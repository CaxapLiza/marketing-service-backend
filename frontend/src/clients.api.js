import instance from "./funcs.js";

export const getList = () => instance.get("clients");

export const get = (id) => instance.get(`clients/${id}`);

export const create = (data) => instance.post("clients", data);

export const update = (id, data) => instance.put(`clients/${id}`, data);

export const remove = (id) => instance.delete(`clients/${id}`);