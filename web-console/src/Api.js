import axios from "axios"
import qs from "qs"

export const HTTP_VERBS = {
  DELETE: "delete",
  GET: "get",
  PATCH: "patch",
  POST: "post",
  PUT: "put",
}

const myAxios = axios.create({
  paramsSerializer: (params) => qs.stringify(params, { arrayFormat: "repeat" }),
})

const newRequest = (method, url, queryParams = {}, data, headers = {}) => {
  return myAxios.request({
    method: method,
    url: getFinalUrl(url),
    data: data,
    headers: { ...headers },
    params: queryParams,
  })
}

const getFinalUrl = (url) => {
  return "/api" + url
}

export const api = {
  version: {
    get: () => newRequest(HTTP_VERBS.GET, "/version", {}, {}),
  },
  youtube: {
    get: (url) => newRequest(HTTP_VERBS.GET, "/youtube/info", { url: url }, {}),
  },
}
