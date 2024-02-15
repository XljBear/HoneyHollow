import axios from "axios";
import config from '../config'
import { ElNotification } from "element-plus";

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
    return res.data.data
}, err => {
    ElNotification.error({
        message: "程序发生故障，请联系开发者！",
    });
    return Promise.reject(err);
})

export default axios