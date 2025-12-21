//mongodb setup script for invoice payment system
// usage : mongosh < mongodb_setup_invoice.js

use invoice_payment_db;

print("setting up invoice payment sytem database..");

// create invoices collection with indexes

db.createCollection("invoices");
db.invoices.createIndex({"_id": 1}, {unique: true});
db.invoices.createIndex({"payerEmail": 1});
db.invoices.createIndex({"senderEmail": 1});
db.invoices.createIndex({"status": 1});
db.invoices.createIndex({"createdAt": -1});

print("invoices collection created with indexes");

// create payment collection to store callback logs

db.createCollection("payments");
db.payments.createIndex({"invoiceID": 1});
db.payments.createIndex({"reference": 1}, {unique: true});
db.payments.createIndex({"status": 1});
db.payments.createIndex({"createdAt": -1});

print("payments collection created with indexes");

//insert sample invoice for testing
db.invoices.insertOne({
    _id: "inv_001",
    senderEmail: "merchant@example.com",
    payerEmail: "payer@example.com",
    amount: 500.0,
    currency: "ETB",
    description: "website hosting invoice",
    status: "PENDING",
    santimPayRef: null,
    createdAt: new Date(),
    updatedAt: new Date()
});

print("sample invoice inserted");

// show collectoins
print("\nDatabase Collectoins:");
db.getCollectionNames().forEach(function(collection){
    print("  - " + collection);
});

print("\MongoDB setup completed successfully!");