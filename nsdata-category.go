//GENERATED via go2wsdl, (and modified for naming conflicts, and functionality)

package netsuite

import (
	"encoding/xml"
	//"github.com/hooklift/gowsdl/soap"
)

type Category_GetRequest struct {
	XMLName xml.Name `xml:"get"`
	XmlNSXSI   string   `xml:"xmlns:xsi,attr,omitempty"` //ADDED
	XmlNSPC   string   `xml:"xmlns:platformCore,attr,omitempty"` //ADDED
	BaseRef *BaseRef `xml:"baseRef,omitempty" json:"baseRef,omitempty"`
}

type Category_GetResponse struct {
	//XMLName xml.Name `xml:"urn:messages_2022_1.platform.webservices.netsuite.com getResponse"`
	XMLName xml.Name `xml:"getResponse"`
	ReadResponse *Category_ReadResponse `xml:"readResponse,omitempty" json:"readResponse,omitempty"`
}

type Category_ReadResponse struct {
	XMLName xml.Name `xml:"readResponse"`
	Status *Status `xml:"status,omitempty" json:"status,omitempty"`
	Record *CustomerCategory `xml:"record,omitempty" json:"record,omitempty"` // ADDED
}

type CustomerCategory struct {
	Name string `xml:"name,omitempty" json:"name,omitempty"`
	IsInactive bool `xml:"isInactive,omitempty" json:"isInactive,omitempty"`
	InternalId string `xml:"internalId,attr,omitempty" json:"internalId,omitempty"`
	ExternalId string `xml:"externalId,attr,omitempty" json:"externalId,omitempty"`
}



















