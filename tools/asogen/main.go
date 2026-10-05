// Command asogen generates a Krateo Azure blueprint from an Azure Service Operator (ASO) CRD.
//
// It reads a CustomResourceDefinition (file or URL), projects its GA spec schema into a
// crdgen-safe values.schema.json, and renders a full blueprint directory
// (chart + CompositionDefinition + customform + README) under <out>/<service>/<resource>.
//
// Usage:
//
//	asogen -crd <file|url> [-version v1api20230101] [-out blueprints] [-chart-version 0.1.0] [-lint]
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	crdSrc := flag.String("crd", "", "path or URL to an Azure Service Operator CustomResourceDefinition (required)")
	version := flag.String("version", "", "CRD version to target (default: newest GA non-storage version)")
	outRoot := flag.String("out", "blueprints", "output root directory for blueprints")
	chartVersion := flag.String("chart-version", "0.1.0", "chart/Composition version to stamp into CompositionDefinition and customform")
	doLint := flag.Bool("lint", false, "run `helm template` on the generated chart (requires helm)")
	flag.Parse()

	if *crdSrc == "" {
		fmt.Fprintln(os.Stderr, "error: -crd is required")
		flag.Usage()
		os.Exit(2)
	}

	if err := run(*crdSrc, *version, *outRoot, *chartVersion, *doLint); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(crdSrc, version, outRoot, chartVersion string, doLint bool) error {
	c, err := loadCRD(crdSrc)
	if err != nil {
		return err
	}
	// ASO-specific: never the storage (conversion hub) version, never a preview.
	ver, err := pickASOVersion(c, version)
	if err != nil {
		return err
	}
	spec, err := ver.specSchema()
	if err != nil {
		return err
	}

	service, err := serviceFromGroup(c.Spec.Group)
	if err != nil {
		return err
	}
	kind := c.Spec.Names.Kind
	cn := chartName(service, kind)
	compKind := compositionKind(cn)

	served := 0
	for _, v := range c.Spec.Versions {
		if v.Served {
			served++
		}
	}

	title := fmt.Sprintf("Azure %s %s", titleCaseService(service), kind)

	data := blueprintData{
		Group:        c.Spec.Group,
		AsoVersion:   ver.Name,
		AsoKind:      kind,
		AsoKindLower: strings.ToLower(kind),
		Service:      service,
		ServiceTitle: titleCaseService(service),
		Slug:         resourceSlug(service, kind),
		ChartName:    cn,
		CompKind:     compKind,
		CompPlural:   compositionPlural(compKind),
		CompVersion:  compositionVersion(chartVersion),
		ChartVersion: chartVersion,
		Namespace:    "azure-" + service + "-system",
		Title:        title,
		ServedCount:  served,
	}

	schema := buildValuesSchema(spec, title,
		fmt.Sprintf("Provision an Azure %s %s via Azure Service Operator. Fields mirror the %s/%s %s spec.",
			data.ServiceTitle, kind, data.Group, data.AsoVersion, kind))
	schemaJSON, err := marshalSchema(schema)
	if err != nil {
		return err
	}

	// Seed values.yaml from a minimal-valid instance of the curated spec, plus credentialFrom.
	values, _ := minimalValid(curate(spec)).(map[string]interface{})
	if values == nil {
		values = map[string]interface{}{}
	}
	values["credentialFrom"] = ""
	valuesYAML, err := marshalValues(values, data)
	if err != nil {
		return err
	}

	outDir := filepath.Join(outRoot, service, data.Slug)
	if err := writeBlueprint(outDir, data, schemaJSON, valuesYAML); err != nil {
		return err
	}
	fmt.Printf("generated %s\n  chart=%s kind=%s plural=%s version=%s apiversion=%s served=%d\n",
		outDir, cn, compKind, data.CompPlural, data.CompVersion, ver.Name, served)

	if doLint {
		if err := lintChart(filepath.Join(outDir, "chart")); err != nil {
			return fmt.Errorf("lint failed: %w", err)
		}
		fmt.Println("  helm template: OK")
	}
	return nil
}

// lintChart stamps a temp version into a copy of the chart and runs helm template.
func lintChart(chartDir string) error {
	helm, err := exec.LookPath("helm")
	if err != nil {
		return fmt.Errorf("helm not found in PATH")
	}

	tmp, err := os.MkdirTemp("", "asogen-lint-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)
	dst := filepath.Join(tmp, "chart")
	if out, err := exec.Command("cp", "-r", chartDir, dst).CombinedOutput(); err != nil {
		return fmt.Errorf("copy chart: %v: %s", err, out)
	}

	chartYaml := filepath.Join(dst, "Chart.yaml")
	b, err := os.ReadFile(chartYaml)
	if err != nil {
		return err
	}
	if err := os.WriteFile(chartYaml, []byte(strings.ReplaceAll(string(b), "CHART_VERSION", "0.0.0-ci")), 0o644); err != nil {
		return err
	}

	cmd := exec.Command(helm, "template", "release", dst, "--namespace", "asogen-test")
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("helm template: %v\n%s", err, out)
	}
	_ = exec.Command(helm, "lint", dst).Run()
	return nil
}
