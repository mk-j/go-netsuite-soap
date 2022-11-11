//GENERATED via go2wsdl, (and modified for naming conflicts, and functionality)

package netsuite

import (
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
)

type Customer_AddRequest struct {
	XMLName xml.Name `xml:"add"`
	XmlNSXSI   string   `xml:"xmlns:xsi,attr,omitempty"` 
	XmlNSPC   string   `xml:"xmlns:platformCore,attr,omitempty"` 
	XmlNS1   string   `xml:"xmlns:ns1,attr,omitempty"` 
	Record *Customer `xml:"record,omitempty" json:"record,omitempty"`
}

type Customer_AddResponse struct {
	XMLName xml.Name `xml:"addResponse"`
	WriteResponse *Customer_WriteResponse `xml:"writeResponse,omitempty" json:"writeResponse,omitempty"`
}

type Customer_WriteResponse struct {
	XMLName xml.Name `xml:"writeResponse"`
	Status *Status `xml:"status,omitempty" json:"status,omitempty"`
	BaseRef *BaseRef `xml:"baseRef,omitempty" json:"baseRef,omitempty"`
}

//----------------------------------------------------

type Customer_GetRequest struct {
	XMLName xml.Name `xml:"get"`
	XmlNSXSI   string   `xml:"xmlns:xsi,attr,omitempty"` 
	XmlNSPC   string   `xml:"xmlns:platformCore,attr,omitempty"` 
	BaseRef *BaseRef `xml:"baseRef,omitempty" json:"baseRef,omitempty"`
}

type Customer_GetResponse struct {
	//XMLName xml.Name `xml:"urn:messages_2022_1.platform.webservices.netsuite.com getResponse"`
	XMLName xml.Name `xml:"getResponse"`
	ReadResponse *Customer_ReadResponse `xml:"readResponse,omitempty" json:"readResponse,omitempty"`
}
//----------------------------------------------------

type Customer_ReadResponse struct {
	XMLName xml.Name `xml:"readResponse"`
	Status *Status `xml:"status,omitempty" json:"status,omitempty"`
	Record *Customer `xml:"record,omitempty" json:"record,omitempty"`
}

//Leave the below fields, commented out, we can comment them back in, as needed
type Customer struct {
	XsiType string `xml:"xsi:type,attr,omitempty"  json:"-"`
	CustomForm *RecordRef `xml:"customForm,omitempty" json:"customForm,omitempty"`
	EntityId string `xml:"entityId,omitempty" json:"entityId,omitempty"`
	AltName string `xml:"altName,omitempty" json:"altName,omitempty"`
	IsPerson bool `xml:"isPerson,omitempty" json:"isPerson,omitempty"`
	PhoneticName string `xml:"phoneticName,omitempty" json:"phoneticName,omitempty"`
	Salutation string `xml:"salutation,omitempty" json:"salutation,omitempty"`
	FirstName string `xml:"firstName,omitempty" json:"firstName,omitempty"`
	MiddleName string `xml:"middleName,omitempty" json:"middleName,omitempty"`
	LastName string `xml:"lastName,omitempty" json:"lastName,omitempty"`
	CompanyName string `xml:"companyName,omitempty" json:"companyName,omitempty"`
	EntityStatus *RecordRef `xml:"entityStatus,omitempty" json:"entityStatus,omitempty"`
//	Parent *RecordRef `xml:"parent,omitempty" json:"parent,omitempty"`
	Phone string `xml:"phone,omitempty" json:"phone,omitempty"`
	Fax string `xml:"fax,omitempty" json:"fax,omitempty"`
	Email string `xml:"email,omitempty" json:"email,omitempty"`
	Url string `xml:"url,omitempty" json:"url,omitempty"`
	DefaultAddress string `xml:"defaultAddress,omitempty" json:"defaultAddress,omitempty"`
	IsInactive bool `xml:"isInactive,omitempty" json:"isInactive,omitempty"`
//	Category *RecordRef `xml:"category,omitempty" json:"category,omitempty"`
	Title string `xml:"title,omitempty" json:"title,omitempty"`
	PrintOnCheckAs string `xml:"printOnCheckAs,omitempty" json:"printOnCheckAs,omitempty"`
	AltPhone string `xml:"altPhone,omitempty" json:"altPhone,omitempty"`
	HomePhone string `xml:"homePhone,omitempty" json:"homePhone,omitempty"`
	MobilePhone string `xml:"mobilePhone,omitempty" json:"mobilePhone,omitempty"`
	AltEmail string `xml:"altEmail,omitempty" json:"altEmail,omitempty"`
//	Language *Language `xml:"language,omitempty" json:"language,omitempty"`
	Comments string `xml:"comments,omitempty" json:"comments,omitempty"`
//	NumberFormat *CustomerNumberFormat `xml:"numberFormat,omitempty" json:"numberFormat,omitempty"`
//	NegativeNumberFormat *CustomerNegativeNumberFormat `xml:"negativeNumberFormat,omitempty" json:"negativeNumberFormat,omitempty"`
	DateCreated soap.XSDDateTime `xml:"dateCreated,omitempty" json:"dateCreated,omitempty"`
//	Image *RecordRef `xml:"image,omitempty" json:"image,omitempty"`
//	EmailPreference *EmailPreference `xml:"emailPreference,omitempty" json:"emailPreference,omitempty"`
	Subsidiary *RecordRef `xml:"subsidiary,omitempty" json:"subsidiary,omitempty"`
//	RepresentingSubsidiary *RecordRef `xml:"representingSubsidiary,omitempty" json:"representingSubsidiary,omitempty"`
//	SalesRep *RecordRef `xml:"salesRep,omitempty" json:"salesRep,omitempty"`
//	Territory *RecordRef `xml:"territory,omitempty" json:"territory,omitempty"`
//	ContribPct string `xml:"contribPct,omitempty" json:"contribPct,omitempty"`
//	Partner *RecordRef `xml:"partner,omitempty" json:"partner,omitempty"`
//	SalesGroup *RecordRef `xml:"salesGroup,omitempty" json:"salesGroup,omitempty"`
//	VatRegNumber string `xml:"vatRegNumber,omitempty" json:"vatRegNumber,omitempty"`
//	AccountNumber string `xml:"accountNumber,omitempty" json:"accountNumber,omitempty"`
//	TaxExempt bool `xml:"taxExempt,omitempty" json:"taxExempt,omitempty"`
	Terms *RecordRef `xml:"terms,omitempty" json:"terms,omitempty"`
//	CreditLimit float64 `xml:"creditLimit,omitempty" json:"creditLimit,omitempty"`
//	CreditHoldOverride *CustomerCreditHoldOverride `xml:"creditHoldOverride,omitempty" json:"creditHoldOverride,omitempty"`
//	MonthlyClosing *CustomerMonthlyClosing `xml:"monthlyClosing,omitempty" json:"monthlyClosing,omitempty"`
//	OverrideCurrencyFormat bool `xml:"overrideCurrencyFormat,omitempty" json:"overrideCurrencyFormat,omitempty"`
//	DisplaySymbol string `xml:"displaySymbol,omitempty" json:"displaySymbol,omitempty"`
//	SymbolPlacement *CurrencySymbolPlacement `xml:"symbolPlacement,omitempty" json:"symbolPlacement,omitempty"`
//	Balance float64 `xml:"balance,omitempty" json:"balance,omitempty"`
//	OverdueBalance float64 `xml:"overdueBalance,omitempty" json:"overdueBalance,omitempty"`
//	DaysOverdue int64 `xml:"daysOverdue,omitempty" json:"daysOverdue,omitempty"`
//	UnbilledOrders float64 `xml:"unbilledOrders,omitempty" json:"unbilledOrders,omitempty"`
//	ConsolUnbilledOrders float64 `xml:"consolUnbilledOrders,omitempty" json:"consolUnbilledOrders,omitempty"`
//	ConsolOverdueBalance float64 `xml:"consolOverdueBalance,omitempty" json:"consolOverdueBalance,omitempty"`
//	ConsolDepositBalance float64 `xml:"consolDepositBalance,omitempty" json:"consolDepositBalance,omitempty"`
//	ConsolBalance float64 `xml:"consolBalance,omitempty" json:"consolBalance,omitempty"`
//	ConsolAging float64 `xml:"consolAging,omitempty" json:"consolAging,omitempty"`
//	ConsolAging1 float64 `xml:"consolAging1,omitempty" json:"consolAging1,omitempty"`
//	ConsolAging2 float64 `xml:"consolAging2,omitempty" json:"consolAging2,omitempty"`
//	ConsolAging3 float64 `xml:"consolAging3,omitempty" json:"consolAging3,omitempty"`
//	ConsolAging4 float64 `xml:"consolAging4,omitempty" json:"consolAging4,omitempty"`
//	ConsolDaysOverdue int64 `xml:"consolDaysOverdue,omitempty" json:"consolDaysOverdue,omitempty"`
//	PriceLevel *RecordRef `xml:"priceLevel,omitempty" json:"priceLevel,omitempty"`
//	Currency *RecordRef `xml:"currency,omitempty" json:"currency,omitempty"`
//	PrefCCProcessor *RecordRef `xml:"prefCCProcessor,omitempty" json:"prefCCProcessor,omitempty"`
//	DepositBalance float64 `xml:"depositBalance,omitempty" json:"depositBalance,omitempty"`
//	ShipComplete bool `xml:"shipComplete,omitempty" json:"shipComplete,omitempty"`
//	Taxable bool `xml:"taxable,omitempty" json:"taxable,omitempty"`
//	TaxItem *RecordRef `xml:"taxItem,omitempty" json:"taxItem,omitempty"`
	ResaleNumber string `xml:"resaleNumber,omitempty" json:"resaleNumber,omitempty"`
//	Aging float64 `xml:"aging,omitempty" json:"aging,omitempty"`
//	Aging1 float64 `xml:"aging1,omitempty" json:"aging1,omitempty"`
//	Aging2 float64 `xml:"aging2,omitempty" json:"aging2,omitempty"`
//	Aging3 float64 `xml:"aging3,omitempty" json:"aging3,omitempty"`
//	Aging4 float64 `xml:"aging4,omitempty" json:"aging4,omitempty"`
	StartDate soap.XSDDateTime `xml:"startDate,omitempty" json:"startDate,omitempty"`
//	AlcoholRecipientType *AlcoholRecipientType `xml:"alcoholRecipientType,omitempty" json:"alcoholRecipientType,omitempty"`
	EndDate soap.XSDDateTime `xml:"endDate,omitempty" json:"endDate,omitempty"`
//	ReminderDays int64 `xml:"reminderDays,omitempty" json:"reminderDays,omitempty"`
//	ShippingItem *RecordRef `xml:"shippingItem,omitempty" json:"shippingItem,omitempty"`
//	ThirdPartyAcct string `xml:"thirdPartyAcct,omitempty" json:"thirdPartyAcct,omitempty"`
//	ThirdPartyZipcode string `xml:"thirdPartyZipcode,omitempty" json:"thirdPartyZipcode,omitempty"`
//	ThirdPartyCountry *Country `xml:"thirdPartyCountry,omitempty" json:"thirdPartyCountry,omitempty"`
//	ThirdPartyCarrier *CustomerThirdPartyCarrier `xml:"thirdPartyCarrier,omitempty" json:"thirdPartyCarrier,omitempty"`
//	GiveAccess bool `xml:"giveAccess,omitempty" json:"giveAccess,omitempty"`
//	EstimatedBudget float64 `xml:"estimatedBudget,omitempty" json:"estimatedBudget,omitempty"`
//	AccessRole *RecordRef `xml:"accessRole,omitempty" json:"accessRole,omitempty"`
//	SendEmail bool `xml:"sendEmail,omitempty" json:"sendEmail,omitempty"`
//	AssignedWebSite *RecordRef `xml:"assignedWebSite,omitempty" json:"assignedWebSite,omitempty"`
//	Password string `xml:"password,omitempty" json:"password,omitempty"`
//	Password2 string `xml:"password2,omitempty" json:"password2,omitempty"`
//	RequirePwdChange bool `xml:"requirePwdChange,omitempty" json:"requirePwdChange,omitempty"`
//	CampaignCategory *RecordRef `xml:"campaignCategory,omitempty" json:"campaignCategory,omitempty"`
//	SourceWebSite *RecordRef `xml:"sourceWebSite,omitempty" json:"sourceWebSite,omitempty"`
//	LeadSource *RecordRef `xml:"leadSource,omitempty" json:"leadSource,omitempty"`
//	ReceivablesAccount *RecordRef `xml:"receivablesAccount,omitempty" json:"receivablesAccount,omitempty"`
//	DrAccount *RecordRef `xml:"drAccount,omitempty" json:"drAccount,omitempty"`
//	FxAccount *RecordRef `xml:"fxAccount,omitempty" json:"fxAccount,omitempty"`
//	DefaultOrderPriority float64 `xml:"defaultOrderPriority,omitempty" json:"defaultOrderPriority,omitempty"`
//	WebLead string `xml:"webLead,omitempty" json:"webLead,omitempty"`
//	Referrer string `xml:"referrer,omitempty" json:"referrer,omitempty"`
//	Keywords string `xml:"keywords,omitempty" json:"keywords,omitempty"`
//	ClickStream string `xml:"clickStream,omitempty" json:"clickStream,omitempty"`
//	LastPageVisited string `xml:"lastPageVisited,omitempty" json:"lastPageVisited,omitempty"`
//	Visits int64 `xml:"visits,omitempty" json:"visits,omitempty"`
//	FirstVisit soap.XSDDateTime `xml:"firstVisit,omitempty" json:"firstVisit,omitempty"`
//	LastVisit soap.XSDDateTime `xml:"lastVisit,omitempty" json:"lastVisit,omitempty"`
//	BillPay bool `xml:"billPay,omitempty" json:"billPay,omitempty"`
//	OpeningBalance float64 `xml:"openingBalance,omitempty" json:"openingBalance,omitempty"`
//	LastModifiedDate soap.XSDDateTime `xml:"lastModifiedDate,omitempty" json:"lastModifiedDate,omitempty"`
//	OpeningBalanceDate soap.XSDDateTime `xml:"openingBalanceDate,omitempty" json:"openingBalanceDate,omitempty"`
//	OpeningBalanceAccount *RecordRef `xml:"openingBalanceAccount,omitempty" json:"openingBalanceAccount,omitempty"`
//	Stage *CustomerStage `xml:"stage,omitempty" json:"stage,omitempty"`
//	EmailTransactions bool `xml:"emailTransactions,omitempty" json:"emailTransactions,omitempty"`
//	PrintTransactions bool `xml:"printTransactions,omitempty" json:"printTransactions,omitempty"`
//	FaxTransactions bool `xml:"faxTransactions,omitempty" json:"faxTransactions,omitempty"`
//	DefaultTaxReg *RecordRef `xml:"defaultTaxReg,omitempty" json:"defaultTaxReg,omitempty"`
//	SyncPartnerTeams bool `xml:"syncPartnerTeams,omitempty" json:"syncPartnerTeams,omitempty"`
//	IsBudgetApproved bool `xml:"isBudgetApproved,omitempty" json:"isBudgetApproved,omitempty"`
//	GlobalSubscriptionStatus *GlobalSubscriptionStatus `xml:"globalSubscriptionStatus,omitempty" json:"globalSubscriptionStatus,omitempty"`
//	SalesReadiness *RecordRef `xml:"salesReadiness,omitempty" json:"salesReadiness,omitempty"`
//	SalesTeamList *CustomerSalesTeamList `xml:"salesTeamList,omitempty" json:"salesTeamList,omitempty"`
//	BuyingReason *RecordRef `xml:"buyingReason,omitempty" json:"buyingReason,omitempty"`
//	DownloadList *CustomerDownloadList `xml:"downloadList,omitempty" json:"downloadList,omitempty"`
//	BuyingTimeFrame *RecordRef `xml:"buyingTimeFrame,omitempty" json:"buyingTimeFrame,omitempty"`
//	AddressbookList *CustomerAddressbookList `xml:"addressbookList,omitempty" json:"addressbookList,omitempty"`
//	SubscriptionsList *SubscriptionsList `xml:"subscriptionsList,omitempty" json:"subscriptionsList,omitempty"`
//	ContactRolesList *ContactAccessRolesList `xml:"contactRolesList,omitempty" json:"contactRolesList,omitempty"`
//	CurrencyList *CustomerCurrencyList `xml:"currencyList,omitempty" json:"currencyList,omitempty"`
//	CreditCardsList *CustomerCreditCardsList `xml:"creditCardsList,omitempty" json:"creditCardsList,omitempty"`
//	PartnersList *CustomerPartnersList `xml:"partnersList,omitempty" json:"partnersList,omitempty"`
//	GroupPricingList *CustomerGroupPricingList `xml:"groupPricingList,omitempty" json:"groupPricingList,omitempty"`
//	ItemPricingList *CustomerItemPricingList `xml:"itemPricingList,omitempty" json:"itemPricingList,omitempty"`
//	TaxRegistrationList *CustomerTaxRegistrationList `xml:"taxRegistrationList,omitempty" json:"taxRegistrationList,omitempty"`
//	DefaultAllocationStrategy *RecordRef `xml:"defaultAllocationStrategy,omitempty" json:"defaultAllocationStrategy,omitempty"`
	CustomFieldList *CustomFieldList `xml:"customFieldList,omitempty" json:"customFieldList,omitempty"`
	InternalId string `xml:"internalId,attr,omitempty" json:"internalId,omitempty"`
	ExternalId string `xml:"externalId,attr,omitempty" json:"externalId,omitempty"`
}

