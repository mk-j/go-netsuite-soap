//GENERATED via go2wsdl, (and modified for naming conflicts, and functionality)

package netsuite

import (
    //"strings"
    "regexp"
    //"reflect"
    "encoding/xml"
    "time"
)

//----------------------------------------------------------------------
type CustomFieldList struct {
    CustomField []*CustomFieldRef `xml:"customField,omitempty" json:"customField,omitempty"`
}

type CustomFieldRef struct {
    XMLName xml.Name `xml:"customField" json:"-"`
    InternalId string `xml:"internalId,attr,omitempty" json:"internalId,omitempty"`
    FieldName string `xml:"scriptId,attr,omitempty" json:"scriptId,omitempty"`
    XsiType string `xml:"xsi:type,attr,omitempty" json:"type,omitempty"`
    FieldValue *CustomFieldData `xml:"value,omitempty" json:"value,omitempty"`
}
type CustomFieldData struct {
    XMLName xml.Name `xml:"value" json:"-"`
    FieldValue string `xml:",chardata" json:"fieldValue,omitempty"`
    InternalId string `xml:"internalId,attr,omitempty" json:"internalId,omitempty"`
}

//--------------------
func NSCustomDateField(name string, value string) *CustomFieldRef {
    dateval:=value
    if regexp.MustCompile(`^[0-9]{4}\-[0-9]{2}\-[0-9]{2}$`).MatchString(value) {
        parsed, _ := time.Parse("2006-01-02", value)
        dateval = parsed.Format("2006-01-02T15:04:05-0700")
    } else if regexp.MustCompile(`^[0-9]{4}\-[0-9]{2}\-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}$`).MatchString(value) {
        parsed, _ := time.Parse("2006-01-02 15:04:05", value)
        dateval = parsed.Format("2006-01-02T15:04:05-0700")
    }
    return &CustomFieldRef{
                    FieldName:name,
                    FieldValue:&CustomFieldData{FieldValue:dateval},
                    XsiType:"ns1:DateCustomFieldRef",
                }
}
func NSCustomStringField(name string, value string) *CustomFieldRef {
    return &CustomFieldRef{
                    FieldName:name,
                    FieldValue:&CustomFieldData{FieldValue:value},
                    XsiType:"ns1:StringCustomFieldRef",
                }
}
func NSCustomSelectField(name string, value string) *CustomFieldRef {
    return &CustomFieldRef{
                    FieldName:name,
                    FieldValue:&CustomFieldData{InternalId:value},
                    XsiType:"ns1:SelectCustomFieldRef",
                }
}
//--------------------
//          <customField scriptId="custbody_source_system" xsi:type="SelectCustomFieldRef">
//            <value internalId="1"/>
//          </customField>

//XsiType: one of: LongCustomFieldRef,DoubleCustomFieldRef,BooleanCustomFieldRef,StringCustomFieldRef,DateCustomFieldRef,SelectCustomFieldRef
//DateCustomFieldRef:2022-10-20T00:00:00-06:00

/*
type CustomFieldRef struct {
    XMLName xml.Name `xml:"customField"`
    InternalId string `xml:"internalId,attr,omitempty" json:"internalId,omitempty"`
    ScriptId string `xml:"scriptId,attr,omitempty" json:"scriptId,omitempty"`
}

type LongCustomFieldRef struct {
    *CustomFieldRef
    Value int64 `xml:"value,omitempty" json:"value,omitempty"`
}

type DoubleCustomFieldRef struct {
    *CustomFieldRef
    Value float64 `xml:"value,omitempty" json:"value,omitempty"`
}

type BooleanCustomFieldRef struct {
    *CustomFieldRef
    Value bool `xml:"value,omitempty" json:"value,omitempty"`
}

type StringCustomFieldRef struct {
    *CustomFieldRef
    Value string `xml:"value,omitempty" json:"value,omitempty"`
}

type DateCustomFieldRef struct {
    *CustomFieldRef
    Value soap.XSDDateTime `xml:"value,omitempty" json:"value,omitempty"`
}

type SelectCustomFieldRef struct {
    XMLName xml.Name `xml:"matrixOption"`
    *CustomFieldRef
    Value *ListOrRecordRef `xml:"value,omitempty" json:"value,omitempty"`
}

type MultiSelectCustomFieldRef struct {
    *CustomFieldRef
    Value []*ListOrRecordRef `xml:"value,omitempty" json:"value,omitempty"`
}*/

