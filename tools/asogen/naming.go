package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"

	"github.com/gobuffalo/flect"
)

// asoGroupSuffix is the API group suffix every Azure Service Operator CRD carries.
const asoGroupSuffix = ".azure.com"

// toGolangName mirrors core-provider's internal strutil.ToGolangName so the Kind we compute
// matches the one core-provider derives from the chart name.
func toGolangName(s string) string {
	buf := bytes.NewBuffer([]byte{})
	for i, v := range splitOnAll(s, isNotAGoNameCharacter) {
		if i == 0 && strings.IndexAny(v, "0123456789") == 0 {
			buf.WriteRune('_')
		}
		buf.WriteString(capitaliseFirstLetter(v))
	}
	return buf.String()
}

func capitaliseFirstLetter(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[0:1]) + s[1:]
}

func splitOnAll(s string, shouldSplit func(r rune) bool) []string {
	rv := []string{}
	buf := bytes.NewBuffer([]byte{})
	for _, c := range s {
		if shouldSplit(c) {
			rv = append(rv, buf.String())
			buf.Reset()
		} else {
			buf.WriteRune(c)
		}
	}
	if buf.Len() > 0 {
		rv = append(rv, buf.String())
	}
	return rv
}

func isNotAGoNameCharacter(r rune) bool {
	return !(unicode.IsLetter(r) || unicode.IsDigit(r))
}

// compositionKind reproduces core-provider's gvr.go: flect.Pascalize(toGolangName(chartName)).
func compositionKind(chartName string) string {
	return flect.Pascalize(toGolangName(chartName))
}

// compositionPlural reproduces the plural that crdgen stamps into the generated Composition
// CRD's spec.names.plural, which is what core-provider later reads back from apiserver discovery
// (it does NOT pluralize offline itself).
//
// ORDER MATTERS: crdgen does Pluralize-then-Lower (plumbing transpile.go:660), not
// Lower-then-Pluralize. The two are not equivalent -- "AzureSubscriptionAlias" gives
// "azuresubscriptionaliases" one way and "azuresubscriptionalias" the other, because flect's
// pluralization is case-sensitive. Getting this backwards puts a plural that does not exist into
// the customform's resource path. Mirror crdgen exactly rather than relying on the two orders
// happening to coincide.
func compositionPlural(kind string) string {
	return strings.ToLower(flect.Pluralize(kind))
}

// serviceFromGroup turns an Azure Service Operator API group like "storage.azure.com"
// into "storage". It rejects groups that are not ASO groups, so pointing the generator at a
// non-ASO CRD fails loudly instead of producing a plausible-looking blueprint.
func serviceFromGroup(group string) (string, error) {
	if !strings.HasSuffix(group, asoGroupSuffix) {
		return "", fmt.Errorf("group %q is not an Azure Service Operator group (expected *%s)", group, asoGroupSuffix)
	}
	service := strings.TrimSuffix(group, asoGroupSuffix)
	if service == "" {
		return "", fmt.Errorf("cannot derive a service id from group %q", group)
	}
	// A few ASO groups nest a sub-service ("network.frontdoor.azure.com"). Flatten the dots to
	// dashes so the id stays a single DNS-safe label usable in a chart name and namespace.
	return strings.ReplaceAll(service, ".", "-"), nil
}

// resourceSlug is the per-resource half of the chart name.
//
// Most ASO Kinds do not repeat their service ("Account" lives in the "storage" group), but some
// do ("NetworkInterface" in "network"). Stripping a redundant service prefix is therefore a
// no-op for the common case and an improvement for the rest, matching the AWS catalog's shape.
// The prefix is only stripped when a non-empty remainder survives.
func resourceSlug(service, kind string) string {
	lower := strings.ToLower(kind)
	if trimmed := strings.TrimPrefix(lower, strings.ToLower(service)); trimmed != "" && trimmed != lower {
		return trimmed
	}
	return lower
}

// chartName builds the blueprint chart name, e.g. azure-storage-account.
func chartName(service, kind string) string {
	return "azure-" + service + "-" + resourceSlug(service, kind)
}

// compositionVersion maps a chart semver (0.1.0) to the Composition CRD version (v0-1-0).
func compositionVersion(chartVersion string) string {
	return "v" + strings.ReplaceAll(chartVersion, ".", "-")
}

// titleCaseService renders a service id for human-facing text (storage -> Storage).
func titleCaseService(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
