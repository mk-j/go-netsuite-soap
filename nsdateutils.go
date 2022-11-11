package netsuite

import (
    "github.com/hooklift/gowsdl/soap"
    "time"
)

//--------------------
func XDateFromISOString(input_date string) soap.XSDDate {
    myDate, _ := time.Parse("2006-01-02 15:04:05", input_date)
    return soap.CreateXsdDate(myDate, false)
}
//--------------------
func XDateTimeFromISOString(input_date string) soap.XSDDateTime {
    myDate, _ := time.Parse("2006-01-02 15:04:05", input_date)
    return soap.CreateXsdDateTime(myDate, false)
}
//--------------------
