package main

import (
	"fmt"
	"time"
)

type (
	Arch       string
	ModuleType string
)

type Helper struct {
	gh     *GitHub
	module ModuleInterface
}

type ModuleInterface interface {
	getName() string
	getLatestVersionOfLocal() (string, error)
	getLatestVersionOfRemote() (string, error)
	shouldUpdate(localVersion, remoteVersion string) (bool, string, error)
}

type Module struct {
	h *Helper
}

type DockerImageTag struct {
	Name        string            `json:"name"`
	LastUpdated time.Time         `json:"last_updated"`
	Images      []DockerImageInfo `json:"images"`
}

func (t *DockerImageTag) getInfoString() string {
	return fmt.Sprintf("[tag:%s] [%s]", t.Name, t.Images[0].Digest)
}

type DockerImageInfo struct {
	Architecture Arch   `json:"architecture"`
	Digest       string `json:"digest"`
	Size         uint64 `json:"size"`
}
