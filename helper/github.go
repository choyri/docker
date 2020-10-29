package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"syscall"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

type GitHub struct {
	*github.Client
	Repository
}

type Repository struct {
	Owner string
	Repo  string
}

func NewGitHub(token string) *GitHub {
	var httpClient *http.Client
	if token != "" {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		httpClient = oauth2.NewClient(nil, ts)
	}

	owner, repo := getCurrentRepository()

	return &GitHub{
		Client: github.NewClient(httpClient),
		Repository: Repository{
			Owner: owner,
			Repo:  repo,
		},
	}
}

func (gh *GitHub) GetLatestVersionOfModule(ctx context.Context, module ModuleType) (ret string, err error) {
	var refs []*github.Reference

	refs, err = gh.ListMatchingTags(ctx, nil, string(module))
	if err != nil {
		return
	}

	if len(refs) == 0 {
		err = errors.New("tags 列表为空，请检查")
		return
	}

	ref := strings.Split(*refs[len(refs)-1].Ref, "-")

	if !strings.HasSuffix(ref[0], string(module)) {
		err = fmt.Errorf("ref 前缀（%s）不匹配，请检查", ref[0])
		return
	}

	ret = ref[1]

	return
}

func (gh *GitHub) GetLatestTagOfRepo(ctx context.Context, repo *Repository, opts *github.ListOptions) (ret string, err error) {
	if repo == nil || repo.Owner == "" || repo.Repo == "" {
		err = errors.New("参数 repo 不能为空")
		return
	}

	var tags []*github.RepositoryTag

	// 返回的 tags 是按 name 倒序的
	tags, _, err = gh.Repositories.ListTags(ctx, repo.Owner, repo.Repo, opts)
	if err != nil {
		return
	}

	if len(tags) == 0 {
		return
	}

	ret = *tags[0].Name

	return
}

func (gh *GitHub) CreateTag(ctx context.Context, tag string) (err error) {
	sha := os.Getenv("GITHUB_SHA")
	if sha == "" {
		err = errors.New("环境变量 GITHUB_SHA 不存在")
		return
	}

	_, _, err = gh.Git.CreateRef(ctx, gh.Owner, gh.Repo, &github.Reference{
		Ref:    github.String("refs/tags/" + tag),
		Object: &github.GitObject{SHA: github.String(sha)},
	})

	return
}

func (gh *GitHub) ListMatchingTags(ctx context.Context, repo *Repository, tag string) (ret []*github.Reference, err error) {
	// 返回的 refs 是正序的
	ret, _, err = gh.Git.ListMatchingRefs(ctx, gh.getOwner(repo), gh.getRepo(repo), &github.ReferenceListOptions{
		Ref: "tags/" + tag,
	})

	return ret, err
}

func (gh *GitHub) getOwner(repo *Repository) string {
	if repo != nil {
		return repo.Owner
	}

	return gh.Owner
}

func (gh *GitHub) getRepo(repo *Repository) string {
	if repo != nil {
		return repo.Repo
	}

	return gh.Repo
}

func getCurrentRepository() (string, string) {
	v, found := syscall.Getenv("GITHUB_REPOSITORY")
	if !found {
		log.Fatal("环境变量 GITHUB_REPOSITORY 不存在")
	}

	vv := strings.Split(v, "/")
	if len(vv) < 2 {
		log.Fatalf("环境变量 GITHUB_REPOSITORY 格式错误：%s", v)
	}

	return vv[0], vv[1]
}
