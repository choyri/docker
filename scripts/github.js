import * as core from '@actions/core'
import * as github from '@actions/github'

let octokit

function getToken() {
    return core.getInput('github_token')
}

function getOctokit() {
    if (!octokit) {
        octokit = github.getOctokit(getToken())
    }

    return octokit
}

async function listMatchingTag(tag) {
    // 返回的 refs 是正序的
    return await getOctokit().git.listMatchingRefs({
        ...github.context.repo,
        ref: `tags/${tag}`,
    })
}

async function listTags(params) {
    params = Object.assign(github.context.repo, params)

    // 返回的 tags 是按 name 倒序的
    return await getOctokit().repos.listTags(params)
}

async function createTag(tag) {
    await getOctokit().git.createRef({
        ...github.context.repo,
        ref: `refs/tags/${tag}`,
        sha: process.env.GITHUB_SHA,
    })
}

export default {
    listMatchingTag,
    listTags,
    createTag,
}
