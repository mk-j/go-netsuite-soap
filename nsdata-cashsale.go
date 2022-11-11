//GENERATED via go2wsdl, (and modified for naming conflicts, and functionality)

package netsuite

import (
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
)
type CashSale_AddRequest struct {
	XMLName xml.Name `xml:"add"`
	XmlNSXSI   string   `xml:"xmlns:xsi,attr,omitempty"` 
	XmlNSPC   string   `xml:"xmlns:platformCore,attr,omitempty"` 
	XmlNS1   string   `xml:"xmlns:ns1,attr,omitempty"` 
	Record *CashSale `xml:"record,omitempty" json:"record,omitempty"`
}
type CashSale_AddResponse struct {
	XMLName xml.Name `xml:"addResponse"`
	WriteResponse *CashSale_WriteResponse `xml:"writeResponse,omitempty" json:"writeResponse,omitempty"`
}
type CashSale_WriteResponse struct {
	XMLName xml.Name `xml:"writeResponse"`
	Status *Status `xml:"status,omitempty" json:"status,omitempty"`
	BaseRef *BaseRef `xml:"baseRef,omitempty" json:"baseRef,omitempty"`
}
//----------------------------------------------------
type CashSale_GetRequest struct {
	XMLName xml.Name `xml:"get"`
	XmlNSXSI   string   `xml:"xmlns:xsi,attr,omitempty"` 
	XmlNSPC   string   `xml:"xmlns:platformCore,attr,omitempty"` 
	BaseRef *BaseRef `xml:"baseRef,omitempty" json:"baseRef,omitempty"`
}
type CashSale_GetResponse struct {
	//XMLName xml.Name `xml:"urn:messages_2022_1.platform.webservices.netsuite.com getResponse"`
	XMLName xml.Name `xml:"getResponse"`
	ReadResponse *CashSale_ReadResponse `xml:"readResponse,omitempty" json:"readResponse,omitempty"`
}
//----------------------------------------------------
type CashSale_ReadResponse struct {
	XMLName xml.Name `xml:"readResponse"`
	Status *Status `xml:"status,omitempty" json:"status,omitempty"`
	Record *CashSale `xml:"record,omitempty" json:"record,omitempty"`
}

type CashSale struct {
	XsiType string `xml:"xsi:type,attr,omitempty"  json:"-"`
	CreatedDate soap.XSDDateTime `xml:"createdDate,omitempty" json:"createdDate,omitempty"`
	LastModifiedDate soap.XSDDateTime `xml:"lastModifiedDate,omitempty" json:"lastModifiedDate,omitempty"`
	Nexus *RecordRef `xml:"nexus,omitempty" json:"nexus,omitempty"`
	//SubsidiaryTaxRegNum *RecordRef `xml:"subsidiaryTaxRegNum,omitempty" json:"subsidiaryTaxRegNum,omitempty"`
	//TaxRegOverride bool `xml:"taxRegOverride,omitempty" json:"taxRegOverride,omitempty"`
	//TaxDetailsOverride bool `xml:"taxDetailsOverride,omitempty" json:"taxDetailsOverride,omitempty"`
	CustomForm *RecordRef `xml:"customForm,omitempty" json:"customForm,omitempty"`
	Entity *RecordRef `xml:"entity,omitempty" json:"entity,omitempty"`
	//BillingAccount *RecordRef `xml:"billingAccount,omitempty" json:"billingAccount,omitempty"`
	//RecurringBill bool `xml:"recurringBill,omitempty" json:"recurringBill,omitempty"`
	TranDate soap.XSDDateTime `xml:"tranDate,omitempty" json:"tranDate,omitempty"`
	TranId string `xml:"tranId,omitempty" json:"tranId,omitempty"`
	//EntityTaxRegNum *RecordRef `xml:"entityTaxRegNum,omitempty" json:"entityTaxRegNum,omitempty"`
	//TaxPointDate soap.XSDDateTime `xml:"taxPointDate,omitempty" json:"taxPointDate,omitempty"`
	Source string `xml:"source,omitempty" json:"source,omitempty"`
	//PostingPeriod *RecordRef `xml:"postingPeriod,omitempty" json:"postingPeriod,omitempty"`
	CreatedFrom *RecordRef `xml:"createdFrom,omitempty" json:"createdFrom,omitempty"`
	//Opportunity *RecordRef `xml:"opportunity,omitempty" json:"opportunity,omitempty"`
	Department *RecordRef `xml:"department,omitempty" json:"department,omitempty"`
	Class *RecordRef `xml:"class,omitempty" json:"class,omitempty"`
	Location *RecordRef `xml:"location,omitempty" json:"location,omitempty"`
	Subsidiary *RecordRef `xml:"subsidiary,omitempty" json:"subsidiary,omitempty"`
	SalesRep *RecordRef `xml:"salesRep,omitempty" json:"salesRep,omitempty"`
	//ContribPct string `xml:"contribPct,omitempty" json:"contribPct,omitempty"`
	//Partner *RecordRef `xml:"partner,omitempty" json:"partner,omitempty"`
	//LeadSource *RecordRef `xml:"leadSource,omitempty" json:"leadSource,omitempty"`
	StartDate soap.XSDDateTime `xml:"startDate,omitempty" json:"startDate,omitempty"`
	EndDate soap.XSDDateTime `xml:"endDate,omitempty" json:"endDate,omitempty"`
	OtherRefNum string `xml:"otherRefNum,omitempty" json:"otherRefNum,omitempty"`
	Memo string `xml:"memo,omitempty" json:"memo,omitempty"`
	SalesEffectiveDate soap.XSDDateTime `xml:"salesEffectiveDate,omitempty" json:"salesEffectiveDate,omitempty"`
	//ExcludeCommission bool `xml:"excludeCommission,omitempty" json:"excludeCommission,omitempty"`
	//RevRecSchedule *RecordRef `xml:"revRecSchedule,omitempty" json:"revRecSchedule,omitempty"`
	//UndepFunds bool `xml:"undepFunds,omitempty" json:"undepFunds,omitempty"`
	//CanHaveStackable bool `xml:"canHaveStackable,omitempty" json:"canHaveStackable,omitempty"`
	Currency *RecordRef `xml:"currency,omitempty" json:"currency,omitempty"`
	//Account *RecordRef `xml:"account,omitempty" json:"account,omitempty"`
	//RevRecStartDate soap.XSDDateTime `xml:"revRecStartDate,omitempty" json:"revRecStartDate,omitempty"`
	//RevRecEndDate soap.XSDDateTime `xml:"revRecEndDate,omitempty" json:"revRecEndDate,omitempty"`
	//TotalCostEstimate float64 `xml:"totalCostEstimate,omitempty" json:"totalCostEstimate,omitempty"`
	//EstGrossProfit float64 `xml:"estGrossProfit,omitempty" json:"estGrossProfit,omitempty"`
	//EstGrossProfitPercent float64 `xml:"estGrossProfitPercent,omitempty" json:"estGrossProfitPercent,omitempty"`
	//ExchangeRate float64 `xml:"exchangeRate,omitempty" json:"exchangeRate,omitempty"`
	//CurrencyName string `xml:"currencyName,omitempty" json:"currencyName,omitempty"`
	//PromoCode *RecordRef `xml:"promoCode,omitempty" json:"promoCode,omitempty"`
	//DiscountItem *RecordRef `xml:"discountItem,omitempty" json:"discountItem,omitempty"`
	//DiscountRate string `xml:"discountRate,omitempty" json:"discountRate,omitempty"`
	//IsTaxable bool `xml:"isTaxable,omitempty" json:"isTaxable,omitempty"`
	//TaxItem *RecordRef `xml:"taxItem,omitempty" json:"taxItem,omitempty"`
	//TaxRate float64 `xml:"taxRate,omitempty" json:"taxRate,omitempty"`
	//ToBePrinted bool `xml:"toBePrinted,omitempty" json:"toBePrinted,omitempty"`
	//ToBeEmailed bool `xml:"toBeEmailed,omitempty" json:"toBeEmailed,omitempty"`
	//ToBeFaxed bool `xml:"toBeFaxed,omitempty" json:"toBeFaxed,omitempty"`
	//Fax string `xml:"fax,omitempty" json:"fax,omitempty"`
	//MessageSel *RecordRef `xml:"messageSel,omitempty" json:"messageSel,omitempty"`
	//Message string `xml:"message,omitempty" json:"message,omitempty"`
	//PaymentOption *RecordRef `xml:"paymentOption,omitempty" json:"paymentOption,omitempty"`
	//InputAuthCode string `xml:"inputAuthCode,omitempty" json:"inputAuthCode,omitempty"`
	//InputReferenceCode string `xml:"inputReferenceCode,omitempty" json:"inputReferenceCode,omitempty"`
	//CheckNumber string `xml:"checkNumber,omitempty" json:"checkNumber,omitempty"`
	//PaymentCardCsc string `xml:"paymentCardCsc,omitempty" json:"paymentCardCsc,omitempty"`
	//PaymentProcessingProfile *RecordRef `xml:"paymentProcessingProfile,omitempty" json:"paymentProcessingProfile,omitempty"`
	//HandlingMode *CashSaleHandlingMode `xml:"handlingMode,omitempty" json:"handlingMode,omitempty"`
	//OutputAuthCode string `xml:"outputAuthCode,omitempty" json:"outputAuthCode,omitempty"`
	//OutputReferenceCode string `xml:"outputReferenceCode,omitempty" json:"outputReferenceCode,omitempty"`
	//PaymentOperation *CashSalePaymentOperation `xml:"paymentOperation,omitempty" json:"paymentOperation,omitempty"`
	//DynamicDescriptor string `xml:"dynamicDescriptor,omitempty" json:"dynamicDescriptor,omitempty"`
	BillingAddress *Address `xml:"billingAddress,omitempty" json:"billingAddress,omitempty"`
	//BillAddressList *RecordRef `xml:"billAddressList,omitempty" json:"billAddressList,omitempty"`
	ShippingAddress *Address `xml:"shippingAddress,omitempty" json:"shippingAddress,omitempty"`
	//ShipIsResidential bool `xml:"shipIsResidential,omitempty" json:"shipIsResidential,omitempty"`
	//ShipAddressList *RecordRef `xml:"shipAddressList,omitempty" json:"shipAddressList,omitempty"`
	//Fob string `xml:"fob,omitempty" json:"fob,omitempty"`
	//ShipDate soap.XSDDateTime `xml:"shipDate,omitempty" json:"shipDate,omitempty"`
	//ShipMethod *RecordRef `xml:"shipMethod,omitempty" json:"shipMethod,omitempty"`
	//ShippingCost float64 `xml:"shippingCost,omitempty" json:"shippingCost,omitempty"`
	//ShippingTax1Rate float64 `xml:"shippingTax1Rate,omitempty" json:"shippingTax1Rate,omitempty"`
	//ShippingTax2Rate string `xml:"shippingTax2Rate,omitempty" json:"shippingTax2Rate,omitempty"`
	//ShippingTaxCode *RecordRef `xml:"shippingTaxCode,omitempty" json:"shippingTaxCode,omitempty"`
	//HandlingTaxCode *RecordRef `xml:"handlingTaxCode,omitempty" json:"handlingTaxCode,omitempty"`
	//HandlingTax1Rate float64 `xml:"handlingTax1Rate,omitempty" json:"handlingTax1Rate,omitempty"`
	//HandlingCost float64 `xml:"handlingCost,omitempty" json:"handlingCost,omitempty"`
	//HandlingTax2Rate string `xml:"handlingTax2Rate,omitempty" json:"handlingTax2Rate,omitempty"`
	//LinkedTrackingNumbers string `xml:"linkedTrackingNumbers,omitempty" json:"linkedTrackingNumbers,omitempty"`
	//TrackingNumbers string `xml:"trackingNumbers,omitempty" json:"trackingNumbers,omitempty"`
	//SalesGroup *RecordRef `xml:"salesGroup,omitempty" json:"salesGroup,omitempty"`
	//RevenueStatus *RevenueStatus `xml:"revenueStatus,omitempty" json:"revenueStatus,omitempty"`
	//RecognizedRevenue float64 `xml:"recognizedRevenue,omitempty" json:"recognizedRevenue,omitempty"`
	//DeferredRevenue float64 `xml:"deferredRevenue,omitempty" json:"deferredRevenue,omitempty"`
	//RevRecOnRevCommitment bool `xml:"revRecOnRevCommitment,omitempty" json:"revRecOnRevCommitment,omitempty"`
	//SyncSalesTeams bool `xml:"syncSalesTeams,omitempty" json:"syncSalesTeams,omitempty"`
	//PaymentMethod *RecordRef `xml:"paymentMethod,omitempty" json:"paymentMethod,omitempty"`
	//PayPalStatus string `xml:"payPalStatus,omitempty" json:"payPalStatus,omitempty"`
	//CreditCard *RecordRef `xml:"creditCard,omitempty" json:"creditCard,omitempty"`
	//CcNumber string `xml:"ccNumber,omitempty" json:"ccNumber,omitempty"`
	//CcExpireDate soap.XSDDateTime `xml:"ccExpireDate,omitempty" json:"ccExpireDate,omitempty"`
	//CcName string `xml:"ccName,omitempty" json:"ccName,omitempty"`
	//CcStreet string `xml:"ccStreet,omitempty" json:"ccStreet,omitempty"`
	//CcZipCode string `xml:"ccZipCode,omitempty" json:"ccZipCode,omitempty"`
	//CreditCardProcessor *RecordRef `xml:"creditCardProcessor,omitempty" json:"creditCardProcessor,omitempty"`
	//CcApproved bool `xml:"ccApproved,omitempty" json:"ccApproved,omitempty"`
	//AuthCode string `xml:"authCode,omitempty" json:"authCode,omitempty"`
	//CcAvsStreetMatch *AvsMatchCode `xml:"ccAvsStreetMatch,omitempty" json:"ccAvsStreetMatch,omitempty"`
	//CcAvsZipMatch *AvsMatchCode `xml:"ccAvsZipMatch,omitempty" json:"ccAvsZipMatch,omitempty"`
	//IsRecurringPayment bool `xml:"isRecurringPayment,omitempty" json:"isRecurringPayment,omitempty"`
	//PayPalTranId string `xml:"payPalTranId,omitempty" json:"payPalTranId,omitempty"`
	//SubTotal float64 `xml:"subTotal,omitempty" json:"subTotal,omitempty"`
	//CcIsPurchaseCardBin bool `xml:"ccIsPurchaseCardBin,omitempty" json:"ccIsPurchaseCardBin,omitempty"`
	//IgnoreAvs bool `xml:"ignoreAvs,omitempty" json:"ignoreAvs,omitempty"`
	//CcProcessAsPurchaseCard bool `xml:"ccProcessAsPurchaseCard,omitempty" json:"ccProcessAsPurchaseCard,omitempty"`
	//ItemCostDiscount *RecordRef `xml:"itemCostDiscount,omitempty" json:"itemCostDiscount,omitempty"`
	//ItemCostDiscRate string `xml:"itemCostDiscRate,omitempty" json:"itemCostDiscRate,omitempty"`
	//ItemCostDiscAmount float64 `xml:"itemCostDiscAmount,omitempty" json:"itemCostDiscAmount,omitempty"`
	//ItemCostTaxRate1 float64 `xml:"itemCostTaxRate1,omitempty" json:"itemCostTaxRate1,omitempty"`
	//ItemCostTaxRate2 float64 `xml:"itemCostTaxRate2,omitempty" json:"itemCostTaxRate2,omitempty"`
	//ItemCostDiscTaxable bool `xml:"itemCostDiscTaxable,omitempty" json:"itemCostDiscTaxable,omitempty"`
	//ItemCostTaxCode *RecordRef `xml:"itemCostTaxCode,omitempty" json:"itemCostTaxCode,omitempty"`
	//ItemCostDiscTax1Amt float64 `xml:"itemCostDiscTax1Amt,omitempty" json:"itemCostDiscTax1Amt,omitempty"`
	//ItemCostDiscPrint bool `xml:"itemCostDiscPrint,omitempty" json:"itemCostDiscPrint,omitempty"`
	//ExpCostDiscount *RecordRef `xml:"expCostDiscount,omitempty" json:"expCostDiscount,omitempty"`
	//ExpCostDiscRate string `xml:"expCostDiscRate,omitempty" json:"expCostDiscRate,omitempty"`
	//ExpCostDiscAmount float64 `xml:"expCostDiscAmount,omitempty" json:"expCostDiscAmount,omitempty"`
	//ExpCostDiscTaxable bool `xml:"expCostDiscTaxable,omitempty" json:"expCostDiscTaxable,omitempty"`
	//ExpCostDiscprint bool `xml:"expCostDiscprint,omitempty" json:"expCostDiscprint,omitempty"`
	//ExpCostTaxRate1 float64 `xml:"expCostTaxRate1,omitempty" json:"expCostTaxRate1,omitempty"`
	//TimeDiscount *RecordRef `xml:"timeDiscount,omitempty" json:"timeDiscount,omitempty"`
	//ExpCostTaxCode *RecordRef `xml:"expCostTaxCode,omitempty" json:"expCostTaxCode,omitempty"`
	//TimeDiscRate string `xml:"timeDiscRate,omitempty" json:"timeDiscRate,omitempty"`
	//ExpCostTaxRate2 float64 `xml:"expCostTaxRate2,omitempty" json:"expCostTaxRate2,omitempty"`
	//ExpCostDiscTax1Amt float64 `xml:"expCostDiscTax1Amt,omitempty" json:"expCostDiscTax1Amt,omitempty"`
	//TimeDiscAmount float64 `xml:"timeDiscAmount,omitempty" json:"timeDiscAmount,omitempty"`
	//TimeDiscTaxable bool `xml:"timeDiscTaxable,omitempty" json:"timeDiscTaxable,omitempty"`
	//TimeDiscPrint bool `xml:"timeDiscPrint,omitempty" json:"timeDiscPrint,omitempty"`
	//DiscountTotal float64 `xml:"discountTotal,omitempty" json:"discountTotal,omitempty"`
	//TaxTotal float64 `xml:"taxTotal,omitempty" json:"taxTotal,omitempty"`
	//TimeTaxRate1 float64 `xml:"timeTaxRate1,omitempty" json:"timeTaxRate1,omitempty"`
	//AltShippingCost float64 `xml:"altShippingCost,omitempty" json:"altShippingCost,omitempty"`
	//TimeTaxCode *RecordRef `xml:"timeTaxCode,omitempty" json:"timeTaxCode,omitempty"`
	//AltHandlingCost float64 `xml:"altHandlingCost,omitempty" json:"altHandlingCost,omitempty"`
	//Total float64 `xml:"total,omitempty" json:"total,omitempty"`
	//TimeDiscTax1Amt float64 `xml:"timeDiscTax1Amt,omitempty" json:"timeDiscTax1Amt,omitempty"`
	//CcSecurityCode string `xml:"ccSecurityCode,omitempty" json:"ccSecurityCode,omitempty"`
	//TimeTaxRate2 float64 `xml:"timeTaxRate2,omitempty" json:"timeTaxRate2,omitempty"`
	//CcSecurityCodeMatch *AvsMatchCode `xml:"ccSecurityCodeMatch,omitempty" json:"ccSecurityCodeMatch,omitempty"`
	//ChargeIt bool `xml:"chargeIt,omitempty" json:"chargeIt,omitempty"`
	//DebitCardIssueNo string `xml:"debitCardIssueNo,omitempty" json:"debitCardIssueNo,omitempty"`
	//ThreeDStatusCode string `xml:"threeDStatusCode,omitempty" json:"threeDStatusCode,omitempty"`
	//PnRefNum string `xml:"pnRefNum,omitempty" json:"pnRefNum,omitempty"`
	//PaypalAuthId string `xml:"paypalAuthId,omitempty" json:"paypalAuthId,omitempty"`
	Status string `xml:"status,omitempty" json:"status,omitempty"`
	//PaypalProcess bool `xml:"paypalProcess,omitempty" json:"paypalProcess,omitempty"`
	//Job *RecordRef `xml:"job,omitempty" json:"job,omitempty"`
	//BillingSchedule *RecordRef `xml:"billingSchedule,omitempty" json:"billingSchedule,omitempty"`
	Email string `xml:"email,omitempty" json:"email,omitempty"`
	//Tax2Total float64 `xml:"tax2Total,omitempty" json:"tax2Total,omitempty"`
	ValidFrom soap.XSDDateTime `xml:"validFrom,omitempty" json:"validFrom,omitempty"`
	//VatRegNum string `xml:"vatRegNum,omitempty" json:"vatRegNum,omitempty"`
	//GiftCertApplied float64 `xml:"giftCertApplied,omitempty" json:"giftCertApplied,omitempty"`
	//TranIsVsoeBundle bool `xml:"tranIsVsoeBundle,omitempty" json:"tranIsVsoeBundle,omitempty"`
	//VsoeAutoCalc bool `xml:"vsoeAutoCalc,omitempty" json:"vsoeAutoCalc,omitempty"`
	//SyncPartnerTeams bool `xml:"syncPartnerTeams,omitempty" json:"syncPartnerTeams,omitempty"`
	//SalesTeamList *CashSaleSalesTeamList `xml:"salesTeamList,omitempty" json:"salesTeamList,omitempty"`
	//PartnersList *CashSalePartnersList `xml:"partnersList,omitempty" json:"partnersList,omitempty"`
	ItemList *CashSaleItemList `xml:"itemList,omitempty" json:"itemList,omitempty"`
	//AccountingBookDetailList *AccountingBookDetailList `xml:"accountingBookDetailList,omitempty" json:"accountingBookDetailList,omitempty"`
	//ItemCostList *CashSaleItemCostList `xml:"itemCostList,omitempty" json:"itemCostList,omitempty"`
	//GiftCertRedemptionList *GiftCertRedemptionList `xml:"giftCertRedemptionList,omitempty" json:"giftCertRedemptionList,omitempty"`
	//PromotionsList *PromotionsList `xml:"promotionsList,omitempty" json:"promotionsList,omitempty"`
	//ExpCostList *CashSaleExpCostList `xml:"expCostList,omitempty" json:"expCostList,omitempty"`
	//TimeList *CashSaleTimeList `xml:"timeList,omitempty" json:"timeList,omitempty"`
	//ShipGroupList *CashSaleShipGroupList `xml:"shipGroupList,omitempty" json:"shipGroupList,omitempty"`
	//TaxDetailsList *TaxDetailsList `xml:"taxDetailsList,omitempty" json:"taxDetailsList,omitempty"`
	CustomFieldList *CustomFieldList `xml:"customFieldList,omitempty" json:"customFieldList,omitempty"`
	InternalId string `xml:"internalId,attr,omitempty" json:"internalId,omitempty"`
	ExternalId string `xml:"externalId,attr,omitempty" json:"externalId,omitempty"`
}
type CashSaleItem struct {
	XMLName xml.Name `xml:"item"`
	Job *RecordRef `xml:"job,omitempty" json:"job,omitempty"`
	Item *RecordRef `xml:"item,omitempty" json:"item,omitempty"`
	Line int64 `xml:"line,omitempty" json:"line,omitempty"`
	QuantityAvailable float64 `xml:"quantityAvailable,omitempty" json:"quantityAvailable,omitempty"`
	QuantityOnHand float64 `xml:"quantityOnHand,omitempty" json:"quantityOnHand,omitempty"`
	QuantityFulfilled float64 `xml:"quantityFulfilled,omitempty" json:"quantityFulfilled,omitempty"`
	Quantity float64 `xml:"quantity,omitempty" json:"quantity,omitempty"`
	Units *RecordRef `xml:"units,omitempty" json:"units,omitempty"`
	//InventoryDetail *InventoryDetail `xml:"inventoryDetail,omitempty" json:"inventoryDetail,omitempty"`
	SerialNumbers string `xml:"serialNumbers,omitempty" json:"serialNumbers,omitempty"`
	//BinNumbers string `xml:"binNumbers,omitempty" json:"binNumbers,omitempty"`
	Description string `xml:"description,omitempty" json:"description,omitempty"`
	Price *RecordRef `xml:"price,omitempty" json:"price,omitempty"`
	Rate string `xml:"rate,omitempty" json:"rate,omitempty"`
	Amount float64 `xml:"amount,omitempty" json:"amount,omitempty"`
	OrderLine int64 `xml:"orderLine,omitempty" json:"orderLine,omitempty"`
	//LicenseCode string `xml:"licenseCode,omitempty" json:"licenseCode,omitempty"`
	//IsTaxable bool `xml:"isTaxable,omitempty" json:"isTaxable,omitempty"`
	Options *CustomFieldList `xml:"options,omitempty" json:"options,omitempty"`
	//DeferRevRec bool `xml:"deferRevRec,omitempty" json:"deferRevRec,omitempty"`
	//CurrentPercent float64 `xml:"currentPercent,omitempty" json:"currentPercent,omitempty"`
	Department *RecordRef `xml:"department,omitempty" json:"department,omitempty"`
	PercentComplete float64 `xml:"percentComplete,omitempty" json:"percentComplete,omitempty"`
	Class *RecordRef `xml:"class,omitempty" json:"class,omitempty"`
	Location *RecordRef `xml:"location,omitempty" json:"location,omitempty"`
	//RevRecSchedule *RecordRef `xml:"revRecSchedule,omitempty" json:"revRecSchedule,omitempty"`
	//RevRecStartDate soap.XSDDateTime `xml:"revRecStartDate,omitempty" json:"revRecStartDate,omitempty"`
	//RevRecEndDate soap.XSDDateTime `xml:"revRecEndDate,omitempty" json:"revRecEndDate,omitempty"`
	//SubscriptionLine *RecordRef `xml:"subscriptionLine,omitempty" json:"subscriptionLine,omitempty"`
	//GrossAmt float64 `xml:"grossAmt,omitempty" json:"grossAmt,omitempty"`
	//CostEstimateType *ItemCostEstimateType `xml:"costEstimateType,omitempty" json:"costEstimateType,omitempty"`
	//ExcludeFromRateRequest bool `xml:"excludeFromRateRequest,omitempty" json:"excludeFromRateRequest,omitempty"`
	//CatchUpPeriod *RecordRef `xml:"catchUpPeriod,omitempty" json:"catchUpPeriod,omitempty"`
	//CostEstimate float64 `xml:"costEstimate,omitempty" json:"costEstimate,omitempty"`
	//TaxDetailsReference string `xml:"taxDetailsReference,omitempty" json:"taxDetailsReference,omitempty"`
	//AmountOrdered float64 `xml:"amountOrdered,omitempty" json:"amountOrdered,omitempty"`
	//Tax1Amt float64 `xml:"tax1Amt,omitempty" json:"tax1Amt,omitempty"`
	QuantityOrdered float64 `xml:"quantityOrdered,omitempty" json:"quantityOrdered,omitempty"`
	QuantityRemaining float64 `xml:"quantityRemaining,omitempty" json:"quantityRemaining,omitempty"`
	//TaxCode *RecordRef `xml:"taxCode,omitempty" json:"taxCode,omitempty"`
	//TaxRate1 float64 `xml:"taxRate1,omitempty" json:"taxRate1,omitempty"`
	//TaxRate2 float64 `xml:"taxRate2,omitempty" json:"taxRate2,omitempty"`
	//GiftCertFrom string `xml:"giftCertFrom,omitempty" json:"giftCertFrom,omitempty"`
	//GiftCertRecipientName string `xml:"giftCertRecipientName,omitempty" json:"giftCertRecipientName,omitempty"`
	//GiftCertRecipientEmail string `xml:"giftCertRecipientEmail,omitempty" json:"giftCertRecipientEmail,omitempty"`
	//GiftCertMessage string `xml:"giftCertMessage,omitempty" json:"giftCertMessage,omitempty"`
	TaxAmount float64 `xml:"taxAmount,omitempty" json:"taxAmount,omitempty"`
	//GiftCertNumber string `xml:"giftCertNumber,omitempty" json:"giftCertNumber,omitempty"`
	//ShipGroup int64 `xml:"shipGroup,omitempty" json:"shipGroup,omitempty"`
	ItemIsFulfilled bool `xml:"itemIsFulfilled,omitempty" json:"itemIsFulfilled,omitempty"`
	ShipAddress *RecordRef `xml:"shipAddress,omitempty" json:"shipAddress,omitempty"`
	//ShipMethod *RecordRef `xml:"shipMethod,omitempty" json:"shipMethod,omitempty"`
	//VsoeSopGroup *VsoeSopGroup `xml:"vsoeSopGroup,omitempty" json:"vsoeSopGroup,omitempty"`
	//VsoeIsEstimate bool `xml:"vsoeIsEstimate,omitempty" json:"vsoeIsEstimate,omitempty"`
	//VsoePrice float64 `xml:"vsoePrice,omitempty" json:"vsoePrice,omitempty"`
	//VsoeAmount float64 `xml:"vsoeAmount,omitempty" json:"vsoeAmount,omitempty"`
	//VsoeAllocation float64 `xml:"vsoeAllocation,omitempty" json:"vsoeAllocation,omitempty"`
	//VsoeDeferral *VsoeDeferral `xml:"vsoeDeferral,omitempty" json:"vsoeDeferral,omitempty"`
	//VsoePermitDiscount *VsoePermitDiscount `xml:"vsoePermitDiscount,omitempty" json:"vsoePermitDiscount,omitempty"`
	//VsoeDelivered bool `xml:"vsoeDelivered,omitempty" json:"vsoeDelivered,omitempty"`
	//ChargeType *RecordRef `xml:"chargeType,omitempty" json:"chargeType,omitempty"`
	//ChargesList *RecordRefList `xml:"chargesList,omitempty" json:"chargesList,omitempty"`
	CustomFieldList *CustomFieldList `xml:"customFieldList,omitempty" json:"customFieldList,omitempty"`
}
type CashSaleItemList struct {
	XMLName xml.Name `xml:"itemList"`
	Item []*CashSaleItem `xml:"item,omitempty" json:"item,omitempty"`
	ReplaceAll bool `xml:"replaceAll,attr,omitempty" json:"replaceAll,omitempty"`
}
