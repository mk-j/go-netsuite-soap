//GENERATED via go2wsdl, (and modified for naming conflicts, and functionality)

package netsuite

import (
	"encoding/xml"
	//"github.com/hooklift/gowsdl/soap"
)

const GetAllRecordTypeCurrency GetAllRecordType = "currency"
type GetAllRecordType string

type Currency_GetAllRequest struct {
	XMLName xml.Name `xml:"getAll"`
	Record *Currency_GetAllRecord `xml:"record,omitempty" json:"record,omitempty"`
}

type Currency_GetAllRecord struct {
	XMLName xml.Name `xml:"record"`
	RecordType *GetAllRecordType `xml:"recordType,attr,omitempty" json:"recordType,omitempty"`
}

type Currency_GetAllResponse struct {
	XMLName xml.Name `xml:"getAllResponse"`
	GetAllResult *Currency_GetAllResult `xml:"getAllResult,omitempty" json:"getAllResult,omitempty"`
}

type Currency_GetAllResult struct {
	XMLName xml.Name `xml:"getAllResult"`
	Status *Status `xml:"status,omitempty" json:"status,omitempty"`
	TotalRecords int32 `xml:"totalRecords,omitempty" json:"totalRecords,omitempty"`
	RecordList *Currency_RecordList `xml:"recordList,omitempty" json:"recordList,omitempty"`
}

type Currency_RecordList struct {
	XMLName xml.Name `xml:"recordList"`
	//Record []*CurrencyRecord `xml:"record,omitempty" json:"record,omitempty"` 
	Record []*Currency `xml:"record,omitempty" json:"record,omitempty"` 
}

type CurrencyRecord struct {
	XMLName xml.Name `xml:"record"`
	Contents string `xml:",innerxml"`
	//NullFieldList *NullField `xml:"nullFieldList,omitempty" json:"nullFieldList,omitempty"`
}

type Currency struct {
	Name string `xml:"name,omitempty" json:"name,omitempty"`
	Symbol string `xml:"symbol,omitempty" json:"symbol,omitempty"`
	IsBaseCurrency bool `xml:"isBaseCurrency,omitempty" json:"isBaseCurrency,omitempty"`
	IsInactive bool `xml:"isInactive,omitempty" json:"isInactive,omitempty"`
	OverrideCurrencyFormat bool `xml:"overrideCurrencyFormat,omitempty" json:"overrideCurrencyFormat,omitempty"`
	DisplaySymbol string `xml:"displaySymbol,omitempty" json:"displaySymbol,omitempty"`
	SymbolPlacement *CurrencySymbolPlacement `xml:"symbolPlacement,omitempty" json:"symbolPlacement,omitempty"`
	Locale *CurrencyLocale `xml:"locale,omitempty" json:"locale,omitempty"`
	FormatSample string `xml:"formatSample,omitempty" json:"formatSample,omitempty"`
	ExchangeRate float64 `xml:"exchangeRate,omitempty" json:"exchangeRate,omitempty"`
	FxRateUpdateTimezone *CurrencyFxRateUpdateTimezone `xml:"fxRateUpdateTimezone,omitempty" json:"fxRateUpdateTimezone,omitempty"`
	InclInFxRateUpdates bool `xml:"inclInFxRateUpdates,omitempty" json:"inclInFxRateUpdates,omitempty"`
	CurrencyPrecision *CurrencyCurrencyPrecision `xml:"currencyPrecision,omitempty" json:"currencyPrecision,omitempty"`
	InternalId string `xml:"internalId,attr,omitempty" json:"internalId,omitempty"`
	ExternalId string `xml:"externalId,attr,omitempty" json:"externalId,omitempty"`
}
type CurrencyCurrencyPrecision string
type CurrencySymbolPlacement string
type CurrencyLocale string
type CurrencyFxRateUpdateTimezone string

