import axios from "axios"

// const isDev = process.env.NODE_ENV === "development"

export const service = axios.create({
  baseURL: "http://webcode/api/v1/",
})

// 可添加鉴权的前处理
// service.interceptors.request.use(config => {
//   config.data = {
//     ...config.data,
//     // authToken: window.localStorage.getItem('authToken'),
//     authToken: "authToken"
//   }
//   return config
// })

// 可添加消息统一处理
// service.interceptors.response.use(res => {
//   if (res.data.code === 200) {
//     return res.data.data
//   } else {
//     //全局处理错误
//     message.error(res.data.errMsg)
//   }
// })
