import * as core from '@actions/core'
import github from './github.js'

function info(module, message) {
    core.info(`[${module}] ${message}`)
}

function error(module, message) {
    core.error(`[${module}] ${message}`)
}

async function getLatestVersionOfGit(module, tag) {
    const tags = await github.listMatchingTag(tag)
    if (!tags || !tags.data) {
        error(module, '获取不到 git tags')
        return
    }

    const latestTag = tags.data.pop()
    if (!latestTag) {
        error(module, '获取不到最新的 git tag')
        return
    }

    const ret = latestTag.ref.split('-').pop()
    if (!ret) {
        error(module, 'git tag 解析失败')
    }

    return ret
}

export default {
    info,
    error,
    getLatestVersionOfGit,
}
