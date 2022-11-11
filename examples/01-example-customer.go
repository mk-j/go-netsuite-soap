package main

import (
    "fmt"
    "github.com/mk-j/go-netsuite-soap"
)

func main() {
    fmt.Println("starting...")

    customerInternalId:= CreateCustomer()
    fmt.Printf("CREATED CUSTOMER: %v\n", customerInternalId)
    if len(customerInternalId)>0 {
        cust:= GetCustomer(customerInternalId)
        fmt.Printf("GET CUSTOMER: %v\n", netsuite.JsonEncode(cust))
    }
    if len(customerInternalId)>0 {
        DeleteCustomer(customerInternalId)
        fmt.Printf("DELETED CUSTOMER: %v\n", customerInternalId)
    }
}

//----------------------------------------------------------------------
func CreateCustomer() string {
    conn := netsuite.NetSuiteConnector{}
    conn.ReadConfig("./config/config.yml")
  
    customer := netsuite.Customer{
        CompanyName:"Elephant and Bubbles Waffle Company",
        Email:"kettle-cooked@example.com",
        //ExternalId:"CX90210" + netsuite.RandString(4),
        TranDate:netsuite.XDateTimeFromISOString("2022-11-11"),
        ExternalId:"CX90210",
        CustomForm:&netsuite.RecordRef{InternalId:"125",},
        Subsidiary:&netsuite.RecordRef{InternalId:"1",},
        Terms:&netsuite.RecordRef{InternalId:"2",},
        //CustomFieldList:&customFields,
    }
    r, err := conn.SoapService().AddCustomer(&customer)
    if err!=nil {
        fmt.Printf("Error: %s\n", err)
    }
    return r
}
//----------------------------------------------------------------------
func GetCustomer(customerInternalId string) *netsuite.Customer {
    conn := netsuite.NetSuiteConnector{}
    conn.ReadConfig("./config/config.yml")
    r, err := conn.SoapService().GetCustomer(customerInternalId)
    if err!=nil {
        fmt.Printf("Error: %s\n", err)
    }
    return r
}
//----------------------------------------------------------------------
func DeleteCustomer(customerInternalId string) bool {
    conn := netsuite.NetSuiteConnector{}
    conn.ReadConfig("./config/config.yml")
    r, err := conn.SoapService().DeleteCustomer(customerInternalId)
    if err!=nil {
        fmt.Printf("Error: %s\n", err)
    }
    return r
}
//----------------------------------------------------------------------
