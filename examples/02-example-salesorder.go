package main

import (
    "fmt"
    "github.com/mk-j/go-netsuite-soap"
)


func main() {
    fmt.Println("starting...")

    salesOrderInternalId:= CreateSalesOrder()
    fmt.Printf("CREATED SALESORDER: %v\n", salesOrderInternalId)
    if len(salesOrderInternalId)>0 {
        salesorder:= GetSalesOrder(salesOrderInternalId)
        fmt.Printf("GET SALESORDER: %v\n", netsuite.JsonEncode(salesorder))
    }
    if len(salesOrderInternalId)>0 {
        DeleteSalesOrder(salesOrderInternalId)
        fmt.Printf("DELETED SALESORDER: %v\n", salesOrderInternalId)
    }
}

//----------------------------------------------------------------------
func CreateSalesOrder() string {
    conn := netsuite.NetSuiteConnector{}
    conn.ReadConfig("./config/config.yml")
  
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
       
    salesOrder := netsuite.SalesOrder{
        ShippingAddress:&address,
        BillingAddress:&address,
        ExternalId:"SO34532-" + netsuite.RandString(4),
        CustomForm:&netsuite.RecordRef{InternalId:"124",},
        Subsidiary:&netsuite.RecordRef{InternalId:"1",},
        TranDate:netsuite.XDateTimeFromISOString("2022-11-11"),
        Terms:&netsuite.RecordRef{InternalId:"2",},
        Entity:&netsuite.RecordRef{ExternalId:"CX90210",},
        ItemList:&netsuite.SalesOrderItemList{
            Item:[]*netsuite.SalesOrderItem{
                &netsuite.SalesOrderItem{
                    Item:&netsuite.RecordRef{InternalId:"555",},
                    Line:1,
                    Description:"Contract Fee",
                    Quantity:1.0,
                    Amount:3000.0,
                },
            },
        },
        //CustomFieldList:&customFields,
    }
    r, err := conn.SoapService().AddSalesOrder(&salesOrder)
    if err!=nil {
        fmt.Printf("Error: %s", err)
    }
    return r
}
//----------------------------------------------------------------------
func GetSalesOrder(salesOrderInternalId string) *netsuite.SalesOrder {
    conn := netsuite.NetSuiteConnector{}
    conn.ReadConfig("./config/config.yml")
    r, err := conn.SoapService().GetSalesOrder(salesOrderInternalId)
    if err!=nil {
        fmt.Printf("Error: %s", err)
    }
    return r
}
//----------------------------------------------------------------------
func DeleteSalesOrder(salesOrderInternalId string) bool {
    conn := netsuite.NetSuiteConnector{}
    conn.ReadConfig("./config/config.yml")
    r, err := conn.SoapService().DeleteSalesOrder(salesOrderInternalId)
    if err!=nil {
        fmt.Printf("Error: %s", err)
    }
    return r
}
//----------------------------------------------------------------------
