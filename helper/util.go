package main

import (
	"golang.org/x/mod/semver"
)

func gte(localVersion, remoteVersion string) bool {
	return semver.Compare("v"+localVersion, "v"+remoteVersion) >= 0
}

func renderTag(module ModuleType, version string) string {
	return string(module) + "-" + version
}
