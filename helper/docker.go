package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ArchAmd64 Arch = "amd64"

	URLDockerRepoTags = "https://hub.docker.com/v2/repositories/%s/tags"
)

func getDockerTag(imageName, tagName string) (ret *DockerImageTag, err error) {
	url := fmt.Sprintf(URLDockerRepoTags+"?name=%s", imageName, tagName)

	var (
		resp *http.Response
		data struct {
			Results []DockerImageTag `json:"results"`
		}
	)

	resp, err = http.Get(url)
	if err != nil {
		err = fmt.Errorf("docker image tags 获取失败：%w", err)
		return
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()

	err = json.Unmarshal(respBody, &data)
	if err != nil {
		err = fmt.Errorf("docker image tags 解析失败：%w", err)
		return
	}

	for k, v := range data.Results {
		if v.Name == tagName {
			ret = &data.Results[k]
			break
		}
	}

	if ret == nil {
		err = fmt.Errorf("docker image tag [%s] 还未生成", imageName)
	}

	return
}

func isExistsTargetDockerImage(imageName, tagName string, arch Arch) (tag *DockerImageTag, err error) {
	tag, err = getDockerTag(imageName, tagName)
	if err != nil {
		return
	}

	var info *DockerImageInfo

	for k, v := range tag.Images {
		if v.Architecture == arch {
			info = &tag.Images[k]
		}
	}

	if info == nil {
		err = fmt.Errorf("docker image tag [%s] arch [%s] 还未生成", imageName, ArchAmd64)
		return
	}

	tag.Images = []DockerImageInfo{*info}

	return
}
