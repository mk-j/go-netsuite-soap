package main

import (
    "fmt"
    "github.com/mk-j/go-netsuite-soap"
)


func main() {
    fmt.Println("starting...")

    creditMemoInternalId:= CreateCreditMemo()
    fmt.Printf("CREATED CREDITMEMO: %v\n", creditMemoInternalId)
    if len(creditMemoInternalId)>0 {
        creditmemo:= GetCreditMemo(creditMemoInternalId)
        fmt.Printf("GET CREDITMEMO: %v\n", netsuite.JsonEncode(creditmemo))
    }
    if len(creditMemoInternalId)>0 {
        DeleteCreditMemo(creditMemoInternalId)
        fmt.Printf("DELETED CREDITMEMO: %v\n", creditMemoInternalId)
    }
}

//----------------------------------------------------------------------
func CreateCreditMemo() string {
	conn := netsuite.NetSuiteConnector{}
	conn.ReadConfig("./config/config.yml")
  
    creditMemo := netsuite.CreditMemo{
        ExternalId:"CM912345-" + netsuite.RandString(4),
        Entity:&netsuite.RecordRef{ExternalId:"CX90210",},
        TranDate:netsuite.XDateTimeFromISOString("2022-11-11"),
        Subsidiary:&netsuite.RecordRef{InternalId:"1",},
        ItemList:&netsuite.CreditMemoItemList{
            Item:[]*netsuite.CreditMemoItem{
                &netsuite.CreditMemoItem{
                    Item:&netsuite.RecordRef{InternalId:"555",},
                    Line:1,
                    Description:"Contract Fee",
                    Quantity:1.0,
                    Amount:3000.0,
                },
            },
        },
    }

    r, err := conn.SoapService().AddCreditMemo(&creditMemo)
    if err!=nil {
        fmt.Printf("Error: %s", err)
    }
    return r
}
//----------------------------------------------------------------------
func GetCreditMemo(creditMemoInternalId string) *netsuite.CreditMemo {
    conn := netsuite.NetSuiteConnector{}
    conn.ReadConfig("./config/config.yml")
    r, err := conn.SoapService().GetCreditMemo(creditMemoInternalId)
    if err!=nil {
        fmt.Printf("Error: %s", err)
    }
    return r
}
//----------------------------------------------------------------------
func DeleteCreditMemo(creditMemoInternalId string) bool {
    conn := netsuite.NetSuiteConnector{}
    conn.ReadConfig("./config/config.yml")
    r, err := conn.SoapService().DeleteCreditMemo(creditMemoInternalId)
    if err!=nil {
        fmt.Printf("Error: %s", err)
    }
    return r
}
//----------------------------------------------------------------------
