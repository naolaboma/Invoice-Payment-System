import { InvoiceList } from "@/components/invoices/invoice-list";
import { CreateInvoiceDialog } from "@/components/invoices/create-invoice-dialog";

export default function Home() {
  return (
    <div className="space-y-8">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Invoices</h1>
          <p className="text-muted-foreground">
            Manage and track your invoices and payments.
          </p>
        </div>
        <CreateInvoiceDialog />
      </div>
      <InvoiceList />
    </div>
  );
}
