import axios from "axios";
import { CreateInvoiceDTO, Invoice } from "@/types";

const API_BASE_URL = "http://localhost:8080/api/v1";

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

export const invoiceApi = {
  getAll: async (email?: string) => {
    const params = email ? { email } : {};
    const response = await api.get<Invoice[]>("/invoices/", { params });
    return response.data;
  },

  getById: async (id: string) => {
    const response = await api.get<Invoice>(`/invoices/${id}`);
    return response.data;
  },

  create: async (data: CreateInvoiceDTO) => {
    const response = await api.post<{ message: string; payment_link: string }>(
      "/invoices/",
      data
    );
    return response.data;
  },
};
