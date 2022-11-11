//GENERATED via go2wsdl, (and modified for naming conflicts, and functionality)

package netsuite

import (
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
)

type SalesOrder_AddRequest struct {
	XMLName xml.Name `xml:"add"`
	XmlNSXSI   string   `xml:"xmlns:xsi,attr,omitempty"` 
	XmlNSPC   string   `xml:"xmlns:platformCore,attr,omitempty"` 
	XmlNS1   string   `xml:"xmlns:ns1,attr,omitempty"` 
	Record *SalesOrder `xml:"record,omitempty" json:"record,omitempty"`
}

type SalesOrder_AddResponse struct {
	XMLName xml.Name `xml:"addResponse"`
	WriteResponse *SalesOrder_WriteResponse `xml:"writeResponse,omitempty" json:"writeResponse,omitempty"`
}

type SalesOrder_WriteResponse struct {
	XMLName xml.Name `xml:"writeResponse"`
	Status *Status `xml:"status,omitempty" json:"status,omitempty"`
	BaseRef *BaseRef `xml:"baseRef,omitempty" json:"baseRef,omitempty"`
}

//----------------------------------------------------

type SalesOrder_GetRequest struct {
	XMLName xml.Name `xml:"get"`
	XmlNSXSI   string   `xml:"xmlns:xsi,attr,omitempty"` 
	XmlNSPC   string   `xml:"xmlns:platformCore,attr,omitempty"` 
	BaseRef *BaseRef `xml:"baseRef,omitempty" json:"baseRef,omitempty"`
}

type SalesOrder_GetResponse struct {
	//XMLName xml.Name `xml:"urn:messages_2022_1.platform.webservices.netsuite.com getResponse"`
	XMLName xml.Name `xml:"getResponse"`
	ReadResponse *SalesOrder_ReadResponse `xml:"readResponse,omitempty" json:"readResponse,omitempty"`
}
//----------------------------------------------------

type SalesOrder_ReadResponse struct {
	XMLName xml.Name `xml:"readResponse"`
	Status *Status `xml:"status,omitempty" json:"status,omitempty"`
	Record *SalesOrder `xml:"record,omitempty" json:"record,omitempty"`
}

type SalesOrder struct {
	XsiType string `xml:"xsi:type,attr,omitempty" json:"-"` //ADDED
	CreatedDate soap.XSDDateTime `xml:"createdDate,omitempty" json:"createdDate,omitempty"`
	CustomForm *RecordRef `xml:"customForm,omitempty" json:"customForm,omitempty"`
	Entity *RecordRef `xml:"entity,omitempty" json:"entity,omitempty"`
	//Job *RecordRef `xml:"job,omitempty" json:"job,omitempty"`
	Currency *RecordRef `xml:"currency,omitempty" json:"currency,omitempty"`
	//DrAccount *RecordRef `xml:"drAccount,omitempty" json:"drAccount,omitempty"`
	//FxAccount *RecordRef `xml:"fxAccount,omitempty" json:"fxAccount,omitempty"`
	TranDate soap.XSDDateTime `xml:"tranDate,omitempty" json:"tranDate,omitempty"`
	TranId string `xml:"tranId,omitempty" json:"tranId,omitempty"`
	EntityTaxRegNum *RecordRef `xml:"entityTaxRegNum,omitempty" json:"entityTaxRegNum,omitempty"`
	Source string `xml:"source,omitempty" json:"source,omitempty"`
	CreatedFrom *RecordRef `xml:"createdFrom,omitempty" json:"createdFrom,omitempty"`
//	OrderStatus *SalesOrderOrderStatus `xml:"orderStatus,omitempty" json:"orderStatus,omitempty"`
//	NextBill soap.XSDDateTime `xml:"nextBill,omitempty" json:"nextBill,omitempty"`
//	Opportunity *RecordRef `xml:"opportunity,omitempty" json:"opportunity,omitempty"`
	SalesRep *RecordRef `xml:"salesRep,omitempty" json:"salesRep,omitempty"`
//	ContribPct string `xml:"contribPct,omitempty" json:"contribPct,omitempty"`
//	Partner *RecordRef `xml:"partner,omitempty" json:"partner,omitempty"`
//	SalesGroup *RecordRef `xml:"salesGroup,omitempty" json:"salesGroup,omitempty"`
//	SyncSalesTeams bool `xml:"syncSalesTeams,omitempty" json:"syncSalesTeams,omitempty"`
//	LeadSource *RecordRef `xml:"leadSource,omitempty" json:"leadSource,omitempty"`
	StartDate soap.XSDDateTime `xml:"startDate,omitempty" json:"startDate,omitempty"`
	EndDate soap.XSDDateTime `xml:"endDate,omitempty" json:"endDate,omitempty"`
	OtherRefNum string `xml:"otherRefNum,omitempty" json:"otherRefNum,omitempty"`
	Memo string `xml:"memo,omitempty" json:"memo,omitempty"`
	SalesEffectiveDate soap.XSDDateTime `xml:"salesEffectiveDate,omitempty" json:"salesEffectiveDate,omitempty"`
//	ExcludeCommission bool `xml:"excludeCommission,omitempty" json:"excludeCommission,omitempty"`
//	TotalCostEstimate float64 `xml:"totalCostEstimate,omitempty" json:"totalCostEstimate,omitempty"`
//	EstGrossProfit float64 `xml:"estGrossProfit,omitempty" json:"estGrossProfit,omitempty"`
//	EstGrossProfitPercent float64 `xml:"estGrossProfitPercent,omitempty" json:"estGrossProfitPercent,omitempty"`
//	ExchangeRate float64 `xml:"exchangeRate,omitempty" json:"exchangeRate,omitempty"`
//	PromoCode *RecordRef `xml:"promoCode,omitempty" json:"promoCode,omitempty"`
//	CurrencyName string `xml:"currencyName,omitempty" json:"currencyName,omitempty"`
//	DiscountItem *RecordRef `xml:"discountItem,omitempty" json:"discountItem,omitempty"`
//	DiscountRate string `xml:"discountRate,omitempty" json:"discountRate,omitempty"`
//	IsTaxable bool `xml:"isTaxable,omitempty" json:"isTaxable,omitempty"`
//	TaxItem *RecordRef `xml:"taxItem,omitempty" json:"taxItem,omitempty"`
//	TaxRate float64 `xml:"taxRate,omitempty" json:"taxRate,omitempty"`
//	ToBePrinted bool `xml:"toBePrinted,omitempty" json:"toBePrinted,omitempty"`
//	ToBeEmailed bool `xml:"toBeEmailed,omitempty" json:"toBeEmailed,omitempty"`
	Email string `xml:"email,omitempty" json:"email,omitempty"`
//	ToBeFaxed bool `xml:"toBeFaxed,omitempty" json:"toBeFaxed,omitempty"`
//	Fax string `xml:"fax,omitempty" json:"fax,omitempty"`
//	MessageSel *RecordRef `xml:"messageSel,omitempty" json:"messageSel,omitempty"`
	Message string `xml:"message,omitempty" json:"message,omitempty"`
//	PaymentOption *RecordRef `xml:"paymentOption,omitempty" json:"paymentOption,omitempty"`
//	InputAuthCode string `xml:"inputAuthCode,omitempty" json:"inputAuthCode,omitempty"`
//	InputReferenceCode string `xml:"inputReferenceCode,omitempty" json:"inputReferenceCode,omitempty"`
//	CheckNumber string `xml:"checkNumber,omitempty" json:"checkNumber,omitempty"`
//	PaymentCardCsc string `xml:"paymentCardCsc,omitempty" json:"paymentCardCsc,omitempty"`
//	PaymentProcessingProfile *RecordRef `xml:"paymentProcessingProfile,omitempty" json:"paymentProcessingProfile,omitempty"`
//	HandlingMode *SalesOrderHandlingMode `xml:"handlingMode,omitempty" json:"handlingMode,omitempty"`
//	OutputAuthCode string `xml:"outputAuthCode,omitempty" json:"outputAuthCode,omitempty"`
//	OutputReferenceCode string `xml:"outputReferenceCode,omitempty" json:"outputReferenceCode,omitempty"`
//	PaymentOperation *SalesOrderPaymentOperation `xml:"paymentOperation,omitempty" json:"paymentOperation,omitempty"`
//	DynamicDescriptor string `xml:"dynamicDescriptor,omitempty" json:"dynamicDescriptor,omitempty"`
	BillingAddress *Address `xml:"billingAddress,omitempty" json:"billingAddress,omitempty"`
//	BillAddressList *RecordRef `xml:"billAddressList,omitempty" json:"billAddressList,omitempty"`
	ShippingAddress *Address `xml:"shippingAddress,omitempty" json:"shippingAddress,omitempty"`
//	ShipIsResidential bool `xml:"shipIsResidential,omitempty" json:"shipIsResidential,omitempty"`
//	ShipAddressList *RecordRef `xml:"shipAddressList,omitempty" json:"shipAddressList,omitempty"`
//	Fob string `xml:"fob,omitempty" json:"fob,omitempty"`
//	ShipDate soap.XSDDateTime `xml:"shipDate,omitempty" json:"shipDate,omitempty"`
//	ActualShipDate soap.XSDDateTime `xml:"actualShipDate,omitempty" json:"actualShipDate,omitempty"`
//	ShipMethod *RecordRef `xml:"shipMethod,omitempty" json:"shipMethod,omitempty"`
//	ShippingCost float64 `xml:"shippingCost,omitempty" json:"shippingCost,omitempty"`
//	ShippingTax1Rate float64 `xml:"shippingTax1Rate,omitempty" json:"shippingTax1Rate,omitempty"`
//	IsMultiShipTo bool `xml:"isMultiShipTo,omitempty" json:"isMultiShipTo,omitempty"`
//	ShippingTax2Rate string `xml:"shippingTax2Rate,omitempty" json:"shippingTax2Rate,omitempty"`
//	ShippingTaxCode *RecordRef `xml:"shippingTaxCode,omitempty" json:"shippingTaxCode,omitempty"`
//	HandlingTaxCode *RecordRef `xml:"handlingTaxCode,omitempty" json:"handlingTaxCode,omitempty"`
//	HandlingTax1Rate float64 `xml:"handlingTax1Rate,omitempty" json:"handlingTax1Rate,omitempty"`
//	HandlingTax2Rate string `xml:"handlingTax2Rate,omitempty" json:"handlingTax2Rate,omitempty"`
//	HandlingCost float64 `xml:"handlingCost,omitempty" json:"handlingCost,omitempty"`
//	TrackingNumbers string `xml:"trackingNumbers,omitempty" json:"trackingNumbers,omitempty"`
//	LinkedTrackingNumbers string `xml:"linkedTrackingNumbers,omitempty" json:"linkedTrackingNumbers,omitempty"`
//	ShipComplete bool `xml:"shipComplete,omitempty" json:"shipComplete,omitempty"`
//	PaymentMethod *RecordRef `xml:"paymentMethod,omitempty" json:"paymentMethod,omitempty"`
//	ShopperIpAddress string `xml:"shopperIpAddress,omitempty" json:"shopperIpAddress,omitempty"`
//	SaveOnAuthDecline bool `xml:"saveOnAuthDecline,omitempty" json:"saveOnAuthDecline,omitempty"`
//	CanHaveStackable bool `xml:"canHaveStackable,omitempty" json:"canHaveStackable,omitempty"`
//	CreditCard *RecordRef `xml:"creditCard,omitempty" json:"creditCard,omitempty"`
//	RevenueStatus *RevenueStatus `xml:"revenueStatus,omitempty" json:"revenueStatus,omitempty"`
//	RecognizedRevenue float64 `xml:"recognizedRevenue,omitempty" json:"recognizedRevenue,omitempty"`
//	DeferredRevenue float64 `xml:"deferredRevenue,omitempty" json:"deferredRevenue,omitempty"`
//	RevRecOnRevCommitment bool `xml:"revRecOnRevCommitment,omitempty" json:"revRecOnRevCommitment,omitempty"`
//	RevCommitStatus *RevenueCommitStatus `xml:"revCommitStatus,omitempty" json:"revCommitStatus,omitempty"`
//	CcNumber string `xml:"ccNumber,omitempty" json:"ccNumber,omitempty"`
//	CcExpireDate soap.XSDDateTime `xml:"ccExpireDate,omitempty" json:"ccExpireDate,omitempty"`
//	CcName string `xml:"ccName,omitempty" json:"ccName,omitempty"`
//	CcStreet string `xml:"ccStreet,omitempty" json:"ccStreet,omitempty"`
//	CcZipCode string `xml:"ccZipCode,omitempty" json:"ccZipCode,omitempty"`
//	PayPalStatus string `xml:"payPalStatus,omitempty" json:"payPalStatus,omitempty"`
//	CreditCardProcessor *RecordRef `xml:"creditCardProcessor,omitempty" json:"creditCardProcessor,omitempty"`
//	PayPalTranId string `xml:"payPalTranId,omitempty" json:"payPalTranId,omitempty"`
//	CcApproved bool `xml:"ccApproved,omitempty" json:"ccApproved,omitempty"`
//	GetAuth bool `xml:"getAuth,omitempty" json:"getAuth,omitempty"`
//	AuthCode string `xml:"authCode,omitempty" json:"authCode,omitempty"`
//	CcAvsStreetMatch *AvsMatchCode `xml:"ccAvsStreetMatch,omitempty" json:"ccAvsStreetMatch,omitempty"`
//	CcAvsZipMatch *AvsMatchCode `xml:"ccAvsZipMatch,omitempty" json:"ccAvsZipMatch,omitempty"`
//	IsRecurringPayment bool `xml:"isRecurringPayment,omitempty" json:"isRecurringPayment,omitempty"`
//	CcSecurityCodeMatch *AvsMatchCode `xml:"ccSecurityCodeMatch,omitempty" json:"ccSecurityCodeMatch,omitempty"`
//	AltSalesTotal float64 `xml:"altSalesTotal,omitempty" json:"altSalesTotal,omitempty"`
//	IgnoreAvs bool `xml:"ignoreAvs,omitempty" json:"ignoreAvs,omitempty"`
//	PaymentEventResult *TransactionPaymentEventResult `xml:"paymentEventResult,omitempty" json:"paymentEventResult,omitempty"`
//	PaymentEventHoldReason *TransactionPaymentEventHoldReason `xml:"paymentEventHoldReason,omitempty" json:"paymentEventHoldReason,omitempty"`
//	PaymentEventType *TransactionPaymentEventType `xml:"paymentEventType,omitempty" json:"paymentEventType,omitempty"`
//	PaymentEventDate soap.XSDDateTime `xml:"paymentEventDate,omitempty" json:"paymentEventDate,omitempty"`
//	PaymentEventUpdatedBy string `xml:"paymentEventUpdatedBy,omitempty" json:"paymentEventUpdatedBy,omitempty"`
//	SubTotal float64 `xml:"subTotal,omitempty" json:"subTotal,omitempty"`
//	DiscountTotal float64 `xml:"discountTotal,omitempty" json:"discountTotal,omitempty"`
//	TaxTotal float64 `xml:"taxTotal,omitempty" json:"taxTotal,omitempty"`
//	AltShippingCost float64 `xml:"altShippingCost,omitempty" json:"altShippingCost,omitempty"`
//	AltHandlingCost float64 `xml:"altHandlingCost,omitempty" json:"altHandlingCost,omitempty"`
//	Total float64 `xml:"total,omitempty" json:"total,omitempty"`
//	RevRecSchedule *RecordRef `xml:"revRecSchedule,omitempty" json:"revRecSchedule,omitempty"`
//	RevRecStartDate soap.XSDDateTime `xml:"revRecStartDate,omitempty" json:"revRecStartDate,omitempty"`
//	RevRecEndDate soap.XSDDateTime `xml:"revRecEndDate,omitempty" json:"revRecEndDate,omitempty"`
//	PaypalAuthId string `xml:"paypalAuthId,omitempty" json:"paypalAuthId,omitempty"`
//	Balance float64 `xml:"balance,omitempty" json:"balance,omitempty"`
//	PaypalProcess bool `xml:"paypalProcess,omitempty" json:"paypalProcess,omitempty"`
//	BillingSchedule *RecordRef `xml:"billingSchedule,omitempty" json:"billingSchedule,omitempty"`
//	CcSecurityCode string `xml:"ccSecurityCode,omitempty" json:"ccSecurityCode,omitempty"`
//	ThreeDStatusCode string `xml:"threeDStatusCode,omitempty" json:"threeDStatusCode,omitempty"`
	Class *RecordRef `xml:"class,omitempty" json:"class,omitempty"`
//	Department *RecordRef `xml:"department,omitempty" json:"department,omitempty"`
	Subsidiary *RecordRef `xml:"subsidiary,omitempty" json:"subsidiary,omitempty"`
//	IntercoTransaction *RecordRef `xml:"intercoTransaction,omitempty" json:"intercoTransaction,omitempty"`
//	IntercoStatus *IntercoStatus `xml:"intercoStatus,omitempty" json:"intercoStatus,omitempty"`
//	DebitCardIssueNo string `xml:"debitCardIssueNo,omitempty" json:"debitCardIssueNo,omitempty"`
//	LastModifiedDate soap.XSDDateTime `xml:"lastModifiedDate,omitempty" json:"lastModifiedDate,omitempty"`
//	Nexus *RecordRef `xml:"nexus,omitempty" json:"nexus,omitempty"`
//	SubsidiaryTaxRegNum *RecordRef `xml:"subsidiaryTaxRegNum,omitempty" json:"subsidiaryTaxRegNum,omitempty"`
//	TaxRegOverride bool `xml:"taxRegOverride,omitempty" json:"taxRegOverride,omitempty"`
//	TaxPointDate soap.XSDDateTime `xml:"taxPointDate,omitempty" json:"taxPointDate,omitempty"`
//	TaxDetailsOverride bool `xml:"taxDetailsOverride,omitempty" json:"taxDetailsOverride,omitempty"`
//	Location *RecordRef `xml:"location,omitempty" json:"location,omitempty"`
//	PnRefNum string `xml:"pnRefNum,omitempty" json:"pnRefNum,omitempty"`
//	Status string `xml:"status,omitempty" json:"status,omitempty"`
//	Tax2Total float64 `xml:"tax2Total,omitempty" json:"tax2Total,omitempty"`
	Terms *RecordRef `xml:"terms,omitempty" json:"terms,omitempty"`
//	ValidFrom soap.XSDDateTime `xml:"validFrom,omitempty" json:"validFrom,omitempty"`
//	VatRegNum string `xml:"vatRegNum,omitempty" json:"vatRegNum,omitempty"`
//	GiftCertApplied float64 `xml:"giftCertApplied,omitempty" json:"giftCertApplied,omitempty"`
//	OneTime float64 `xml:"oneTime,omitempty" json:"oneTime,omitempty"`
//	RecurWeekly float64 `xml:"recurWeekly,omitempty" json:"recurWeekly,omitempty"`
//	RecurMonthly float64 `xml:"recurMonthly,omitempty" json:"recurMonthly,omitempty"`
//	RecurQuarterly float64 `xml:"recurQuarterly,omitempty" json:"recurQuarterly,omitempty"`
//	RecurAnnually float64 `xml:"recurAnnually,omitempty" json:"recurAnnually,omitempty"`
//	TranIsVsoeBundle bool `xml:"tranIsVsoeBundle,omitempty" json:"tranIsVsoeBundle,omitempty"`
//	VsoeAutoCalc bool `xml:"vsoeAutoCalc,omitempty" json:"vsoeAutoCalc,omitempty"`
//	SyncPartnerTeams bool `xml:"syncPartnerTeams,omitempty" json:"syncPartnerTeams,omitempty"`
//	SalesTeamList *SalesOrderSalesTeamList `xml:"salesTeamList,omitempty" json:"salesTeamList,omitempty"`
//	PartnersList *SalesOrderPartnersList `xml:"partnersList,omitempty" json:"partnersList,omitempty"`
//	GiftCertRedemptionList *GiftCertRedemptionList `xml:"giftCertRedemptionList,omitempty" json:"giftCertRedemptionList,omitempty"`
//	PromotionsList *PromotionsList `xml:"promotionsList,omitempty" json:"promotionsList,omitempty"`
	ItemList *SalesOrderItemList `xml:"itemList,omitempty" json:"itemList,omitempty"`
//	ShipGroupList *SalesOrderShipGroupList `xml:"shipGroupList,omitempty" json:"shipGroupList,omitempty"`
//	AccountingBookDetailList *AccountingBookDetailList `xml:"accountingBookDetailList,omitempty" json:"accountingBookDetailList,omitempty"`
//	TaxDetailsList *TaxDetailsList `xml:"taxDetailsList,omitempty" json:"taxDetailsList,omitempty"`
	CustomFieldList *CustomFieldList `xml:"customFieldList,omitempty" json:"customFieldList,omitempty"`
	InternalId string `xml:"internalId,attr,omitempty" json:"internalId,omitempty"`
	ExternalId string `xml:"externalId,attr,omitempty" json:"externalId,omitempty"`
}

type SalesOrderItem struct {
	XMLName xml.Name `xml:"item" json:"-"`
	Job *RecordRef `xml:"job,omitempty" json:"job,omitempty"`
	Subscription *RecordRef `xml:"subscription,omitempty" json:"subscription,omitempty"`
	Item *RecordRef `xml:"item,omitempty" json:"item,omitempty"`
//	QuantityAvailable float64 `xml:"quantityAvailable,omitempty" json:"quantityAvailable,omitempty"`
//	ExpandItemGroup bool `xml:"expandItemGroup,omitempty" json:"expandItemGroup,omitempty"`
//	LineUniqueKey int64 `xml:"lineUniqueKey,omitempty" json:"lineUniqueKey,omitempty"`
//	QuantityOnHand float64 `xml:"quantityOnHand,omitempty" json:"quantityOnHand,omitempty"`
	Quantity float64 `xml:"quantity,omitempty" json:"quantity,omitempty"`
//	Units *RecordRef `xml:"units,omitempty" json:"units,omitempty"`
//	InventoryDetail *InventoryDetail `xml:"inventoryDetail,omitempty" json:"inventoryDetail,omitempty"`
	Description string `xml:"description,omitempty" json:"description,omitempty"`
//	Price *RecordRef `xml:"price,omitempty" json:"price,omitempty"`
//	Rate string `xml:"rate,omitempty" json:"rate,omitempty"`
//	SerialNumbers string `xml:"serialNumbers,omitempty" json:"serialNumbers,omitempty"`
	Amount float64 `xml:"amount,omitempty" json:"amount,omitempty"`
//	IsTaxable bool `xml:"isTaxable,omitempty" json:"isTaxable,omitempty"`
//	CommitInventory *SalesOrderItemCommitInventory `xml:"commitInventory,omitempty" json:"commitInventory,omitempty"`
//	OrderPriority float64 `xml:"orderPriority,omitempty" json:"orderPriority,omitempty"`
//	LicenseCode string `xml:"licenseCode,omitempty" json:"licenseCode,omitempty"`
//	Options *CustomFieldList `xml:"options,omitempty" json:"options,omitempty"`
//	Department *RecordRef `xml:"department,omitempty" json:"department,omitempty"`
//	Class *RecordRef `xml:"class,omitempty" json:"class,omitempty"`
//	Location *RecordRef `xml:"location,omitempty" json:"location,omitempty"`
//	CreatePo *SalesOrderItemCreatePo `xml:"createPo,omitempty" json:"createPo,omitempty"`
//	CreatedPo *RecordRef `xml:"createdPo,omitempty" json:"createdPo,omitempty"`
//	AltSalesAmt float64 `xml:"altSalesAmt,omitempty" json:"altSalesAmt,omitempty"`
//	CreateWo bool `xml:"createWo,omitempty" json:"createWo,omitempty"`
//	PoVendor *RecordRef `xml:"poVendor,omitempty" json:"poVendor,omitempty"`
//	PoCurrency string `xml:"poCurrency,omitempty" json:"poCurrency,omitempty"`
//	PoRate float64 `xml:"poRate,omitempty" json:"poRate,omitempty"`
//	RevRecSchedule *RecordRef `xml:"revRecSchedule,omitempty" json:"revRecSchedule,omitempty"`
//	RevRecStartDate soap.XSDDateTime `xml:"revRecStartDate,omitempty" json:"revRecStartDate,omitempty"`
//	RevRecTermInMonths int64 `xml:"revRecTermInMonths,omitempty" json:"revRecTermInMonths,omitempty"`
//	RevRecEndDate soap.XSDDateTime `xml:"revRecEndDate,omitempty" json:"revRecEndDate,omitempty"`
//	DeferRevRec bool `xml:"deferRevRec,omitempty" json:"deferRevRec,omitempty"`
//	IsClosed bool `xml:"isClosed,omitempty" json:"isClosed,omitempty"`
//	ItemFulfillmentChoice *SalesOrderItemFulfillmentChoice `xml:"itemFulfillmentChoice,omitempty" json:"itemFulfillmentChoice,omitempty"`
//	CatchUpPeriod *RecordRef `xml:"catchUpPeriod,omitempty" json:"catchUpPeriod,omitempty"`
//	BillingSchedule *RecordRef `xml:"billingSchedule,omitempty" json:"billingSchedule,omitempty"`
//	FromJob bool `xml:"fromJob,omitempty" json:"fromJob,omitempty"`
//	GrossAmt float64 `xml:"grossAmt,omitempty" json:"grossAmt,omitempty"`
//	TaxAmount float64 `xml:"taxAmount,omitempty" json:"taxAmount,omitempty"`
//	ExcludeFromRateRequest bool `xml:"excludeFromRateRequest,omitempty" json:"excludeFromRateRequest,omitempty"`
//	IsEstimate bool `xml:"isEstimate,omitempty" json:"isEstimate,omitempty"`
//	InventoryLocation *RecordRef `xml:"inventoryLocation,omitempty" json:"inventoryLocation,omitempty"`
//	InventorySubsidiary *RecordRef `xml:"inventorySubsidiary,omitempty" json:"inventorySubsidiary,omitempty"`
	Line int64 `xml:"line,omitempty" json:"line,omitempty"`
//	PercentComplete float64 `xml:"percentComplete,omitempty" json:"percentComplete,omitempty"`
//	CostEstimateType *ItemCostEstimateType `xml:"costEstimateType,omitempty" json:"costEstimateType,omitempty"`
//	CostEstimate float64 `xml:"costEstimate,omitempty" json:"costEstimate,omitempty"`
//	QuantityBackOrdered float64 `xml:"quantityBackOrdered,omitempty" json:"quantityBackOrdered,omitempty"`
//	QuantityBilled float64 `xml:"quantityBilled,omitempty" json:"quantityBilled,omitempty"`
//	QuantityCommitted float64 `xml:"quantityCommitted,omitempty" json:"quantityCommitted,omitempty"`
//	QuantityFulfilled float64 `xml:"quantityFulfilled,omitempty" json:"quantityFulfilled,omitempty"`
//	QuantityPacked float64 `xml:"quantityPacked,omitempty" json:"quantityPacked,omitempty"`
//	QuantityPicked float64 `xml:"quantityPicked,omitempty" json:"quantityPicked,omitempty"`
//	Tax1Amt float64 `xml:"tax1Amt,omitempty" json:"tax1Amt,omitempty"`
//	TaxCode *RecordRef `xml:"taxCode,omitempty" json:"taxCode,omitempty"`
//	TaxRate1 float64 `xml:"taxRate1,omitempty" json:"taxRate1,omitempty"`
//	TaxRate2 float64 `xml:"taxRate2,omitempty" json:"taxRate2,omitempty"`
//	GiftCertFrom string `xml:"giftCertFrom,omitempty" json:"giftCertFrom,omitempty"`
//	GiftCertRecipientName string `xml:"giftCertRecipientName,omitempty" json:"giftCertRecipientName,omitempty"`
//	GiftCertRecipientEmail string `xml:"giftCertRecipientEmail,omitempty" json:"giftCertRecipientEmail,omitempty"`
//	GiftCertMessage string `xml:"giftCertMessage,omitempty" json:"giftCertMessage,omitempty"`
//	GiftCertNumber string `xml:"giftCertNumber,omitempty" json:"giftCertNumber,omitempty"`
//	ShipGroup int64 `xml:"shipGroup,omitempty" json:"shipGroup,omitempty"`
//	ItemIsFulfilled bool `xml:"itemIsFulfilled,omitempty" json:"itemIsFulfilled,omitempty"`
//	ShipAddress *RecordRef `xml:"shipAddress,omitempty" json:"shipAddress,omitempty"`
//	ShipMethod *RecordRef `xml:"shipMethod,omitempty" json:"shipMethod,omitempty"`
//	VsoeSopGroup *VsoeSopGroup `xml:"vsoeSopGroup,omitempty" json:"vsoeSopGroup,omitempty"`
//	VsoeIsEstimate bool `xml:"vsoeIsEstimate,omitempty" json:"vsoeIsEstimate,omitempty"`
//	VsoePrice float64 `xml:"vsoePrice,omitempty" json:"vsoePrice,omitempty"`
//	VsoeAmount float64 `xml:"vsoeAmount,omitempty" json:"vsoeAmount,omitempty"`
//	VsoeAllocation float64 `xml:"vsoeAllocation,omitempty" json:"vsoeAllocation,omitempty"`
//	VsoeDeferral *VsoeDeferral `xml:"vsoeDeferral,omitempty" json:"vsoeDeferral,omitempty"`
//	VsoePermitDiscount *VsoePermitDiscount `xml:"vsoePermitDiscount,omitempty" json:"vsoePermitDiscount,omitempty"`
//	VsoeDelivered bool `xml:"vsoeDelivered,omitempty" json:"vsoeDelivered,omitempty"`
//	ExpectedShipDate soap.XSDDateTime `xml:"expectedShipDate,omitempty" json:"expectedShipDate,omitempty"`
//	NoAutoAssignLocation bool `xml:"noAutoAssignLocation,omitempty" json:"noAutoAssignLocation,omitempty"`
//	LocationAutoAssigned bool `xml:"locationAutoAssigned,omitempty" json:"locationAutoAssigned,omitempty"`
//	TaxDetailsReference string `xml:"taxDetailsReference,omitempty" json:"taxDetailsReference,omitempty"`
//	ChargeType *RecordRef `xml:"chargeType,omitempty" json:"chargeType,omitempty"`
//	OrderAllocationStrategy *RecordRef `xml:"orderAllocationStrategy,omitempty" json:"orderAllocationStrategy,omitempty"`
//	RequestedDate soap.XSDDateTime `xml:"requestedDate,omitempty" json:"requestedDate,omitempty"`
	CustomFieldList *CustomFieldList `xml:"customFieldList,omitempty" json:"customFieldList,omitempty"`
}

type SalesOrderItemList struct {
	XMLName xml.Name `xml:"itemList" json:"-"`
	Item []*SalesOrderItem `xml:"item,omitempty" json:"item,omitempty"`
	ReplaceAll bool `xml:"replaceAll,attr,omitempty" json:"replaceAll,omitempty"`
}
