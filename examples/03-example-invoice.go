package main

import (
    "fmt"
    "github.com/mk-j/go-netsuite-soap"
)


func main() {
    fmt.Println("starting...")

    invoiceInternalId:= CreateInvoice()
    fmt.Printf("CREATED INVOICE: %v\n", invoiceInternalId)
    if len(invoiceInternalId)>0 {
        invoice:= GetInvoice(invoiceInternalId)
        fmt.Printf("GET INVOICE: %v\n", netsuite.JsonEncode(invoice))
    }
    if len(invoiceInternalId)>0 {
        DeleteInvoice(invoiceInternalId)
        fmt.Printf("DELETED INVOICE: %v\n", invoiceInternalId)
    }
}

//----------------------------------------------------------------------
func CreateInvoice() string {
    conn := netsuite.NetSuiteConnector{}
    conn.ReadConfig("./config/config.yml")

    customFields:= netsuite.CustomFieldList{
        CustomField:[]*netsuite.CustomFieldRef{
                netsuite.NSCustomStringField("custbody_invoice_method","email"),
                netsuite.NSCustomDateField("custbody_upstream_create_date","2020-01-01"),
            },
        }

    address := netsuite.Address{
        Attention: "Harrison Fjord",
        Addressee: "LLC Corp Inc.",
        AddrPhone: "867-5309",
        Addr1:     "1600 Amphitheatre Parkway",
        Addr2:     "",
        City:      "Mountain View",
        State:     "CA",
        Zip:       "94043",
        Country:   "_unitedStates",
    }

    invoice := netsuite.Invoice{
        BillingAddress:&address,
        Email:"kettle-cooked@example.com",
        ExternalId:"INV23455-" + netsuite.RandString(4),
        CustomForm:&netsuite.RecordRef{InternalId:"123",},
        Subsidiary:&netsuite.RecordRef{InternalId:"1",},
        TranDate:netsuite.XDateTimeFromISOString("2022-11-11"),
        Terms:&netsuite.RecordRef{InternalId:"2",},
        Entity:&netsuite.RecordRef{ExternalId:"CX90210",},
        ItemList:&netsuite.InvoiceItemList{
            Item:[]*netsuite.InvoiceItem{
                &netsuite.InvoiceItem{
                    Item:&netsuite.RecordRef{InternalId:"555",},
                    Line:1,
                    Description:"Contract Fee",
                    Quantity:1.0,
                    Amount:3000.0,
                },
            },
        },
        CustomFieldList:&customFields,
    }
    r, err := conn.SoapService().AddInvoice(&invoice)
    if err!=nil {
        fmt.Printf("Error: %s", err)
    }
    return r
}
//----------------------------------------------------------------------
func GetInvoice(invoiceInternalId string) *netsuite.Invoice{
    conn := netsuite.NetSuiteConnector{}
    conn.ReadConfig("./config/config.yml")
    r, err := conn.SoapService().GetInvoice(invoiceInternalId)
    if err!=nil {
        fmt.Printf("Error: %s", err)
    }
    return r
}
//----------------------------------------------------------------------
func DeleteInvoice(invoiceInternalId string) bool{
    conn := netsuite.NetSuiteConnector{}
    conn.ReadConfig("./config/config.yml")
    r, err := conn.SoapService().DeleteInvoice(invoiceInternalId)
    if err!=nil {
        fmt.Printf("Error: %s", err)
    }
    return r
}
//----------------------------------------------------------------------
