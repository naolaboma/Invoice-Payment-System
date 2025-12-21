export interface Invoice {
  id: string;
  senderEmail: string;
  payerEmail: string;
  amount: number;
  currency: string;
  description: string;
  status: "PENDING" | "PAID" | "FAILED";
  payment_link?: string;
  santimPayRef?: string;
  createdAt: string;
  updatedAt: string;
}

export interface CreateInvoiceDTO {
  senderEmail: string;
  payerEmail: string;
  amount: number;
  currency: string;
  description: string;
}
