package main

import (
	"flag"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	githubToken  string
	targetModule string
)

func init() {
	flag.StringVar(&githubToken, "github_token", os.Getenv("INPUT_GITHUB_TOKEN"), "GitHub Token")
	flag.StringVar(&targetModule, "module", "", "处理的模块")
}

func main() {
	flag.Parse()
	NewHelper(githubToken).LoadModule(targetModule).Run()
}
