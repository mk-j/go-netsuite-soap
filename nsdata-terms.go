//GENERATED via go2wsdl, (and modified for naming conflicts, and functionality)

package netsuite

import (
	"encoding/xml"
	//"github.com/hooklift/gowsdl/soap"
)
type SearchRequest struct {
	XMLName xml.Name `xml:"search"`
	XmlNSXSI   string   `xml:"xmlns:xsi,attr,omitempty"` 
	XmlNSPC   string   `xml:"xmlns:platformCore,attr,omitempty"` 
	SearchRecord *SearchRecord `xml:"searchRecord,omitempty" json:"searchRecord,omitempty"`
}
type SearchRecord struct {
	XMLName xml.Name `xml:"searchRecord"`
	XmlNSXSI string `xml:"xmlns:xsi,attr,omitempty"` 
	XsiType string `xml:"xsi:type,attr,omitempty"  json:"-"`
	//*TermSearchBasic
}
type TermSearchBasic struct {
	XMLName xml.Name `xml:"basic"`
	//*SearchRecordBasic
	//DateDriven *SearchBooleanField `xml:"dateDriven,omitempty" json:"dateDriven,omitempty"`
	//DayDiscountExpires *SearchLongField `xml:"dayDiscountExpires,omitempty" json:"dayDiscountExpires,omitempty"`
	//DayOfMonthNetDue *SearchLongField `xml:"dayOfMonthNetDue,omitempty" json:"dayOfMonthNetDue,omitempty"`
	//DaysUntilExpiry *SearchLongField `xml:"daysUntilExpiry,omitempty" json:"daysUntilExpiry,omitempty"`
	//DaysUntilNetDue *SearchLongField `xml:"daysUntilNetDue,omitempty" json:"daysUntilNetDue,omitempty"`
	//DiscountPercent *SearchDoubleField `xml:"discountPercent,omitempty" json:"discountPercent,omitempty"`
	//DiscountPercentDateDriven *SearchDoubleField `xml:"discountPercentDateDriven,omitempty" json:"discountPercentDateDriven,omitempty"`
	//DueNextMonthIfWithinDays *SearchLongField `xml:"dueNextMonthIfWithinDays,omitempty" json:"dueNextMonthIfWithinDays,omitempty"`
	//ExternalId *SearchMultiSelectField `xml:"externalId,omitempty" json:"externalId,omitempty"`
	//ExternalIdString *SearchStringField `xml:"externalIdString,omitempty" json:"externalIdString,omitempty"`
	//Installment *SearchBooleanField `xml:"installment,omitempty" json:"installment,omitempty"`
	//InternalId *SearchMultiSelectField `xml:"internalId,omitempty" json:"internalId,omitempty"`
	//InternalIdNumber *SearchLongField `xml:"internalIdNumber,omitempty" json:"internalIdNumber,omitempty"`
	//IsInactive *SearchBooleanField `xml:"isInactive,omitempty" json:"isInactive,omitempty"`
	//Name *SearchStringField `xml:"name,omitempty" json:"name,omitempty"`
	//Preferred *SearchBooleanField `xml:"preferred,omitempty" json:"preferred,omitempty"`
	//RecurrenceCount *SearchLongField `xml:"recurrenceCount,omitempty" json:"recurrenceCount,omitempty"`
	//RecurrenceFrequency *SearchEnumMultiSelectField `xml:"recurrenceFrequency,omitempty" json:"recurrenceFrequency,omitempty"`
	//RepeatEvery *SearchLongField `xml:"repeatEvery,omitempty" json:"repeatEvery,omitempty"`
	//SplitEvenly *SearchBooleanField `xml:"splitEvenly,omitempty" json:"splitEvenly,omitempty"`
}

type SearchResponse struct {
	XMLName xml.Name `xml:"searchResponse"`
	SearchResult *SearchResult `xml:"searchResult,omitempty" json:"searchResult,omitempty"`
}
type SearchResult struct {
	XMLName xml.Name `xml:"searchResult"`
	Status *Status `xml:"status,omitempty" json:"status,omitempty"`
	TotalRecords int32 `xml:"totalRecords,omitempty" json:"totalRecords,omitempty"`
	PageSize int32 `xml:"pageSize,omitempty" json:"pageSize,omitempty"`
	TotalPages int32 `xml:"totalPages,omitempty" json:"totalPages,omitempty"`
	PageIndex int32 `xml:"pageIndex,omitempty" json:"pageIndex,omitempty"`
	SearchId string `xml:"searchId,omitempty" json:"searchId,omitempty"`
	RecordList *TermRecordList `xml:"recordList,omitempty" json:"recordList,omitempty"`
	//SearchRowList *SearchRowList `xml:"searchRowList,omitempty" json:"searchRowList,omitempty"`
}
type TermRecordList struct {
	XMLName xml.Name `xml:"recordList"`
	TermRecord []*Term `xml:"record,omitempty" json:"record,omitempty"`
}
type Term struct {
	//*Record
	Name string `xml:"name,omitempty" json:"name,omitempty"`
	DateDriven bool `xml:"dateDriven,omitempty" json:"dateDriven,omitempty"`
	DaysUntilNetDue int64 `xml:"daysUntilNetDue,omitempty" json:"daysUntilNetDue,omitempty"`
	DiscountPercent float64 `xml:"discountPercent,omitempty" json:"discountPercent,omitempty"`
	DaysUntilExpiry int64 `xml:"daysUntilExpiry,omitempty" json:"daysUntilExpiry,omitempty"`
	DayOfMonthNetDue int64 `xml:"dayOfMonthNetDue,omitempty" json:"dayOfMonthNetDue,omitempty"`
	DueNextMonthIfWithinDays int64 `xml:"dueNextMonthIfWithinDays,omitempty" json:"dueNextMonthIfWithinDays,omitempty"`
	DiscountPercentDateDriven float64 `xml:"discountPercentDateDriven,omitempty" json:"discountPercentDateDriven,omitempty"`
	DayDiscountExpires int64 `xml:"dayDiscountExpires,omitempty" json:"dayDiscountExpires,omitempty"`
	Preferred bool `xml:"preferred,omitempty" json:"preferred,omitempty"`
	IsInactive bool `xml:"isInactive,omitempty" json:"isInactive,omitempty"`
	Installment bool `xml:"installment,omitempty" json:"installment,omitempty"`
//	RecurrenceFrequency *TermRecurrenceFrequency `xml:"recurrenceFrequency,omitempty" json:"recurrenceFrequency,omitempty"`
	RecurrenceCount int64 `xml:"recurrenceCount,omitempty" json:"recurrenceCount,omitempty"`
	RepeatEvery int64 `xml:"repeatEvery,omitempty" json:"repeatEvery,omitempty"`
	SplitEvenly bool `xml:"splitEvenly,omitempty" json:"splitEvenly,omitempty"`
	//PercentagesList *TermPercentagesList `xml:"percentagesList,omitempty" json:"percentagesList,omitempty"`
	InternalId string `xml:"internalId,attr,omitempty" json:"internalId,omitempty"`
	ExternalId string `xml:"externalId,attr,omitempty" json:"externalId,omitempty"`
}



//NOT REALLY USED................
/*
type SearchBooleanField struct {
	SearchValue bool `xml:"searchValue,omitempty" json:"searchValue,omitempty"`
}
type SearchStringField struct {
	SearchValue string `xml:"searchValue,omitempty" json:"searchValue,omitempty"`
	Operator *SearchStringFieldOperator `xml:"operator,attr,omitempty" json:"operator,omitempty"`
}
type SearchLongField struct {
	SearchValue int64 `xml:"searchValue,omitempty" json:"searchValue,omitempty"`
	SearchValue2 int64 `xml:"searchValue2,omitempty" json:"searchValue2,omitempty"`
	Operator *SearchLongFieldOperator `xml:"operator,attr,omitempty" json:"operator,omitempty"`
}
type SearchTextNumberField struct {
	SearchValue string `xml:"searchValue,omitempty" json:"searchValue,omitempty"`
	SearchValue2 string `xml:"searchValue2,omitempty" json:"searchValue2,omitempty"`
	Operator *SearchTextNumberFieldOperator `xml:"operator,attr,omitempty" json:"operator,omitempty"`
}
type SearchDoubleField struct {
	SearchValue float64 `xml:"searchValue,omitempty" json:"searchValue,omitempty"`
	SearchValue2 float64 `xml:"searchValue2,omitempty" json:"searchValue2,omitempty"`
	Operator *SearchDoubleFieldOperator `xml:"operator,attr,omitempty" json:"operator,omitempty"`
}
type SearchDateField struct {
	PredefinedSearchValue *SearchDate `xml:"predefinedSearchValue,omitempty" json:"predefinedSearchValue,omitempty"`
	SearchValue soap.XSDDateTime `xml:"searchValue,omitempty" json:"searchValue,omitempty"`
	SearchValue2 soap.XSDDateTime `xml:"searchValue2,omitempty" json:"searchValue2,omitempty"`
	Operator *SearchDateFieldOperator `xml:"operator,attr,omitempty" json:"operator,omitempty"`
}
type SearchEnumMultiSelectField struct {
	SearchValue []string `xml:"searchValue,omitempty" json:"searchValue,omitempty"`
	Operator *SearchEnumMultiSelectFieldOperator `xml:"operator,attr,omitempty" json:"operator,omitempty"`
}
type SearchMultiSelectField struct {
	SearchValue []*RecordRef `xml:"searchValue,omitempty" json:"searchValue,omitempty"`
	Operator *SearchMultiSelectFieldOperator `xml:"operator,attr,omitempty" json:"operator,omitempty"`
}
type SearchStringFieldOperator string
const (
	SearchStringFieldOperatorContains SearchStringFieldOperator = "contains"
	SearchStringFieldOperatorDoesNotContain SearchStringFieldOperator = "doesNotContain"
	SearchStringFieldOperatorDoesNotStartWith SearchStringFieldOperator = "doesNotStartWith"
	SearchStringFieldOperatorEmpty SearchStringFieldOperator = "empty"
	SearchStringFieldOperatorHasKeywords SearchStringFieldOperator = "hasKeywords"
	SearchStringFieldOperatorIs SearchStringFieldOperator = "is"
	SearchStringFieldOperatorIsNot SearchStringFieldOperator = "isNot"
	SearchStringFieldOperatorNotEmpty SearchStringFieldOperator = "notEmpty"
	SearchStringFieldOperatorStartsWith SearchStringFieldOperator = "startsWith"
)
type SearchLongFieldOperator string
const (
	SearchLongFieldOperatorBetween SearchLongFieldOperator = "between"
	SearchLongFieldOperatorEmpty SearchLongFieldOperator = "empty"
	SearchLongFieldOperatorEqualTo SearchLongFieldOperator = "equalTo"
	SearchLongFieldOperatorGreaterThan SearchLongFieldOperator = "greaterThan"
	SearchLongFieldOperatorGreaterThanOrEqualTo SearchLongFieldOperator = "greaterThanOrEqualTo"
	SearchLongFieldOperatorLessThan SearchLongFieldOperator = "lessThan"
	SearchLongFieldOperatorLessThanOrEqualTo SearchLongFieldOperator = "lessThanOrEqualTo"
	SearchLongFieldOperatorNotBetween SearchLongFieldOperator = "notBetween"
	SearchLongFieldOperatorNotEmpty SearchLongFieldOperator = "notEmpty"
	SearchLongFieldOperatorNotEqualTo SearchLongFieldOperator = "notEqualTo"
	SearchLongFieldOperatorNotGreaterThan SearchLongFieldOperator = "notGreaterThan"
	SearchLongFieldOperatorNotGreaterThanOrEqualTo SearchLongFieldOperator = "notGreaterThanOrEqualTo"
	SearchLongFieldOperatorNotLessThan SearchLongFieldOperator = "notLessThan"
	SearchLongFieldOperatorNotLessThanOrEqualTo SearchLongFieldOperator = "notLessThanOrEqualTo"
)
type SearchTextNumberFieldOperator string
const (
	SearchTextNumberFieldOperatorBetween SearchTextNumberFieldOperator = "between"
	SearchTextNumberFieldOperatorEmpty SearchTextNumberFieldOperator = "empty"
	SearchTextNumberFieldOperatorEqualTo SearchTextNumberFieldOperator = "equalTo"
	SearchTextNumberFieldOperatorGreaterThan SearchTextNumberFieldOperator = "greaterThan"
	SearchTextNumberFieldOperatorGreaterThanOrEqualTo SearchTextNumberFieldOperator = "greaterThanOrEqualTo"
	SearchTextNumberFieldOperatorLessThan SearchTextNumberFieldOperator = "lessThan"
	SearchTextNumberFieldOperatorLessThanOrEqualTo SearchTextNumberFieldOperator = "lessThanOrEqualTo"
	SearchTextNumberFieldOperatorNotBetween SearchTextNumberFieldOperator = "notBetween"
	SearchTextNumberFieldOperatorNotEmpty SearchTextNumberFieldOperator = "notEmpty"
	SearchTextNumberFieldOperatorNotEqualTo SearchTextNumberFieldOperator = "notEqualTo"
	SearchTextNumberFieldOperatorNotGreaterThan SearchTextNumberFieldOperator = "notGreaterThan"
	SearchTextNumberFieldOperatorNotGreaterThanOrEqualTo SearchTextNumberFieldOperator = "notGreaterThanOrEqualTo"
	SearchTextNumberFieldOperatorNotLessThan SearchTextNumberFieldOperator = "notLessThan"
	SearchTextNumberFieldOperatorNotLessThanOrEqualTo SearchTextNumberFieldOperator = "notLessThanOrEqualTo"
)
type SearchDoubleFieldOperator string
const (
	SearchDoubleFieldOperatorBetween SearchDoubleFieldOperator = "between"
	SearchDoubleFieldOperatorEmpty SearchDoubleFieldOperator = "empty"
	SearchDoubleFieldOperatorEqualTo SearchDoubleFieldOperator = "equalTo"
	SearchDoubleFieldOperatorGreaterThan SearchDoubleFieldOperator = "greaterThan"
	SearchDoubleFieldOperatorGreaterThanOrEqualTo SearchDoubleFieldOperator = "greaterThanOrEqualTo"
	SearchDoubleFieldOperatorLessThan SearchDoubleFieldOperator = "lessThan"
	SearchDoubleFieldOperatorLessThanOrEqualTo SearchDoubleFieldOperator = "lessThanOrEqualTo"
	SearchDoubleFieldOperatorNotBetween SearchDoubleFieldOperator = "notBetween"
	SearchDoubleFieldOperatorNotEmpty SearchDoubleFieldOperator = "notEmpty"
	SearchDoubleFieldOperatorNotEqualTo SearchDoubleFieldOperator = "notEqualTo"
	SearchDoubleFieldOperatorNotGreaterThan SearchDoubleFieldOperator = "notGreaterThan"
	SearchDoubleFieldOperatorNotGreaterThanOrEqualTo SearchDoubleFieldOperator = "notGreaterThanOrEqualTo"
	SearchDoubleFieldOperatorNotLessThan SearchDoubleFieldOperator = "notLessThan"
	SearchDoubleFieldOperatorNotLessThanOrEqualTo SearchDoubleFieldOperator = "notLessThanOrEqualTo"
)
type SearchDateFieldOperator string
const (
	SearchDateFieldOperatorAfter SearchDateFieldOperator = "after"
	SearchDateFieldOperatorBefore SearchDateFieldOperator = "before"
	SearchDateFieldOperatorEmpty SearchDateFieldOperator = "empty"
	SearchDateFieldOperatorNotAfter SearchDateFieldOperator = "notAfter"
	SearchDateFieldOperatorNotBefore SearchDateFieldOperator = "notBefore"
	SearchDateFieldOperatorNotEmpty SearchDateFieldOperator = "notEmpty"
	SearchDateFieldOperatorNotOn SearchDateFieldOperator = "notOn"
	SearchDateFieldOperatorNotOnOrAfter SearchDateFieldOperator = "notOnOrAfter"
	SearchDateFieldOperatorNotOnOrBefore SearchDateFieldOperator = "notOnOrBefore"
	SearchDateFieldOperatorNotWithin SearchDateFieldOperator = "notWithin"
	SearchDateFieldOperatorOn SearchDateFieldOperator = "on"
	SearchDateFieldOperatorOnOrAfter SearchDateFieldOperator = "onOrAfter"
	SearchDateFieldOperatorOnOrBefore SearchDateFieldOperator = "onOrBefore"
	SearchDateFieldOperatorWithin SearchDateFieldOperator = "within"
)
type SearchEnumMultiSelectFieldOperator string
const (
	SearchEnumMultiSelectFieldOperatorAnyOf SearchEnumMultiSelectFieldOperator = "anyOf"
	SearchEnumMultiSelectFieldOperatorNoneOf SearchEnumMultiSelectFieldOperator = "noneOf"
)
type SearchMultiSelectFieldOperator string
const (
	SearchMultiSelectFieldOperatorAnyOf SearchMultiSelectFieldOperator = "anyOf"
	SearchMultiSelectFieldOperatorNoneOf SearchMultiSelectFieldOperator = "noneOf"
)
type SearchDate string
const (
	SearchDateFiscalHalfBeforeLast SearchDate = "fiscalHalfBeforeLast"
	SearchDateFiscalHalfBeforeLastToDate SearchDate = "fiscalHalfBeforeLastToDate"
	SearchDateFiscalQuarterBeforeLast SearchDate = "fiscalQuarterBeforeLast"
	SearchDateFiscalQuarterBeforeLastToDate SearchDate = "fiscalQuarterBeforeLastToDate"
	SearchDateFiscalYearBeforeLast SearchDate = "fiscalYearBeforeLast"
	SearchDateFiscalYearBeforeLastToDate SearchDate = "fiscalYearBeforeLastToDate"
	SearchDateFiveDaysAgo SearchDate = "fiveDaysAgo"
	SearchDateFiveDaysFromNow SearchDate = "fiveDaysFromNow"
	SearchDateFourDaysAgo SearchDate = "fourDaysAgo"
	SearchDateFourDaysFromNow SearchDate = "fourDaysFromNow"
	SearchDateFourWeeksStartingThisWeek SearchDate = "fourWeeksStartingThisWeek"
	SearchDateLastBusinessWeek SearchDate = "lastBusinessWeek"
	SearchDateLastFiscalHalf SearchDate = "lastFiscalHalf"
	SearchDateLastFiscalHalfOneFiscalYearAgo SearchDate = "lastFiscalHalfOneFiscalYearAgo"
	SearchDateLastFiscalHalfToDate SearchDate = "lastFiscalHalfToDate"
	SearchDateLastFiscalQuarter SearchDate = "lastFiscalQuarter"
	SearchDateLastFiscalQuarterOneFiscalYearAgo SearchDate = "lastFiscalQuarterOneFiscalYearAgo"
	SearchDateLastFiscalQuarterToDate SearchDate = "lastFiscalQuarterToDate"
	SearchDateLastFiscalQuarterTwoFiscalYearsAgo SearchDate = "lastFiscalQuarterTwoFiscalYearsAgo"
	SearchDateLastFiscalYear SearchDate = "lastFiscalYear"
	SearchDateLastFiscalYearToDate SearchDate = "lastFiscalYearToDate"
	SearchDateLastMonth SearchDate = "lastMonth"
	SearchDateLastMonthOneFiscalQuarterAgo SearchDate = "lastMonthOneFiscalQuarterAgo"
	SearchDateLastMonthOneFiscalYearAgo SearchDate = "lastMonthOneFiscalYearAgo"
	SearchDateLastMonthToDate SearchDate = "lastMonthToDate"
	SearchDateLastMonthTwoFiscalQuartersAgo SearchDate = "lastMonthTwoFiscalQuartersAgo"
	SearchDateLastMonthTwoFiscalYearsAgo SearchDate = "lastMonthTwoFiscalYearsAgo"
	SearchDateLastRollingHalf SearchDate = "lastRollingHalf"
	SearchDateLastRollingQuarter SearchDate = "lastRollingQuarter"
	SearchDateLastRollingYear SearchDate = "lastRollingYear"
	SearchDateLastWeek SearchDate = "lastWeek"
	SearchDateLastWeekToDate SearchDate = "lastWeekToDate"
	SearchDateMonthAfterNext SearchDate = "monthAfterNext"
	SearchDateMonthAfterNextToDate SearchDate = "monthAfterNextToDate"
	SearchDateMonthBeforeLast SearchDate = "monthBeforeLast"
	SearchDateMonthBeforeLastToDate SearchDate = "monthBeforeLastToDate"
	SearchDateNextBusinessWeek SearchDate = "nextBusinessWeek"
	SearchDateNextFiscalHalf SearchDate = "nextFiscalHalf"
	SearchDateNextFiscalQuarter SearchDate = "nextFiscalQuarter"
	SearchDateNextFiscalYear SearchDate = "nextFiscalYear"
	SearchDateNextFourWeeks SearchDate = "nextFourWeeks"
	SearchDateNextMonth SearchDate = "nextMonth"
	SearchDateNextOneHalf SearchDate = "nextOneHalf"
	SearchDateNextOneMonth SearchDate = "nextOneMonth"
	SearchDateNextOneQuarter SearchDate = "nextOneQuarter"
	SearchDateNextOneWeek SearchDate = "nextOneWeek"
	SearchDateNextOneYear SearchDate = "nextOneYear"
	SearchDateNextWeek SearchDate = "nextWeek"
	SearchDateNinetyDaysAgo SearchDate = "ninetyDaysAgo"
	SearchDateNinetyDaysFromNow SearchDate = "ninetyDaysFromNow"
	SearchDateOneYearBeforeLast SearchDate = "oneYearBeforeLast"
	SearchDatePreviousFiscalQuartersLastFiscalYear SearchDate = "previousFiscalQuartersLastFiscalYear"
	SearchDatePreviousFiscalQuartersThisFiscalYear SearchDate = "previousFiscalQuartersThisFiscalYear"
	SearchDatePreviousMonthsLastFiscalHalf SearchDate = "previousMonthsLastFiscalHalf"
	SearchDatePreviousMonthsLastFiscalQuarter SearchDate = "previousMonthsLastFiscalQuarter"
	SearchDatePreviousMonthsLastFiscalYear SearchDate = "previousMonthsLastFiscalYear"
	SearchDatePreviousMonthsSameFiscalHalfLastFiscalYear SearchDate = "previousMonthsSameFiscalHalfLastFiscalYear"
	SearchDatePreviousMonthsSameFiscalQuarterLastFiscalYear SearchDate = "previousMonthsSameFiscalQuarterLastFiscalYear"
	SearchDatePreviousMonthsThisFiscalHalf SearchDate = "previousMonthsThisFiscalHalf"
	SearchDatePreviousMonthsThisFiscalQuarter SearchDate = "previousMonthsThisFiscalQuarter"
	SearchDatePreviousMonthsThisFiscalYear SearchDate = "previousMonthsThisFiscalYear"
	SearchDatePreviousOneDay SearchDate = "previousOneDay"
	SearchDatePreviousOneHalf SearchDate = "previousOneHalf"
	SearchDatePreviousOneMonth SearchDate = "previousOneMonth"
	SearchDatePreviousOneQuarter SearchDate = "previousOneQuarter"
	SearchDatePreviousOneWeek SearchDate = "previousOneWeek"
	SearchDatePreviousOneYear SearchDate = "previousOneYear"
	SearchDatePreviousRollingHalf SearchDate = "previousRollingHalf"
	SearchDatePreviousRollingQuarter SearchDate = "previousRollingQuarter"
	SearchDatePreviousRollingYear SearchDate = "previousRollingYear"
	SearchDateSameDayFiscalQuarterBeforeLast SearchDate = "sameDayFiscalQuarterBeforeLast"
	SearchDateSameDayFiscalYearBeforeLast SearchDate = "sameDayFiscalYearBeforeLast"
	SearchDateSameDayLastFiscalQuarter SearchDate = "sameDayLastFiscalQuarter"
	SearchDateSameDayLastFiscalYear SearchDate = "sameDayLastFiscalYear"
	SearchDateSameDayLastMonth SearchDate = "sameDayLastMonth"
	SearchDateSameDayLastWeek SearchDate = "sameDayLastWeek"
	SearchDateSameDayMonthBeforeLast SearchDate = "sameDayMonthBeforeLast"
	SearchDateSameDayWeekBeforeLast SearchDate = "sameDayWeekBeforeLast"
	SearchDateSameFiscalHalfLastFiscalYear SearchDate = "sameFiscalHalfLastFiscalYear"
	SearchDateSameFiscalHalfLastFiscalYearToDate SearchDate = "sameFiscalHalfLastFiscalYearToDate"
	SearchDateSameFiscalQuarterFiscalYearBeforeLast SearchDate = "sameFiscalQuarterFiscalYearBeforeLast"
	SearchDateSameFiscalQuarterLastFiscalYear SearchDate = "sameFiscalQuarterLastFiscalYear"
	SearchDateSameFiscalQuarterLastFiscalYearToDate SearchDate = "sameFiscalQuarterLastFiscalYearToDate"
	SearchDateSameMonthFiscalQuarterBeforeLast SearchDate = "sameMonthFiscalQuarterBeforeLast"
	SearchDateSameMonthFiscalYearBeforeLast SearchDate = "sameMonthFiscalYearBeforeLast"
	SearchDateSameMonthLastFiscalQuarter SearchDate = "sameMonthLastFiscalQuarter"
	SearchDateSameMonthLastFiscalQuarterToDate SearchDate = "sameMonthLastFiscalQuarterToDate"
	SearchDateSameMonthLastFiscalYear SearchDate = "sameMonthLastFiscalYear"
	SearchDateSameMonthLastFiscalYearToDate SearchDate = "sameMonthLastFiscalYearToDate"
	SearchDateSameWeekFiscalYearBeforeLast SearchDate = "sameWeekFiscalYearBeforeLast"
	SearchDateSameWeekLastFiscalYear SearchDate = "sameWeekLastFiscalYear"
	SearchDateSixtyDaysAgo SearchDate = "sixtyDaysAgo"
	SearchDateSixtyDaysFromNow SearchDate = "sixtyDaysFromNow"
	SearchDateStartOfFiscalHalfBeforeLast SearchDate = "startOfFiscalHalfBeforeLast"
	SearchDateStartOfFiscalQuarterBeforeLast SearchDate = "startOfFiscalQuarterBeforeLast"
	SearchDateStartOfFiscalYearBeforeLast SearchDate = "startOfFiscalYearBeforeLast"
	SearchDateStartOfLastBusinessWeek SearchDate = "startOfLastBusinessWeek"
	SearchDateStartOfLastFiscalHalf SearchDate = "startOfLastFiscalHalf"
	SearchDateStartOfLastFiscalHalfOneFiscalYearAgo SearchDate = "startOfLastFiscalHalfOneFiscalYearAgo"
	SearchDateStartOfLastFiscalQuarter SearchDate = "startOfLastFiscalQuarter"
	SearchDateStartOfLastFiscalQuarterOneFiscalYearAgo SearchDate = "startOfLastFiscalQuarterOneFiscalYearAgo"
	SearchDateStartOfLastFiscalYear SearchDate = "startOfLastFiscalYear"
	SearchDateStartOfLastMonth SearchDate = "startOfLastMonth"
	SearchDateStartOfLastMonthOneFiscalQuarterAgo SearchDate = "startOfLastMonthOneFiscalQuarterAgo"
	SearchDateStartOfLastMonthOneFiscalYearAgo SearchDate = "startOfLastMonthOneFiscalYearAgo"
	SearchDateStartOfLastRollingHalf SearchDate = "startOfLastRollingHalf"
	SearchDateStartOfLastRollingQuarter SearchDate = "startOfLastRollingQuarter"
	SearchDateStartOfLastRollingYear SearchDate = "startOfLastRollingYear"
	SearchDateStartOfLastWeek SearchDate = "startOfLastWeek"
	SearchDateStartOfMonthBeforeLast SearchDate = "startOfMonthBeforeLast"
	SearchDateStartOfNextBusinessWeek SearchDate = "startOfNextBusinessWeek"
	SearchDateStartOfNextFiscalHalf SearchDate = "startOfNextFiscalHalf"
	SearchDateStartOfNextFiscalQuarter SearchDate = "startOfNextFiscalQuarter"
	SearchDateStartOfNextFiscalYear SearchDate = "startOfNextFiscalYear"
	SearchDateStartOfNextMonth SearchDate = "startOfNextMonth"
	SearchDateStartOfNextWeek SearchDate = "startOfNextWeek"
	SearchDateStartOfPreviousRollingHalf SearchDate = "startOfPreviousRollingHalf"
	SearchDateStartOfPreviousRollingQuarter SearchDate = "startOfPreviousRollingQuarter"
	SearchDateStartOfPreviousRollingYear SearchDate = "startOfPreviousRollingYear"
	SearchDateStartOfSameFiscalHalfLastFiscalYear SearchDate = "startOfSameFiscalHalfLastFiscalYear"
	SearchDateStartOfSameFiscalQuarterLastFiscalYear SearchDate = "startOfSameFiscalQuarterLastFiscalYear"
	SearchDateStartOfSameMonthLastFiscalQuarter SearchDate = "startOfSameMonthLastFiscalQuarter"
	SearchDateStartOfSameMonthLastFiscalYear SearchDate = "startOfSameMonthLastFiscalYear"
	SearchDateStartOfThisBusinessWeek SearchDate = "startOfThisBusinessWeek"
	SearchDateStartOfThisFiscalHalf SearchDate = "startOfThisFiscalHalf"
	SearchDateStartOfThisFiscalQuarter SearchDate = "startOfThisFiscalQuarter"
	SearchDateStartOfThisFiscalYear SearchDate = "startOfThisFiscalYear"
	SearchDateStartOfThisMonth SearchDate = "startOfThisMonth"
	SearchDateStartOfThisWeek SearchDate = "startOfThisWeek"
	SearchDateStartOfThisYear SearchDate = "startOfThisYear"
	SearchDateStartOfWeekBeforeLast SearchDate = "startOfWeekBeforeLast"
	SearchDateTenDaysAgo SearchDate = "tenDaysAgo"
	SearchDateTenDaysFromNow SearchDate = "tenDaysFromNow"
	SearchDateThirtyDaysAgo SearchDate = "thirtyDaysAgo"
	SearchDateThirtyDaysFromNow SearchDate = "thirtyDaysFromNow"
	SearchDateThisBusinessWeek SearchDate = "thisBusinessWeek"
	SearchDateThisFiscalHalf SearchDate = "thisFiscalHalf"
	SearchDateThisFiscalHalfToDate SearchDate = "thisFiscalHalfToDate"
	SearchDateThisFiscalQuarter SearchDate = "thisFiscalQuarter"
	SearchDateThisFiscalQuarterToDate SearchDate = "thisFiscalQuarterToDate"
	SearchDateThisFiscalYear SearchDate = "thisFiscalYear"
	SearchDateThisFiscalYearToDate SearchDate = "thisFiscalYearToDate"
	SearchDateThisMonth SearchDate = "thisMonth"
	SearchDateThisMonthToDate SearchDate = "thisMonthToDate"
	SearchDateThisRollingHalf SearchDate = "thisRollingHalf"
	SearchDateThisRollingQuarter SearchDate = "thisRollingQuarter"
	SearchDateThisRollingYear SearchDate = "thisRollingYear"
	SearchDateThisWeek SearchDate = "thisWeek"
	SearchDateThisWeekToDate SearchDate = "thisWeekToDate"
	SearchDateThisYear SearchDate = "thisYear"
	SearchDateThreeDaysAgo SearchDate = "threeDaysAgo"
	SearchDateThreeDaysFromNow SearchDate = "threeDaysFromNow"
	SearchDateThreeFiscalQuartersAgo SearchDate = "threeFiscalQuartersAgo"
	SearchDateThreeFiscalQuartersAgoToDate SearchDate = "threeFiscalQuartersAgoToDate"
	SearchDateThreeFiscalYearsAgo SearchDate = "threeFiscalYearsAgo"
	SearchDateThreeFiscalYearsAgoToDate SearchDate = "threeFiscalYearsAgoToDate"
	SearchDateThreeMonthsAgo SearchDate = "threeMonthsAgo"
	SearchDateThreeMonthsAgoToDate SearchDate = "threeMonthsAgoToDate"
	SearchDateToday SearchDate = "today"
	SearchDateTomorrow SearchDate = "tomorrow"
	SearchDateTwoDaysAgo SearchDate = "twoDaysAgo"
	SearchDateTwoDaysFromNow SearchDate = "twoDaysFromNow"
	SearchDateWeekAfterNext SearchDate = "weekAfterNext"
	SearchDateWeekAfterNextToDate SearchDate = "weekAfterNextToDate"
	SearchDateWeekBeforeLast SearchDate = "weekBeforeLast"
	SearchDateWeekBeforeLastToDate SearchDate = "weekBeforeLastToDate"
	SearchDateYesterday SearchDate = "yesterday"
)

type SearchRowList struct {
	XMLName xml.Name `xml:"searchRowList"`
	SearchRow []*SearchRow `xml:"searchRow,omitempty" json:"searchRow,omitempty"`
}
type SearchRow struct {
	XMLName xml.Name `xml:"searchRow"`
}

*/
