"use client";

import { useQuery } from "@tanstack/react-query";
import { invoiceApi } from "@/lib/api";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { formatCurrency, formatDate } from "@/lib/utils";
import { ArrowLeft, CreditCard, Download, Share2 } from "lucide-react";
import Link from "next/link";
import { useParams } from "next/navigation";

export default function InvoiceDetailsPage() {
  const params = useParams();
  const id = params.id as string;

  const {
    data: invoice,
    isLoading,
    error,
  } = useQuery({
    queryKey: ["invoice", id],
    queryFn: () => invoiceApi.getById(id),
  });

  if (isLoading) {
    return (
      <div className="space-y-6">
        <div className="flex items-center gap-4">
          <div className="h-10 w-10 rounded-full bg-muted animate-pulse" />
          <div className="h-8 w-48 bg-muted animate-pulse rounded-md" />
        </div>
        <div className="grid gap-6 md:grid-cols-2">
          <div className="h-[400px] bg-muted animate-pulse rounded-lg" />
          <div className="h-[200px] bg-muted animate-pulse rounded-lg" />
        </div>
      </div>
    );
  }

  if (error || !invoice) {
    return (
      <div className="flex flex-col items-center justify-center h-[50vh] gap-4">
        <h2 className="text-2xl font-bold text-destructive">
          Invoice not found
        </h2>
        <Button asChild variant="outline">
          <Link href="/">Go Back</Link>
        </Button>
      </div>
    );
  }

  const getStatusVariant = (status: string) => {
    switch (status) {
      case "PAID":
        return "success";
      case "PENDING":
        return "warning";
      case "FAILED":
        return "destructive";
      default:
        return "default";
    }
  };

  return (
    <div className="space-y-6 max-w-5xl mx-auto">
      <div className="flex items-center justify-between">
        <div className="flex items-center gap-4">
          <Button variant="ghost" size="icon" asChild>
            <Link href="/">
              <ArrowLeft className="h-4 w-4" />
            </Link>
          </Button>
          <div>
            <h1 className="text-2xl font-bold tracking-tight">
              Invoice #{invoice.id.slice(-6).toUpperCase()}
            </h1>
            <p className="text-sm text-muted-foreground">
              Created on {formatDate(invoice.createdAt)}
            </p>
          </div>
        </div>
        <div className="flex items-center gap-2">
          <Button variant="outline" size="sm">
            <Download className="mr-2 h-4 w-4" />
            Download PDF
          </Button>
          <Button variant="outline" size="sm">
            <Share2 className="mr-2 h-4 w-4" />
            Share
          </Button>
        </div>
      </div>

      <div className="grid gap-6 md:grid-cols-3">
        <div className="md:col-span-2 space-y-6">
          <Card>
            <CardHeader>
              <CardTitle>Invoice Details</CardTitle>
            </CardHeader>
            <CardContent className="space-y-6">
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <h3 className="font-semibold mb-1">From</h3>
                  <p className="text-sm text-muted-foreground">
                    {invoice.senderEmail}
                  </p>
                </div>
                <div>
                  <h3 className="font-semibold mb-1">To</h3>
                  <p className="text-sm text-muted-foreground">
                    {invoice.payerEmail}
                  </p>
                </div>
              </div>

              <div className="border-t pt-4">
                <h3 className="font-semibold mb-2">Description</h3>
                <p className="text-sm text-muted-foreground">
                  {invoice.description}
                </p>
              </div>

              <div className="border-t pt-4">
                <div className="flex justify-between items-center font-medium">
                  <span>Total Amount</span>
                  <span className="text-xl">
                    {formatCurrency(invoice.amount, invoice.currency)}
                  </span>
                </div>
              </div>
            </CardContent>
          </Card>
        </div>

        <div className="space-y-6">
          <Card>
            <CardHeader>
              <CardTitle>Payment Status</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="flex justify-between items-center">
                <span className="text-sm text-muted-foreground">Status</span>
                <Badge variant={getStatusVariant(invoice.status)}>
                  {invoice.status}
                </Badge>
              </div>

              {invoice.status === "PENDING" && invoice.payment_link && (
                <Button className="w-full" asChild>
                  <a
                    href={invoice.payment_link}
                    target="_blank"
                    rel="noopener noreferrer"
                  >
                    <CreditCard className="mr-2 h-4 w-4" />
                    Pay Now
                  </a>
                </Button>
              )}

              {invoice.status === "PAID" && (
                <div className="rounded-md bg-green-50 p-3 text-sm text-green-700 dark:bg-green-900/20 dark:text-green-400">
                  Payment completed successfully.
                </div>
              )}
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
}
