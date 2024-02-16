import axios from "axios";
import config from '../config'
import { ElMessage } from "element-plus";

const env: string = import.meta.env.MODE
axios.defaults.baseURL = config[env as keyof typeof config].baseUrl
axios.defaults.timeout = 5000
axios.defaults.withCredentials = false

axios.interceptors.request.use(function (config) {
    return config;
}, function (error) {
    return Promise.reject(error);
});

axios.interceptors.response.use(res => {

    if (typeof res !== 'object') {
        return Promise.reject(res)
    }
    if (res.status != 200) {
        return Promise.reject(res)
    }
    if (res.data.code === 0) {
        if (res.data.message === "") {
            ElMessage({
                message: '嗷，操作好像失败了～',
                type: 'error',
            })
        } else {
            ElMessage({
                message: res.data.message,
                type: 'warning',
            })
        }
        return Promise.reject(res)
    }
    return res.data.data
}, err => {
    ElMessage({
        message: '糟糕服务器好像出错了！～',
        type: 'error',
    })
    return Promise.reject(err);
})

export default axios