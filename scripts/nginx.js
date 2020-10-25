import api from './api.js'
import common from './common.js'
import github from './github.js'
import semverGte from 'semver/functions/gte.js'

const MODULE_NAME = 'nginx'
const URL_DOCKER_TAGS = 'https://hub.docker.com/v2/repositories/library/nginx/tags'

async function getLatestVersionOfLocal() {
    const version = await common.getLatestVersionOfGit(MODULE_NAME, MODULE_NAME)
    if (!version) {
        return
    }

    common.info(MODULE_NAME, `当前本仓库最新 version 为 [${version}]`)

    return version
}

async function getLatestVersionOfRemoteGit() {
    const tags = await github.listTags({
        owner: 'nginxinc',
        repo: 'docker-nginx',
        per_page: 1,
    })
    if (!tags || !tags.data || tags.data.length < 1) {
        common.error(module, '获取不到 docker-nginx git tags')
        return
    }

    const version = tags.data[0].name
    common.info(MODULE_NAME, `当前最新的 docker-nginx version 为 [${version}]`)

    return version
}

async function isExistTargetDockerImage(version) {
    const tagName = version + '-alpine'

    const resp = await api.get(URL_DOCKER_TAGS, {name: tagName})
    if (!resp) {
        common.error(MODULE_NAME, '获取不到指定的 docker-nginx docker tags')
        return
    }

    const info = resp.results.find(item => item.name === tagName)
    if (!info) {
        common.info(MODULE_NAME, '指定的 docker-nginx docker tags 还未生成')
        return
    }

    const image = info.images.find(item => item.architecture === 'amd64')
    if (!image) {
        common.error(MODULE_NAME, '指定的 docker-nginx docker arch 还未生成')
        return
    }

    common.info(MODULE_NAME, `目标镜像已出现 [nginx:${tagName}] ${image.digest}`)

    return true
}

export default async function run() {
    const localVersion = await getLatestVersionOfLocal()
    if (!localVersion) {
        common.error(MODULE_NAME, '获取不到 local version，当前模块停止执行')
        return
    }

    const remoteGitVersion = await getLatestVersionOfRemoteGit()
    if (!remoteGitVersion) {
        common.error(MODULE_NAME, '获取不到 remote git version，当前模块停止执行')
        return
    }

    if (semverGte(localVersion, remoteGitVersion)) {
        common.info(MODULE_NAME, '本地版本大于等于远程版本，当前模块停止执行')
        return
    }

    const exists = await isExistTargetDockerImage(remoteGitVersion)
    if (!exists) {
        common.info(MODULE_NAME, 'docker image 不存在，当前模块停止执行')
        return
    }

    const newTag = `nginx-${remoteGitVersion}`
    await github.createTag(newTag)

    common.info(MODULE_NAME, `已创建新 git tag ${newTag}`)
}
