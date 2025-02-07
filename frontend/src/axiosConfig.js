import axios from 'axios';

const BASE_URL = import.meta.env.VITE_BASE_URL || 'http://127.0.0.1:8080';

const axiosInstance = axios.create({
  baseURL: BASE_URL,
});

export default axiosInstance;