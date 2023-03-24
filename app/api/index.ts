import axios from "axios";

export const ApiClient = axios.create({
  //   baseURL: env.API_URL,
  baseURL: "http://localhost:4000",
  headers: {
    "Content-type": "application/json",
  },
});

export const { get, put, delete: del, post } = ApiClient;
