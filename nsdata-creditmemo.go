//GENERATED via go2wsdl, (and modified for naming conflicts, and functionality)
package netsuite
import (
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
)
type CreditMemo_AddRequest struct {
	XMLName xml.Name `xml:"add"`
	XmlNSXSI   string   `xml:"xmlns:xsi,attr,omitempty"` 
	XmlNSPC   string   `xml:"xmlns:platformCore,attr,omitempty"` 
	XmlNS1   string   `xml:"xmlns:ns1,attr,omitempty"` 
	Record *CreditMemo `xml:"record,omitempty" json:"record,omitempty"`
}
type CreditMemo_AddResponse struct {
	XMLName xml.Name `xml:"addResponse"`
	WriteResponse *CreditMemo_WriteResponse `xml:"writeResponse,omitempty" json:"writeResponse,omitempty"`
}
type CreditMemo_WriteResponse struct {
	XMLName xml.Name `xml:"writeResponse"`
	Status *Status `xml:"status,omitempty" json:"status,omitempty"`
	BaseRef *BaseRef `xml:"baseRef,omitempty" json:"baseRef,omitempty"`
}
//----------------------------------------------------
type CreditMemo_GetRequest struct {
	XMLName xml.Name `xml:"get"`
	XmlNSXSI   string   `xml:"xmlns:xsi,attr,omitempty"` 
	XmlNSPC   string   `xml:"xmlns:platformCore,attr,omitempty"` 
	BaseRef *BaseRef `xml:"baseRef,omitempty" json:"baseRef,omitempty"`
}
type CreditMemo_GetResponse struct {
	//XMLName xml.Name `xml:"urn:messages_2022_1.platform.webservices.netsuite.com getResponse"`
	XMLName xml.Name `xml:"getResponse"`
	ReadResponse *CreditMemo_ReadResponse `xml:"readResponse,omitempty" json:"readResponse,omitempty"`
}
//----------------------------------------------------
type CreditMemo_ReadResponse struct {
	XMLName xml.Name `xml:"readResponse"`
	Status *Status `xml:"status,omitempty" json:"status,omitempty"`
	Record *CreditMemo `xml:"record,omitempty" json:"record,omitempty"`
}
type CreditMemo struct {
	XsiType string `xml:"xsi:type,attr,omitempty"  json:"-"`
	CreatedDate soap.XSDDateTime `xml:"createdDate,omitempty" json:"createdDate,omitempty"`
	LastModifiedDate soap.XSDDateTime `xml:"lastModifiedDate,omitempty" json:"lastModifiedDate,omitempty"`
	//Nexus *RecordRef `xml:"nexus,omitempty" json:"nexus,omitempty"`
	//SubsidiaryTaxRegNum *RecordRef `xml:"subsidiaryTaxRegNum,omitempty" json:"subsidiaryTaxRegNum,omitempty"`
	//TaxRegOverride bool `xml:"taxRegOverride,omitempty" json:"taxRegOverride,omitempty"`
	//TaxDetailsOverride bool `xml:"taxDetailsOverride,omitempty" json:"taxDetailsOverride,omitempty"`
	CustomForm *RecordRef `xml:"customForm,omitempty" json:"customForm,omitempty"`
	Currency *RecordRef `xml:"currency,omitempty" json:"currency,omitempty"`
	Entity *RecordRef `xml:"entity,omitempty" json:"entity,omitempty"`
	//VatRegNum string `xml:"vatRegNum,omitempty" json:"vatRegNum,omitempty"`
	TranDate soap.XSDDateTime `xml:"tranDate,omitempty" json:"tranDate,omitempty"`
	TranId string `xml:"tranId,omitempty" json:"tranId,omitempty"`
	//EntityTaxRegNum *RecordRef `xml:"entityTaxRegNum,omitempty" json:"entityTaxRegNum,omitempty"`
	//TaxPointDate soap.XSDDateTime `xml:"taxPointDate,omitempty" json:"taxPointDate,omitempty"`
	CreatedFrom *RecordRef `xml:"createdFrom,omitempty" json:"createdFrom,omitempty"`
	//PostingPeriod *RecordRef `xml:"postingPeriod,omitempty" json:"postingPeriod,omitempty"`
	Department *RecordRef `xml:"department,omitempty" json:"department,omitempty"`
	Class *RecordRef `xml:"class,omitempty" json:"class,omitempty"`
	Location *RecordRef `xml:"location,omitempty" json:"location,omitempty"`
	Subsidiary *RecordRef `xml:"subsidiary,omitempty" json:"subsidiary,omitempty"`
	//Job *RecordRef `xml:"job,omitempty" json:"job,omitempty"`
	SalesRep *RecordRef `xml:"salesRep,omitempty" json:"salesRep,omitempty"`
	//Partner *RecordRef `xml:"partner,omitempty" json:"partner,omitempty"`
	//ContribPct string `xml:"contribPct,omitempty" json:"contribPct,omitempty"`
	OtherRefNum string `xml:"otherRefNum,omitempty" json:"otherRefNum,omitempty"`
	Memo string `xml:"memo,omitempty" json:"memo,omitempty"`
	//ExcludeCommission bool `xml:"excludeCommission,omitempty" json:"excludeCommission,omitempty"`
	//LeadSource *RecordRef `xml:"leadSource,omitempty" json:"leadSource,omitempty"`
	//Balance float64 `xml:"balance,omitempty" json:"balance,omitempty"`
	//Account *RecordRef `xml:"account,omitempty" json:"account,omitempty"`
	//ExchangeRate float64 `xml:"exchangeRate,omitempty" json:"exchangeRate,omitempty"`
	//OnCreditHold string `xml:"onCreditHold,omitempty" json:"onCreditHold,omitempty"`
	AmountPaid float64 `xml:"amountPaid,omitempty" json:"amountPaid,omitempty"`
	SalesEffectiveDate soap.XSDDateTime `xml:"salesEffectiveDate,omitempty" json:"salesEffectiveDate,omitempty"`
	//TotalCostEstimate float64 `xml:"totalCostEstimate,omitempty" json:"totalCostEstimate,omitempty"`
	//EstGrossProfit float64 `xml:"estGrossProfit,omitempty" json:"estGrossProfit,omitempty"`
	//EstGrossProfitPercent float64 `xml:"estGrossProfitPercent,omitempty" json:"estGrossProfitPercent,omitempty"`
	//CurrencyName string `xml:"currencyName,omitempty" json:"currencyName,omitempty"`
	//PromoCode *RecordRef `xml:"promoCode,omitempty" json:"promoCode,omitempty"`
	AmountRemaining float64 `xml:"amountRemaining,omitempty" json:"amountRemaining,omitempty"`
	//DiscountItem *RecordRef `xml:"discountItem,omitempty" json:"discountItem,omitempty"`
	Source string `xml:"source,omitempty" json:"source,omitempty"`
	//DiscountRate string `xml:"discountRate,omitempty" json:"discountRate,omitempty"`
	//IsTaxable bool `xml:"isTaxable,omitempty" json:"isTaxable,omitempty"`
	//TaxItem *RecordRef `xml:"taxItem,omitempty" json:"taxItem,omitempty"`
	//TaxRate float64 `xml:"taxRate,omitempty" json:"taxRate,omitempty"`
	//Unapplied float64 `xml:"unapplied,omitempty" json:"unapplied,omitempty"`
	//AutoApply bool `xml:"autoApply,omitempty" json:"autoApply,omitempty"`
	//Applied float64 `xml:"applied,omitempty" json:"applied,omitempty"`
	//ToBePrinted bool `xml:"toBePrinted,omitempty" json:"toBePrinted,omitempty"`
	//ToBeEmailed bool `xml:"toBeEmailed,omitempty" json:"toBeEmailed,omitempty"`
	//Email string `xml:"email,omitempty" json:"email,omitempty"`
	//ToBeFaxed bool `xml:"toBeFaxed,omitempty" json:"toBeFaxed,omitempty"`
	//Fax string `xml:"fax,omitempty" json:"fax,omitempty"`
	//MessageSel *RecordRef `xml:"messageSel,omitempty" json:"messageSel,omitempty"`
	//Message string `xml:"message,omitempty" json:"message,omitempty"`
	BillingAddress *Address `xml:"billingAddress,omitempty" json:"billingAddress,omitempty"`
	//BillAddressList *RecordRef `xml:"billAddressList,omitempty" json:"billAddressList,omitempty"`
	//ShipMethod *RecordRef `xml:"shipMethod,omitempty" json:"shipMethod,omitempty"`
	//ShippingCost float64 `xml:"shippingCost,omitempty" json:"shippingCost,omitempty"`
	//ShippingTax1Rate float64 `xml:"shippingTax1Rate,omitempty" json:"shippingTax1Rate,omitempty"`
	//ShippingTaxCode *RecordRef `xml:"shippingTaxCode,omitempty" json:"shippingTaxCode,omitempty"`
	//HandlingTaxCode *RecordRef `xml:"handlingTaxCode,omitempty" json:"handlingTaxCode,omitempty"`
	//ShippingTax2Rate string `xml:"shippingTax2Rate,omitempty" json:"shippingTax2Rate,omitempty"`
	//HandlingTax1Rate float64 `xml:"handlingTax1Rate,omitempty" json:"handlingTax1Rate,omitempty"`
	//HandlingTax2Rate string `xml:"handlingTax2Rate,omitempty" json:"handlingTax2Rate,omitempty"`
	//HandlingCost float64 `xml:"handlingCost,omitempty" json:"handlingCost,omitempty"`
	SubTotal float64 `xml:"subTotal,omitempty" json:"subTotal,omitempty"`
	//DiscountTotal float64 `xml:"discountTotal,omitempty" json:"discountTotal,omitempty"`
	//RevenueStatus *RevenueStatus `xml:"revenueStatus,omitempty" json:"revenueStatus,omitempty"`
	//RecognizedRevenue float64 `xml:"recognizedRevenue,omitempty" json:"recognizedRevenue,omitempty"`
	//DeferredRevenue float64 `xml:"deferredRevenue,omitempty" json:"deferredRevenue,omitempty"`
	//RevRecOnRevCommitment bool `xml:"revRecOnRevCommitment,omitempty" json:"revRecOnRevCommitment,omitempty"`
	//TaxTotal float64 `xml:"taxTotal,omitempty" json:"taxTotal,omitempty"`
	//Tax2Total float64 `xml:"tax2Total,omitempty" json:"tax2Total,omitempty"`
	//AltShippingCost float64 `xml:"altShippingCost,omitempty" json:"altShippingCost,omitempty"`
	//AltHandlingCost float64 `xml:"altHandlingCost,omitempty" json:"altHandlingCost,omitempty"`
	//IsMultiShipTo bool `xml:"isMultiShipTo,omitempty" json:"isMultiShipTo,omitempty"`
	//Total float64 `xml:"total,omitempty" json:"total,omitempty"`
	//SalesGroup *RecordRef `xml:"salesGroup,omitempty" json:"salesGroup,omitempty"`
	//SyncSalesTeams bool `xml:"syncSalesTeams,omitempty" json:"syncSalesTeams,omitempty"`
	Status string `xml:"status,omitempty" json:"status,omitempty"`
	//GiftCert *RecordRef `xml:"giftCert,omitempty" json:"giftCert,omitempty"`
	//GiftCertTotal float64 `xml:"giftCertTotal,omitempty" json:"giftCertTotal,omitempty"`
	//GiftCertApplied float64 `xml:"giftCertApplied,omitempty" json:"giftCertApplied,omitempty"`
	//GiftCertAvailable float64 `xml:"giftCertAvailable,omitempty" json:"giftCertAvailable,omitempty"`
	//TranIsVsoeBundle bool `xml:"tranIsVsoeBundle,omitempty" json:"tranIsVsoeBundle,omitempty"`
	//VsoeAutoCalc bool `xml:"vsoeAutoCalc,omitempty" json:"vsoeAutoCalc,omitempty"`
	//SyncPartnerTeams bool `xml:"syncPartnerTeams,omitempty" json:"syncPartnerTeams,omitempty"`
	//SalesTeamList *CreditMemoSalesTeamList `xml:"salesTeamList,omitempty" json:"salesTeamList,omitempty"`
	ItemList *CreditMemoItemList `xml:"itemList,omitempty" json:"itemList,omitempty"`
	//AccountingBookDetailList *AccountingBookDetailList `xml:"accountingBookDetailList,omitempty" json:"accountingBookDetailList,omitempty"`
	//PartnersList *CreditMemoPartnersList `xml:"partnersList,omitempty" json:"partnersList,omitempty"`
	ApplyList *CreditMemoApplyList `xml:"applyList,omitempty" json:"applyList,omitempty"`
	//TaxDetailsList *TaxDetailsList `xml:"taxDetailsList,omitempty" json:"taxDetailsList,omitempty"`
	CustomFieldList *CustomFieldList `xml:"customFieldList,omitempty" json:"customFieldList,omitempty"`
	InternalId string `xml:"internalId,attr,omitempty" json:"internalId,omitempty"`
	ExternalId string `xml:"externalId,attr,omitempty" json:"externalId,omitempty"`
}
type CreditMemoItemList struct {
	XMLName xml.Name `xml:"itemList"`
	Item []*CreditMemoItem `xml:"item,omitempty" json:"item,omitempty"`
	ReplaceAll bool `xml:"replaceAll,attr,omitempty" json:"replaceAll,omitempty"`
}
type CreditMemoItem struct {
	XMLName xml.Name `xml:"item"`
	Job *RecordRef `xml:"job,omitempty" json:"job,omitempty"`
	Item *RecordRef `xml:"item,omitempty" json:"item,omitempty"`
	OrderLine int64 `xml:"orderLine,omitempty" json:"orderLine,omitempty"`
	Line int64 `xml:"line,omitempty" json:"line,omitempty"`
	Quantity float64 `xml:"quantity,omitempty" json:"quantity,omitempty"`
	Description string `xml:"description,omitempty" json:"description,omitempty"`
	//BinNumbers string `xml:"binNumbers,omitempty" json:"binNumbers,omitempty"`
	Price *RecordRef `xml:"price,omitempty" json:"price,omitempty"`
	Rate string `xml:"rate,omitempty" json:"rate,omitempty"`
	Amount float64 `xml:"amount,omitempty" json:"amount,omitempty"`
	//IsTaxable bool `xml:"isTaxable,omitempty" json:"isTaxable,omitempty"`
	//Options *CustomFieldList `xml:"options,omitempty" json:"options,omitempty"`
	ShipAddress *RecordRef `xml:"shipAddress,omitempty" json:"shipAddress,omitempty"`
	//ShipCarrier *ShippingCarrier `xml:"shipCarrier,omitempty" json:"shipCarrier,omitempty"`
	ShipMethod *RecordRef `xml:"shipMethod,omitempty" json:"shipMethod,omitempty"`
	TaxCode *RecordRef `xml:"taxCode,omitempty" json:"taxCode,omitempty"`
	TaxRate1 float64 `xml:"taxRate1,omitempty" json:"taxRate1,omitempty"`
	TaxRate2 float64 `xml:"taxRate2,omitempty" json:"taxRate2,omitempty"`
	Tax1Amt float64 `xml:"tax1Amt,omitempty" json:"tax1Amt,omitempty"`
	GrossAmt float64 `xml:"grossAmt,omitempty" json:"grossAmt,omitempty"`
	Department *RecordRef `xml:"department,omitempty" json:"department,omitempty"`
	Class *RecordRef `xml:"class,omitempty" json:"class,omitempty"`
	Location *RecordRef `xml:"location,omitempty" json:"location,omitempty"`
	//CostEstimateType *ItemCostEstimateType `xml:"costEstimateType,omitempty" json:"costEstimateType,omitempty"`
	//CostEstimate float64 `xml:"costEstimate,omitempty" json:"costEstimate,omitempty"`
	//TaxDetailsReference string `xml:"taxDetailsReference,omitempty" json:"taxDetailsReference,omitempty"`
	//RevRecSchedule *RecordRef `xml:"revRecSchedule,omitempty" json:"revRecSchedule,omitempty"`
	//RevRecStartDate soap.XSDDateTime `xml:"revRecStartDate,omitempty" json:"revRecStartDate,omitempty"`
	//RevRecTermInMonths int64 `xml:"revRecTermInMonths,omitempty" json:"revRecTermInMonths,omitempty"`
	//RevRecEndDate soap.XSDDateTime `xml:"revRecEndDate,omitempty" json:"revRecEndDate,omitempty"`
	Units *RecordRef `xml:"units,omitempty" json:"units,omitempty"`
	//InventoryDetail *InventoryDetail `xml:"inventoryDetail,omitempty" json:"inventoryDetail,omitempty"`
	//SerialNumbers string `xml:"serialNumbers,omitempty" json:"serialNumbers,omitempty"`
	//DeferRevRec bool `xml:"deferRevRec,omitempty" json:"deferRevRec,omitempty"`
	//GiftCertFrom string `xml:"giftCertFrom,omitempty" json:"giftCertFrom,omitempty"`
	//GiftCertRecipientName string `xml:"giftCertRecipientName,omitempty" json:"giftCertRecipientName,omitempty"`
	//GiftCertRecipientEmail string `xml:"giftCertRecipientEmail,omitempty" json:"giftCertRecipientEmail,omitempty"`
	//GiftCertMessage string `xml:"giftCertMessage,omitempty" json:"giftCertMessage,omitempty"`
	//TaxAmount float64 `xml:"taxAmount,omitempty" json:"taxAmount,omitempty"`
	//GiftCertNumber string `xml:"giftCertNumber,omitempty" json:"giftCertNumber,omitempty"`
	//VsoeSopGroup *VsoeSopGroup `xml:"vsoeSopGroup,omitempty" json:"vsoeSopGroup,omitempty"`
	//VsoeIsEstimate bool `xml:"vsoeIsEstimate,omitempty" json:"vsoeIsEstimate,omitempty"`
	//VsoePrice float64 `xml:"vsoePrice,omitempty" json:"vsoePrice,omitempty"`
	//VsoeAmount float64 `xml:"vsoeAmount,omitempty" json:"vsoeAmount,omitempty"`
	//VsoeAllocation float64 `xml:"vsoeAllocation,omitempty" json:"vsoeAllocation,omitempty"`
	//VsoeDeferral *VsoeDeferral `xml:"vsoeDeferral,omitempty" json:"vsoeDeferral,omitempty"`
	//VsoePermitDiscount *VsoePermitDiscount `xml:"vsoePermitDiscount,omitempty" json:"vsoePermitDiscount,omitempty"`
	//VsoeDelivered bool `xml:"vsoeDelivered,omitempty" json:"vsoeDelivered,omitempty"`
	//CatchUpPeriod *RecordRef `xml:"catchUpPeriod,omitempty" json:"catchUpPeriod,omitempty"`
	ChargeType *RecordRef `xml:"chargeType,omitempty" json:"chargeType,omitempty"`
	//SubscriptionLine *RecordRef `xml:"subscriptionLine,omitempty" json:"subscriptionLine,omitempty"`
	//ChargesList *RecordRefList `xml:"chargesList,omitempty" json:"chargesList,omitempty"`
	CustomFieldList *CustomFieldList `xml:"customFieldList,omitempty" json:"customFieldList,omitempty"`
}
type CreditMemoApply struct {
	XMLName xml.Name `xml:"apply"`
	Apply bool `xml:"apply,omitempty" json:"apply,omitempty"`
	Doc int64 `xml:"doc,omitempty" json:"doc,omitempty"`
	ApplyDate soap.XSDDateTime `xml:"applyDate,omitempty" json:"applyDate,omitempty"`
	Job string `xml:"job,omitempty" json:"job,omitempty"`
	Type_ string `xml:"type,omitempty" json:"type,omitempty"`
	RefNum string `xml:"refNum,omitempty" json:"refNum,omitempty"`
	Total float64 `xml:"total,omitempty" json:"total,omitempty"`
	Due float64 `xml:"due,omitempty" json:"due,omitempty"`
	Currency string `xml:"currency,omitempty" json:"currency,omitempty"`
	Amount float64 `xml:"amount,omitempty" json:"amount,omitempty"`
	Line int64 `xml:"line,omitempty" json:"line,omitempty"`
}
type CreditMemoApplyList struct {
	XMLName xml.Name `xml:"applyList"`
	Apply []*CreditMemoApply `xml:"apply,omitempty" json:"apply,omitempty"`
	ReplaceAll bool `xml:"replaceAll,attr,omitempty" json:"replaceAll,omitempty"`
}
