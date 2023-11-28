import axios from "axios";

const baseURL = "http://localhost:8081/";
const instance = axios.create({
  baseURL: baseURL,
  headers: {
    "Content-Type": "application/json",
  },
});

export const getList = (requestConfig = {}) => {
  return instance.get(`${baseURL}clients`, {
    ...requestConfig,
  })
}

export const get = (id, requestConfig = {}) => {
  return instance.get(`${baseURL}clients/${id}`, {
    ...requestConfig,
  })
}

export const create = (data, requestConfig = {}) => {
  return instance.post(`${baseURL}clients`, data, {
    ...requestConfig,
  });
}

export const update = (id, data, requestConfig = {}) => {
  return instance.put(`${baseURL}clients/${id}`, data, {
    ...requestConfig,
  });
}

export const remove = (id, requestConfig = {}) => {
  return instance.delete(`${baseURL}clients/${id}`, {
    ...requestConfig,
  });
}