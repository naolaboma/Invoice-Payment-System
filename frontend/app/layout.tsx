import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { Navbar } from "@/components/layout/navbar";
import QueryProvider from "@/providers/query-provider";
import { cn } from "@/lib/utils";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Invoice Payment System",
  description: "Modern invoice payment system",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body
        className={cn(
          inter.className,
          "min-h-screen bg-background font-sans antialiased"
        )}
      >
        <QueryProvider>
          <div className="relative flex min-h-screen flex-col">
            <Navbar />
            <main className="flex-1 container py-6">{children}</main>
          </div>
        </QueryProvider>
      </body>
    </html>
  );
}
