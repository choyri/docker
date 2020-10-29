package main

import (
	"context"
	"log"
)

func NewHelper(githubToken string) *Helper {
	return &Helper{
		gh: NewGitHub(githubToken),
	}
}

func (h *Helper) LoadModule(module string) *Helper {
	if h == nil {
		log.Fatal("nil helper，停止执行")
	}

	var m ModuleInterface

	switch ModuleType(module) {
	case ModuleNginx:
		m = NewNginx(h)
	default:
		log.Fatalf("不支持的 Module [%s]", module)
	}

	h.module = m

	return h
}

func (h *Helper) Run() {
	if h == nil {
		log.Fatal("nil helper，停止执行")
	}

	var (
		err           error
		localVersion  string
		remoteVersion string
		shouldUpdate  bool
		updateInfo    string
		newTag        string
	)

	localVersion, err = h.module.getLatestVersionOfLocal()
	if err != nil {
		h.logf("获取本地最新版本失败，停止执行：%s", err)
		return
	}

	remoteVersion, err = h.module.getLatestVersionOfRemote()
	if err != nil {
		h.logf("获取远程最新版本失败，停止执行：%s", err)
		return
	}

	h.logf("本地版本：%s / 远程版本：%s", localVersion, remoteVersion)

	if gte(localVersion, remoteVersion) {
		h.logf("本地版未低于远程版本，停止执行")
		return
	}

	shouldUpdate, updateInfo, err = h.module.shouldUpdate(localVersion, remoteVersion)
	if err != nil {
		h.logf(err.Error())
		return
	}

	if !shouldUpdate {
		h.logf("无需更新，停止执行")
		return
	}

	if updateInfo != "" {
		h.logf(updateInfo)
	}

	newTag = h.renderTag(remoteVersion)

	err = h.gh.CreateTag(context.Background(), newTag)
	if err != nil {
		h.logf("创建 tag [%s] 失败：%s", newTag, err)
		return
	}

	h.logf("创建 tag [%s] 成功", newTag)
}

func (h *Helper) logf(format string, v ...interface{}) {
	log.Printf("[%s] "+format, append([]interface{}{h.module.getName()}, v...)...)
}

func (h *Helper) renderTag(remoteVersion string) string {
	return h.module.getName() + "-" + remoteVersion
}
