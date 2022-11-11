//GENERATED via go2wsdl, (and modified for naming conflicts, and functionality)

package netsuite

import (
	//"context"
	"encoding/xml"
	//"github.com/hooklift/gowsdl/soap"
	//"time"
)

type RecordType string
const (
	RecordTypeCustomer RecordType = "customer"    
	RecordTypeInvoice RecordType = "invoice"
	RecordTypeSalesOrder RecordType = "salesOrder"
	RecordTypeCashRefund RecordType = "cashRefund"
	RecordTypeCashSale RecordType = "cashSale"
	RecordTypeCreditMemo RecordType = "creditMemo"
	RecordTypeCurrency RecordType = "currency"
	RecordTypeCustomerPayment RecordType = "customerPayment"
	RecordTypeCustomerCategory RecordType = "customerCategory"
	RecordTypeEmployee RecordType = "employee"
	RecordTypeNonInventorySaleItem RecordType = "nonInventorySaleItem"
)

type Status struct {
	XMLName xml.Name `xml:"status"`
	StatusDetail []*StatusDetail `xml:"statusDetail,omitempty" json:"statusDetail,omitempty"`
	IsSuccess bool `xml:"isSuccess,attr,omitempty" json:"isSuccess,omitempty"`
}

type StatusDetail struct {
	XMLName xml.Name `xml:"statusDetail"`
	Code *StatusDetailCodeType `xml:"code,omitempty" json:"code,omitempty"`
	Message string `xml:"message,omitempty" json:"message,omitempty"`
	AfterSubmitFailed bool `xml:"afterSubmitFailed,omitempty" json:"afterSubmitFailed,omitempty"`
	Type *StatusDetailType `xml:"type,attr,omitempty" json:"type,omitempty"`
}
type StatusDetailCodeType string
type StatusDetailType string

type RecordRef struct {
//	*BaseRef
	InternalId string `xml:"internalId,attr,omitempty" json:"internalId,omitempty"`
	ExternalId string `xml:"externalId,attr,omitempty" json:"externalId,omitempty"`
	Type *RecordType `xml:"type,attr,omitempty" json:"type,omitempty"`
}

type BaseRef struct {
	Name string `xml:"name,omitempty" json:"name,omitempty"`
	InternalId string `xml:"internalId,attr,omitempty" json:"internalId,omitempty"` 
	ExternalId string `xml:"externalId,attr,omitempty" json:"externalId,omitempty"` 
	Type *RecordType `xml:"type,attr,omitempty" json:"type,omitempty"`
    XsiType string `xml:"xsi:type,attr,omitempty"`
}

type Address struct {
	//*Record
	//InternalId string `xml:"internalId,omitempty" json:"internalId,omitempty"`
	Country string `xml:"country,omitempty" json:"country,omitempty"`
	Attention string `xml:"attention,omitempty" json:"attention,omitempty"`
	Addressee string `xml:"addressee,omitempty" json:"addressee,omitempty"`
	AddrPhone string `xml:"addrPhone,omitempty" json:"addrPhone,omitempty"`
	Addr1 string `xml:"addr1,omitempty" json:"addr1,omitempty"`
	Addr2 string `xml:"addr2,omitempty" json:"addr2,omitempty"`
	Addr3 string `xml:"addr3,omitempty" json:"addr3,omitempty"`
	City string `xml:"city,omitempty" json:"city,omitempty"`
	State string `xml:"state,omitempty" json:"state,omitempty"`
	Zip string `xml:"zip,omitempty" json:"zip,omitempty"`
	//AddrText string `xml:"addrText,omitempty" json:"addrText,omitempty"`
	//Override bool `xml:"override,omitempty" json:"override,omitempty"`
	CustomFieldList *CustomFieldList `xml:"customFieldList,omitempty" json:"customFieldList,omitempty"`
}
//------------------
type AddRequest struct {
	XMLName xml.Name `xml:"add"`
	XmlNSXSI   string   `xml:"xmlns:xsi,attr,omitempty"` 
	XmlNSPC   string   `xml:"xmlns:platformCore,attr,omitempty"` 
	Record *Customer `xml:"record,omitempty" json:"record,omitempty"`
}

type AddResponse struct {
	XMLName xml.Name `xml:"addResponse"`
	WriteResponse *WriteResponse `xml:"writeResponse,omitempty" json:"writeResponse,omitempty"`
}

type WriteResponse struct {
	XMLName xml.Name `xml:"writeResponse"`
	Status *Status `xml:"status,omitempty" json:"status,omitempty"`
	BaseRef *BaseRef `xml:"baseRef,omitempty" json:"baseRef,omitempty"`
}

type SoapErrorXML struct {
	XMLName xml.Name `xml:"Envelope"`
	FaultCode string `xml:"Body>Fault>faultcode,omitempty"`
	FaultString string `xml:"Body>Fault>faultstring,omitempty"`
}
/*
<?xml version="1.0" encoding="utf-8"?>
 <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
 <soapenv:Body>
   <soapenv:Fault>
     <faultcode>soapenv:Server.userException</faultcode>
     <faultstring>org.xml.sax.SAXException: Invalid type: {urn:relationships_2022_1.lists.webservices.com}Customer</faultstring>
     <detail><ns1:hostname xmlns:ns1="http://xml.apache.org/axis/">partners002</ns1:hostname></detail>
   </soapenv:Fault>
 </soapenv:Body>
 </soapenv:Envelope>*/


//------------------
type GetRequest struct {
	XMLName xml.Name `xml:"get"`
	XmlNSXSI   string   `xml:"xmlns:xsi,attr,omitempty"` 
	XmlNSPC   string   `xml:"xmlns:platformCore,attr,omitempty"` 
	BaseRef *BaseRef `xml:"baseRef,omitempty" json:"baseRef,omitempty"`
}

type GetResponse struct {
	//XMLName xml.Name `xml:"urn:messages_2022_1.platform.webservices.netsuite.com getResponse"`
	XMLName xml.Name `xml:"getResponse"`
	ReadResponse *ReadResponse `xml:"readResponse,omitempty" json:"readResponse,omitempty"`
}

type ReadResponse struct {
	XMLName xml.Name `xml:"readResponse"`
	Status *Status `xml:"status,omitempty" json:"status,omitempty"`
	Record *Customer `xml:"record,omitempty" json:"record,omitempty"`
}
//------------------
type DeleteRequest struct {
	XMLName xml.Name `xml:"delete"`
	BaseRef *BaseRef `xml:"baseRef,omitempty" json:"baseRef,omitempty"`
	DeletionReason *DeletionReason `xml:"deletionReason,omitempty" json:"deletionReason,omitempty"`
	XmlNSXSI   string   `xml:"xmlns:xsi,attr,omitempty"` 
	XmlNSPC   string   `xml:"xmlns:platformCore,attr,omitempty"` 
}

type DeletionReason struct {
	XMLName xml.Name `xml:"deletionReason"`
	DeletionReasonCode *RecordRef `xml:"deletionReasonCode,omitempty" json:"deletionReasonCode,omitempty"`
	DeletionReasonMemo string `xml:"deletionReasonMemo,omitempty" json:"deletionReasonMemo,omitempty"`
}

type DeleteResponse struct {
	XMLName xml.Name `xml:"deleteResponse"`
	DeleteReturn *DeleteReturn `xml:"deleteReturn,omitempty" json:"writeResponse,omitempty"`
	//WriteResponse *WriteResponse `xml:"writeResponse,omitempty" json:"writeResponse,omitempty"`
}
/*
//<deleteResponse xmlns="">
//  <deleteReturn xsi:type="WriteResponse" xmlns:platformMsgs="urn:messages_2022_1.platform.webservices.netsuite.com">
//    <platformCore:status isSuccess="true" xmlns:platformCore="urn:core_2022_1.platform.webservices.netsuite.com"/>
//    <platformMsgs:baseRef internalId="15822105" type="salesOrder" xsi:type="platformCore:RecordRef" xmlns:platformCore="urn:core_2022_1.platform.webservices.netsuite.com"/>
//  </deleteReturn>
//</deleteResponse>
//I manually added this struct DeleteReturn, because the soap/xml parser wasn't automatically promoting 
//[<deleteReturn xsi:type="WriteResponse"/>] to [<WriteResponse/>]
 * */
type DeleteReturn struct {
	XMLName xml.Name `xml:"deleteReturn"`
	Status *Status `xml:"status,omitempty" json:"status,omitempty"`
	BaseRef *BaseRef `xml:"baseRef,omitempty" json:"baseRef,omitempty"`
}

//----------------------------------------------------------------------
