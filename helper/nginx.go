package main

import (
	"context"

	"github.com/google/go-github/v32/github"
)

const (
	ModuleNginx ModuleType = "nginx"
	DockerNginx string     = "library/nginx"
)

type nginxModule Module

func (m *nginxModule) getName() string {
	return string(ModuleNginx)
}

func (m *nginxModule) getLatestVersionOfLocal() (ret string, err error) {
	return m.h.gh.GetLatestVersionOfModule(context.Background(), ModuleNginx)
}

func (m *nginxModule) getLatestVersionOfRemote() (ret string, err error) {
	return m.h.gh.GetLatestTagOfRepo(context.Background(), &Repository{
		Owner: "nginxinc",
		Repo:  "docker-nginx",
	}, &github.ListOptions{PerPage: 1})
}

func (m *nginxModule) shouldUpdate(localVersion, remoteVersion string) (ret bool, info string, err error) {
	var (
		tag     *DockerImageTag
		tagName = remoteVersion + "-alpine"
	)

	tag, err = isExistsTargetDockerImage(DockerNginx, tagName, ArchAmd64)
	if err != nil {
		return
	}

	ret = true
	info = "发现新镜像 " + tag.getInfoString()

	return
}

func NewNginx(h *Helper) ModuleInterface {
	return &nginxModule{
		h: h,
	}
}
