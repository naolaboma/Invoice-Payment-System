import Link from "next/link";
import { Button } from "@/components/ui/button";
import { CreditCard } from "lucide-react";

export function Navbar() {
  return (
    <nav className="border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
      <div className="container flex h-16 items-center justify-between">
        <Link href="/" className="flex items-center gap-2 font-bold text-xl">
          <div className="bg-primary text-primary-foreground p-1.5 rounded-lg">
            <CreditCard className="h-5 w-5" />
          </div>
          <span>InvoicePay</span>
        </Link>
        <div className="flex items-center gap-4">
          <Button variant="ghost" asChild>
            <Link href="/">Invoices</Link>
          </Button>
          <Button asChild>
            <Link href="/?create=true">New Invoice</Link>
          </Button>
        </div>
      </div>
    </nav>
  );
}
