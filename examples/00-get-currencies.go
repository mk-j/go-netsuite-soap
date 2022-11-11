package main

import (
    "fmt"
    "github.com/mk-j/go-netsuite-soap"
)

func main() {
    fmt.Println("starting...")
    fmt.Println("--getting-currencies-------------")
    curr()
    fmt.Println("--getting-terms-------------")
    terms()
    fmt.Println("--getting-category-by-id-------------")
    categ()
    fmt.Println("---------------")
}
//----------------------------------------------------------------------
func curr() {
    conn := netsuite.NetSuiteConnector{}
    conn.ReadConfig("./config/config.yml")
    currencyList, err := conn.SoapService().GetCurrencies()
    if err!=nil {
        fmt.Printf("Error: %s\n", err)
    }
    for _,c := range(currencyList) {
        fmt.Printf("[Currency]%s %s %s\n", c.InternalId, c.ExternalId, c.Name)
    }
}
//----------------------------------------------------------------------
func terms() {
    conn := netsuite.NetSuiteConnector{}
    conn.ReadConfig("./config/config.yml")
    termsList, err := conn.SoapService().GetTerms()
    if err!=nil {
        fmt.Printf("Error: %s\n", err)
    }
    for _,t := range(termsList) {
        fmt.Printf("[Terms] InternalId:%s; Name:%s\n", t.InternalId, t.Name)
    }
}
//----------------------------------------------------------------------
func categ() {
    conn := netsuite.NetSuiteConnector{}
    conn.ReadConfig("./config/config.yml")
    categoryInternalId:="123"
    customerCategory, err:= conn.SoapService().GetCategory(categoryInternalId)
    if customerCategory!=nil {
        fmt.Printf("CategoryName: %s\n", customerCategory.Name)
        fmt.Printf("CategoryInternalId: %s\n", customerCategory.InternalId)
        fmt.Printf("CategoryExternalId: %s\n", customerCategory.ExternalId)
        fmt.Printf("CategoryIsInactive: %v\n", customerCategory.IsInactive)
    }
    if err!=nil {
        fmt.Printf("Error: %s", err)
    }
}
//----------------------------------------------------------------------
