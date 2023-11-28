import axios from "axios";

const instance = axios.create({
    baseURL: "http://localhost:8080/",
    headers: {
        "Content-Type": "application/json",
    },
});

const get = (uri, requestConfig = {}) => {
    return instance.get(`http://localhost:8080/${uri}`, {
        ...requestConfig,
    });
};

const post = (uri, data, requestConfig = {}) => {
    return instance.post(`http://localhost:8080/${uri}`, data, {
        ...requestConfig,
    });
};

const put = (uri, data, requestConfig = {}) => {
    return instance.put(`http://localhost:8080/${uri}`, data, {
        ...requestConfig,
    });
};

const remove = (uri, requestConfig = {}) => {
    return instance.delete(`http://localhost:8080/${uri}`, {
        ...requestConfig,
    });
};

const patch = (uri, data = {}, requestConfig = {}) => {
    return instance.patch(`http://localhost:8080/${uri}`, data, {
        ...requestConfig,
    });
};

const request = (requestConfig) => {
    // ToDo: add base uri
    return instance.request(requestConfig);
};

export default {
    instance,
    get,
    post,
    put,
    remove,
    request,
    patch,
};