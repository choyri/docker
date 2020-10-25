import _ from 'dotenv/config.js'
import * as core from '@actions/core'
import disposeNginx from './nginx.js'

async function main() {
    if (!process.env.GITHUB_SHA) {
        throw 'Missing GITHUB_SHA'
    }

    await disposeNginx()
}


main().catch(error => {
    core.setFailed(error)
    throw error
})
