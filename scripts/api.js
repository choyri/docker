import axios from 'axios'
import common from './common.js';

export default {
    async get(url, params) {
        return new Promise(resolve => {
            axios.get(url, {params: params})
                .then(resp => {
                    resolve(resp.data)
                })
                .catch(err => {
                    common.error('axios', 'get failed: ' + err.message)
                    resolve(null)
                })
        })
    },
}
