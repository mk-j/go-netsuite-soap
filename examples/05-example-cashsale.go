package main

import (
    "fmt"
    "github.com/mk-j/go-netsuite-soap"
)


func main() {
    fmt.Println("starting...")

    cashSaleInternalId:= CreateCashSale()
    fmt.Printf("CREATED CREDITMEMO: %v\n", cashSaleInternalId)
    if len(cashSaleInternalId)>0 {
        cashsale:= GetCashSale(cashSaleInternalId)
        fmt.Printf("GET CREDITMEMO: %v\n", netsuite.JsonEncode(cashsale))
    }
    if len(cashSaleInternalId)>0 {
        DeleteCashSale(cashSaleInternalId)
        fmt.Printf("DELETED CREDITMEMO: %v\n", cashSaleInternalId)
    }
}

//----------------------------------------------------------------------
func CreateCashSale() string {
	conn := netsuite.NetSuiteConnector{}
	conn.ReadConfig("./config/config.yml")
  
    cashSale := netsuite.CashSale{
        ExternalId:"CS76543-" + netsuite.RandString(4),
        Entity:&netsuite.RecordRef{ExternalId:"CX90210",},
        TranDate:netsuite.XDateTimeFromISOString("2022-11-11"),
        Subsidiary:&netsuite.RecordRef{InternalId:"1",},
        ItemList:&netsuite.CashSaleItemList{
            Item:[]*netsuite.CashSaleItem{
                &netsuite.CashSaleItem{
                    Item:&netsuite.RecordRef{InternalId:"555",},
                    Line:1,
                    Description:"Contract Fee",
                    Quantity:1.0,
                    Amount:3000.0,
                },
            },
        },
    }

    r, err := conn.SoapService().AddCashSale(&cashSale)
    if err!=nil {
        fmt.Printf("Error: %s", err)
    }
    return r
}
//----------------------------------------------------------------------
func GetCashSale(cashSaleInternalId string) *netsuite.CashSale {
    conn := netsuite.NetSuiteConnector{}
    conn.ReadConfig("./config/config.yml")
    r, err := conn.SoapService().GetCashSale(cashSaleInternalId)
    if err!=nil {
        fmt.Printf("Error: %s", err)
    }
    return r
}
//----------------------------------------------------------------------
func DeleteCashSale(cashSaleInternalId string) bool {
    conn := netsuite.NetSuiteConnector{}
    conn.ReadConfig("./config/config.yml")
    r, err := conn.SoapService().DeleteCashSale(cashSaleInternalId)
    if err!=nil {
        fmt.Printf("Error: %s", err)
    }
    return r
}
//----------------------------------------------------------------------
