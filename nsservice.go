package netsuite

import ( 
    "log"
    "fmt"
    "encoding/xml"
    "errors"
)

type NetSuiteService struct {
    UpstreamService NetSuitePortType
}

//----------------------------------------------------------------------
func (service *NetSuiteService) AddCustomer(customer *Customer) (string, error) {
    customer.XsiType="platformCore:Customer" //this must be set, for the transaction to succeed
    request := Customer_AddRequest{
        Record: customer,
        XmlNSXSI: "http://www.w3.org/2001/XMLSchema-instance",
        XmlNSPC: "urn:relationships_2022_1.lists.webservices.netsuite.com",
        XmlNS1: "urn:core_2022_1.platform.webservices.netsuite.com", //needed for StringCustomFieldRef types
    }
    internalId:=""
    response, err := service.UpstreamService.Customer_Add(&request)
    if (response!=nil && response.WriteResponse.Status.IsSuccess) {
        internalId = response.WriteResponse.BaseRef.InternalId
    } else if err != nil {
        err = distillSoapError(err)
    } else if response.WriteResponse.Status!=nil {
        if len(response.WriteResponse.Status.StatusDetail)>0 {
            err=errors.New(response.WriteResponse.Status.StatusDetail[0].Message)
        }
    }
    return internalId, err
}
//----------------------------------------------------------------------
func (service *NetSuiteService) GetCustomer(customerInternalId string) (*Customer,error) {
    recordType := RecordTypeCustomer
    baseRef := BaseRef{
            InternalId:customerInternalId,
            Type:&recordType,
            XsiType:"platformCore:RecordRef",                
        }
    request := Customer_GetRequest{
        BaseRef:&baseRef,
        XmlNSXSI: "http://www.w3.org/2001/XMLSchema-instance",
        XmlNSPC:"urn:core_2022_1.platform.webservices.netsuite.com",
    }
    var customer *Customer=nil
    response, err := service.UpstreamService.Customer_Get(&request)
    if (response!=nil && response.ReadResponse.Status.IsSuccess) {
        customer = response.ReadResponse.Record
    } else if err != nil {
        err = distillSoapError(err)
    } else if response.ReadResponse.Status!=nil {
        if len(response.ReadResponse.Status.StatusDetail)>0 {
            err=errors.New(response.ReadResponse.Status.StatusDetail[0].Message)
        }
    }
    return customer, err
}
//----------------------------------------------------------------------
func (service *NetSuiteService) DeleteCustomer(customerInternalId string) (bool,error) {
    recordType := RecordTypeCustomer
    recordInternalId := customerInternalId
    return service.deleteRecord(recordType, recordInternalId)
}
//----------------------------------------------------------------------
//----------------------------------------------------------------------
func (service *NetSuiteService) AddSalesOrder(salesOrder *SalesOrder) (string,error) {
    salesOrder.XsiType="platformCore:SalesOrder" //this must be set, for the transaction to succeed
    request := SalesOrder_AddRequest{
        Record:salesOrder,
        XmlNSXSI: "http://www.w3.org/2001/XMLSchema-instance",
        XmlNSPC:"urn:sales_2022_1.transactions.webservices.netsuite.com",
        XmlNS1: "urn:core_2022_1.platform.webservices.netsuite.com", //needed for StringCustomFieldRef types
    }
    internalId:=""
    response, err := service.UpstreamService.SalesOrder_Add(&request)
    if (response!=nil && response.WriteResponse.Status.IsSuccess) {
        internalId = response.WriteResponse.BaseRef.InternalId
    } else if err != nil {
        err = distillSoapError(err)
    } else if response.WriteResponse.Status!=nil {
        if len(response.WriteResponse.Status.StatusDetail)>0 {
            err=errors.New(response.WriteResponse.Status.StatusDetail[0].Message)
        }
    }
    return internalId, err
}
//----------------------------------------------------------------------
func (service *NetSuiteService) GetSalesOrder(salesOrderInternalId string) (*SalesOrder,error) {
    recordType := RecordTypeSalesOrder 
    baseRef :=BaseRef{
            InternalId:salesOrderInternalId,
            Type:&recordType,
            XsiType:"platformCore:RecordRef",
        }
    request := SalesOrder_GetRequest{
        BaseRef:&baseRef,
        XmlNSXSI: "http://www.w3.org/2001/XMLSchema-instance",
        XmlNSPC: "urn:core_2022_1.platform.webservices.netsuite.com",
    }
    var salesOrder *SalesOrder=nil
    response, err := service.UpstreamService.SalesOrder_Get(&request)
    if (response!=nil && response.ReadResponse.Status.IsSuccess) {
        salesOrder = response.ReadResponse.Record
    } else if err != nil {
        err = distillSoapError(err)
    } else if response.ReadResponse.Status!=nil {
        if len(response.ReadResponse.Status.StatusDetail)>0 {
            err=errors.New(response.ReadResponse.Status.StatusDetail[0].Message)
        }
    }
    return salesOrder, err
}
//----------------------------------------------------------------------
func (service *NetSuiteService) DeleteSalesOrder(salesOrderInternalId string) (bool,error) {
    recordType := RecordTypeSalesOrder
    recordInternalId := salesOrderInternalId
    return service.deleteRecord(recordType, recordInternalId)
}
//----------------------------------------------------------------------
//----------------------------------------------------------------------
func (service *NetSuiteService) AddInvoice(invoice *Invoice) (string,error) {
    invoice.XsiType = "platformCore:Invoice" //this must be set, for the transaction to succeed
    request := Invoice_AddRequest{
        Record:invoice,
        XmlNSXSI: "http://www.w3.org/2001/XMLSchema-instance",
        XmlNSPC: "urn:sales_2022_1.transactions.webservices.netsuite.com",
        XmlNS1: "urn:core_2022_1.platform.webservices.netsuite.com", //needed for StringCustomFieldRef types
    }
    internalId:=""
    response, err := service.UpstreamService.Invoice_Add(&request)
    if (response!=nil && response.WriteResponse.Status.IsSuccess) {
        internalId = response.WriteResponse.BaseRef.InternalId
    } else if err != nil {
        err = distillSoapError(err)
    } else if response.WriteResponse.Status!=nil {
        if len(response.WriteResponse.Status.StatusDetail)>0 {
            err=errors.New(response.WriteResponse.Status.StatusDetail[0].Message)
        }
    }
    return internalId, err
}
//----------------------------------------------------------------------
func (service *NetSuiteService) GetInvoice(invoiceInternalId string) (*Invoice,error) {
    recordType := RecordTypeInvoice
    baseRef := BaseRef{
            InternalId:invoiceInternalId,
            Type:&recordType,
            XsiType:"platformCore:RecordRef", 
        }
    request := Invoice_GetRequest{
        BaseRef:&baseRef,
        XmlNSXSI: "http://www.w3.org/2001/XMLSchema-instance",
        XmlNSPC: "urn:core_2022_1.platform.webservices.netsuite.com",
    }
    var invoice *Invoice=nil
    response, err := service.UpstreamService.Invoice_Get(&request)
    if (response!=nil && response.ReadResponse.Status.IsSuccess) {
        invoice = response.ReadResponse.Record
    } else if err != nil {
        err = distillSoapError(err)
    } else if response.ReadResponse.Status!=nil {
        if len(response.ReadResponse.Status.StatusDetail)>0 {
            err=errors.New(response.ReadResponse.Status.StatusDetail[0].Message)
        }
    }
    return invoice, err
}
//----------------------------------------------------------------------
func (service *NetSuiteService) DeleteInvoice(invoiceInternalId string) (bool,error) {
    recordType := RecordTypeInvoice
    recordInternalId := invoiceInternalId
    return service.deleteRecord(recordType, recordInternalId)
}
//----------------------------------------------------------------------
//----------------------------------------------------------------------
func (service *NetSuiteService) AddCreditMemo(creditmemo *CreditMemo) (string,error) {
    creditmemo.XsiType = "platformCore:CreditMemo" //this must be set, for the transaction to succeed
    request := CreditMemo_AddRequest{
        Record:creditmemo,
        XmlNSXSI: "http://www.w3.org/2001/XMLSchema-instance",
        XmlNSPC: "urn:customers_2022_1.transactions.webservices.netsuite.com",
        XmlNS1: "urn:core_2022_1.platform.webservices.netsuite.com", //needed for StringCustomFieldRef types
    }
    internalId:=""
    response, err := service.UpstreamService.CreditMemo_Add(&request)
    if (response!=nil && response.WriteResponse.Status.IsSuccess) {
        internalId = response.WriteResponse.BaseRef.InternalId
    } else if err != nil {
        err = distillSoapError(err)
    } else if response.WriteResponse.Status!=nil {
        if len(response.WriteResponse.Status.StatusDetail)>0 {
            err=errors.New(response.WriteResponse.Status.StatusDetail[0].Message)
        }
    }
    return internalId, err
}
//----------------------------------------------------------------------
func (service *NetSuiteService) GetCreditMemo(creditMemoInternalId string) (*CreditMemo,error) {
    recordType := RecordTypeCreditMemo
    baseRef := BaseRef{
            InternalId:creditMemoInternalId,
            Type:&recordType,
            XsiType:"platformCore:RecordRef", 
        }
    request := CreditMemo_GetRequest{
        BaseRef:&baseRef,
        XmlNSXSI: "http://www.w3.org/2001/XMLSchema-instance",
        XmlNSPC: "urn:core_2022_1.platform.webservices.netsuite.com",
    }
    var creditmemo *CreditMemo=nil
    response, err := service.UpstreamService.CreditMemo_Get(&request)
    if (response!=nil && response.ReadResponse.Status.IsSuccess) {
        creditmemo = response.ReadResponse.Record
    } else if err != nil {
        err = distillSoapError(err)
    } else if response.ReadResponse.Status!=nil {
        if len(response.ReadResponse.Status.StatusDetail)>0 {
            err=errors.New(response.ReadResponse.Status.StatusDetail[0].Message)
        }
    }
    return creditmemo, err
}
//----------------------------------------------------------------------
func (service *NetSuiteService) DeleteCreditMemo(creditMemoInternalId string) (bool,error) {
    recordType := RecordTypeCreditMemo
    recordInternalId := creditMemoInternalId
    return service.deleteRecord(recordType, recordInternalId)
}
//----------------------------------------------------------------------
//----------------------------------------------------------------------
func (service *NetSuiteService) AddCashSale(cashsale *CashSale) (string,error) {
    cashsale.XsiType = "platformCore:CashSale" //this must be set, for the transaction to succeed
    request := CashSale_AddRequest{
        Record:cashsale,
        XmlNSXSI: "http://www.w3.org/2001/XMLSchema-instance",
        XmlNSPC: "urn:sales_2022_1.transactions.webservices.netsuite.com",
    }
    internalId:=""
    response, err := service.UpstreamService.CashSale_Add(&request)
    if (response!=nil && response.WriteResponse.Status.IsSuccess) {
        internalId = response.WriteResponse.BaseRef.InternalId
    } else if err != nil {
        err = distillSoapError(err)
    } else if response.WriteResponse.Status!=nil {
        if len(response.WriteResponse.Status.StatusDetail)>0 {
            err=errors.New(response.WriteResponse.Status.StatusDetail[0].Message)
        }
    }
    return internalId, err
}
//----------------------------------------------------------------------
func (service *NetSuiteService) GetCashSale(creditMemoInternalId string) (*CashSale,error) {
    recordType := RecordTypeCashSale
    baseRef := BaseRef{
            InternalId:creditMemoInternalId,
            Type:&recordType,
            XsiType:"platformCore:RecordRef", 
        }
    request := CashSale_GetRequest{
        BaseRef:&baseRef,
        XmlNSXSI: "http://www.w3.org/2001/XMLSchema-instance",
        XmlNSPC: "urn:core_2022_1.platform.webservices.netsuite.com",
    }
    var cashsale *CashSale=nil
    response, err := service.UpstreamService.CashSale_Get(&request)
    if (response!=nil && response.ReadResponse.Status.IsSuccess) {
        cashsale = response.ReadResponse.Record
    } else if err != nil {
        err = distillSoapError(err)
    } else if response.ReadResponse.Status!=nil {
        if len(response.ReadResponse.Status.StatusDetail)>0 {
            err=errors.New(response.ReadResponse.Status.StatusDetail[0].Message)
        }
    }
    return cashsale, err
}
//----------------------------------------------------------------------
func (service *NetSuiteService) DeleteCashSale(cashSaleInternalId string) (bool,error) {
    recordType := RecordTypeCashSale
    recordInternalId := cashSaleInternalId
    return service.deleteRecord(recordType, recordInternalId)
}
//----------------------------------------------------------------------
//----------------------------------------------------------------------
func (service *NetSuiteService) deleteRecord(recordType RecordType, recordInternalId string) (bool,error) {
    baseRef := BaseRef{
            InternalId:recordInternalId,
            Type:&recordType,
            XsiType:"platformCore:RecordRef",
        }
    request := DeleteRequest{
        BaseRef:&baseRef,
        XmlNSXSI: "http://www.w3.org/2001/XMLSchema-instance",
        XmlNSPC:"urn:core_2022_1.platform.webservices.netsuite.com",
    }
    delResponse, err := service.UpstreamService.Delete(&request)
    if (delResponse!=nil && delResponse.DeleteReturn.Status.IsSuccess) {
        return true, nil
    } 
    log.Printf("deleteRecord Failure %e\n", err)
    return false, distillSoapError(err)
}
//----------------------------------------------------------------------
//----------------------------------------------------------------------
func (service *NetSuiteService) GetCategory(categoryInternalId string) (*CustomerCategory,error) {
    recordType := RecordTypeCustomerCategory
    baseRef := BaseRef{
            InternalId:categoryInternalId,
            Type:&recordType,
            XsiType:"platformCore:RecordRef",                
        }
    request := Category_GetRequest{
        BaseRef:&baseRef,
        XmlNSXSI: "http://www.w3.org/2001/XMLSchema-instance",
        XmlNSPC:"urn:core_2022_1.platform.webservices.netsuite.com",
    }
    var customerCategory *CustomerCategory=nil
    response, err := service.UpstreamService.Category_Get(&request)
    if (response!=nil && response.ReadResponse.Status.IsSuccess) {
        customerCategory = response.ReadResponse.Record
    }    
    return customerCategory, distillSoapError(err)
}
//----------------------------------------------------------------------
//----------------------------------------------------------------------
func (service *NetSuiteService) GetCurrencies() ([]*Currency,error) {
    currencytype := GetAllRecordTypeCurrency
    request := Currency_GetAllRequest{
        Record: &Currency_GetAllRecord{
            RecordType: &currencytype,
        },
    }
    var currencyList []*Currency
    getAllResponse, err := service.UpstreamService.Currency_GetAll(&request)
    if (getAllResponse!=nil && getAllResponse.GetAllResult.Status.IsSuccess) {
        for _,currency:= range(getAllResponse.GetAllResult.RecordList.Record) {
            currencyList = append(currencyList, currency)
        }
    }
    return currencyList, distillSoapError(err)
}
//----------------------------------------------------------------------
func (service *NetSuiteService) GetTerms() ([]*Term,error) {
    request:=SearchRequest{
        XmlNSXSI: "http://www.w3.org/2001/XMLSchema-instance",
        XmlNSPC:"urn:common_2022_1.platform.webservices.netsuite.com",        
        SearchRecord:&SearchRecord{
            XsiType:"platformCore:TermSearchBasic",
            //we are searching for [ALL], so basically no search criteria
        },
    }
    var termList []*Term
    searchResponse, err := service.UpstreamService.Terms_Search(&request)
    if (searchResponse!=nil && searchResponse.SearchResult.Status.IsSuccess) {
        for _,termRecord := range(searchResponse.SearchResult.RecordList.TermRecord) {
            //fmt.Printf("[Terms] InternalID:%s Name:%s\n", termRecord.InternalId, termRecord.Name)
            termList = append(termList, termRecord)
        }
    }
    return termList, distillSoapError(err)
}
//----------------------------------------------------------------------
func distillSoapError(Error error) error {
    //|| len(Error)<100
    if Error==nil {
        return Error
    }
    chopme:="HTTP Status 500: "
    xmlstr:=fmt.Sprintf("%s", Error)[len(chopme):]

    var soapError SoapErrorXML
    if err := xml.Unmarshal([]byte(xmlstr), &soapError); err != nil {
        log.Printf("Unable to parse error %s\n", err)
        return Error
    }
    return fmt.Errorf("%s", soapError.FaultString)
}
//----------------------------------------------------------------------
