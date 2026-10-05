package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// Azure Service Operator serves several version families per resource:
//
//	v1api20240101          GA, user-facing            <- what a blueprint should target
//	v1api20240101storage   internal conversion hub    <- never user-facing
//	v20240101              GA, legacy alias
//	v20240101preview       preview (pre-GA)
//
// Half of all served versions (767 of 1539 at ASO v2.21.1) are `*storage` hub versions, so the
// "pick the storage version" default that is correct for ACK and Config Connector is WRONG here:
// it would point every blueprint at an internal conversion type.
//
// pickASOVersion therefore selects the newest GA, non-storage version, preferring the v1api
// family over the legacy alias. Preview versions are excluded outright: a resource whose only
// versions are preview has no GA surface and is skipped rather than pinned to a pre-release.
var dateRe = regexp.MustCompile(`(\d{8})`)

func isStorageVersion(v string) bool { return strings.HasSuffix(v, "storage") }
func isPreviewVersion(v string) bool { return strings.Contains(v, "preview") }

func pickASOVersion(c *crd, want string) (*crdVersion, error) {
	byName := map[string]*crdVersion{}
	var served []string
	for i := range c.Spec.Versions {
		v := &c.Spec.Versions[i]
		if !v.Served {
			continue
		}
		byName[v.Name] = v
		served = append(served, v.Name)
	}
	if want != "" {
		if v, ok := byName[want]; ok {
			return v, nil
		}
		return nil, fmt.Errorf("version %q not served by this CRD (served: %s)", want, strings.Join(served, ", "))
	}

	var ga []string
	for _, v := range served {
		if !isStorageVersion(v) && !isPreviewVersion(v) {
			ga = append(ga, v)
		}
	}
	if len(ga) == 0 {
		return nil, fmt.Errorf("no GA version: this resource is preview-only (served: %s); "+
			"skipped rather than pinned to a pre-release", strings.Join(served, ", "))
	}

	// Prefer the v1api family; fall back to the legacy alias when that is all there is.
	pool := []string{}
	for _, v := range ga {
		if strings.HasPrefix(v, "v1api") {
			pool = append(pool, v)
		}
	}
	if len(pool) == 0 {
		pool = ga
	}
	sort.Slice(pool, func(i, j int) bool {
		return dateOf(pool[i]) < dateOf(pool[j])
	})
	return byName[pool[len(pool)-1]], nil
}

func dateOf(v string) string {
	if m := dateRe.FindStringSubmatch(v); m != nil {
		return m[1]
	}
	return ""
}
