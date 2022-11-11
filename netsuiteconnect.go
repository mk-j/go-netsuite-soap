package netsuite

import (
	"fmt"
	"gopkg.in/yaml.v2"
//	"io"
	"io/ioutil"
//	"log"
//	"net/http"
//	"net/url"
//	"sort"
	"strings"
	"encoding/xml"
	"regexp"
	"time"
//	"reflect"
	"strconv"
	"github.com/hooklift/gowsdl/soap"
)

type NetSuiteConnector struct {
	Auth struct {
		Account        string `yaml:"account"`
		ConsumerKey    string `yaml:"consumerKey"`
		ConsumerSecret string `yaml:"consumerSecret"`
		Token          string `yaml:"token"`
		TokenSecret    string `yaml:"tokenSecret"`
		APIURL         string `yaml:"apiURL"`
		SoapURL        string `yaml:"soapURL"`
	} `yaml:"netsuite"`
	ConfigFile     string
}

type NetSuiteRequest struct {
	Connection *NetSuiteConnector
	RequestURL string
	QueryParams map[string]string
	RequestBody interface{} //can pass in json encoded string, or a struct or map here
}

type NSTokenPassportSignature struct {
	XMLName xml.Name `xml:"signature"`
	Value string `xml:",chardata" json:"-,"`
	Algorithm string `xml:"algorithm,attr,omitempty" json:"algorithm,omitempty"`
}

type NSTokenPassport struct {
	XMLName xml.Name `xml:"tokenPassport"`
	Account string `xml:"account,omitempty" json:"account,omitempty"`
	ConsumerKey string `xml:"consumerKey,omitempty" json:"consumerKey,omitempty"`
	Token string `xml:"token,omitempty" json:"token,omitempty"`
	Nonce string `xml:"nonce,omitempty" json:"nonce,omitempty"`
	Timestamp int64 `xml:"timestamp,omitempty" json:"timestamp,omitempty"`
	Signature *NSTokenPassportSignature `xml:"signature,omitempty" json:"signature,omitempty"`
}

// https://docs.oracle.com/en/cloud/saas/netsuite/ns-online-help/chapter_1540391670.html
/* REST web services always load record instances in "Edit" mode, including GET requests. 
 * When a user without the Administrator role tries to run a GET request for an employee 
 * record of a user with the Administrator role, this error message will appear: 
 * "For security reasons, only an administrator is allowed to edit an administrator record". 
 * This prevents users without the Administrator role from editing the employee record of a 
 * user with the Administrator role. */

//----------------------------------------------------------------------
func (r *NetSuiteConnector) ReadConfig(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(content, r)
	if err != nil {
		panic(err)
	}
	if r.Auth.Account == "" {
		panic("Missing from config.yml file: " + "auth: Account")
	}
	if r.Auth.APIURL == "" {
		panic("Missing from config.yml file: " + "auth: APIURL")
	}
	if r.Auth.SoapURL == "" {
		panic("Missing from config.yml file: " + "auth: SoapURL")
	}
	if r.Auth.Token == "" {
		panic("Missing from config.yml file: " + "auth: Token")
	}
	if r.Auth.TokenSecret == "" {
		panic("Missing from config.yml file: " + "auth: TokenSecret")
	}
	if r.Auth.ConsumerKey == "" {
		panic("Missing from config.yml file: " + "auth: ConsumerKey")
	}
	if r.Auth.ConsumerSecret == "" {
		panic("Missing from config.yml file: " + "auth: ConsumerSecret")
	}
}
//----------------------------------------------------------------------
// SOAP STUFF
//----------------------------------------------------------------------
func (r *NetSuiteConnector) SoapService() *NetSuiteService {    
    if len(r.Auth.SoapURL)==0 {
        panic("Missing from config.yml file: " + "auth: SoapURL")
    }
    wsdlRegex := regexp.MustCompile(`^https:\/\/([0-9\-sb]+)\.suitetalk\.api\.netsuite\.com\/wsdl\/v(2[0-9][0-9][0-9])_1_0\/netsuite\.wsdl`)
    pieces:= wsdlRegex.FindStringSubmatch(r.Auth.SoapURL)
    if len(pieces)!=3 {
        panic("Unable to parse soap netsuite.wsdl url")
    }
    subdomain := pieces[1]
    apiYear := pieces[2]
    soapAccessUrl := fmt.Sprintf("https://%s.suitetalk.api.netsuite.com/services/NetSuitePort_%s_1", subdomain, apiYear)

    client := soap.NewClient(soapAccessUrl)
    client.AddHeader(r.SoapAuthHeaderData())
    service := &NetSuiteService{
        UpstreamService:  NewNetSuitePortType(client),
    }
    return service
}
//----------------------------------------------------------------------
func (r *NetSuiteConnector) SoapAuthHeaderData() NSTokenPassport {
    nonce := RandString(20)
    timestamp := time.Now().Unix()
    timestr := strconv.FormatInt(timestamp,10)
    account:=r.Auth.Account

    consumer_key := r.Auth.ConsumerKey
    consumer_secret := r.Auth.ConsumerSecret
    token := r.Auth.Token
    token_secret := r.Auth.TokenSecret

    baseString:= strings.Join([]string{account, consumer_key, token, nonce, timestr}, "&")
    secretString:= strings.Join([]string{consumer_secret, token_secret}, "&")
    signature := HashHmacSha256(baseString, secretString)

    tokenPassportSignature:= NSTokenPassportSignature {
        Value: signature,
        Algorithm: "HMAC-SHA256",
    }
    tokenPassport:= NSTokenPassport {
        Account: account,
        ConsumerKey: consumer_key,
        Token: token,
        Nonce: nonce,
        Timestamp: timestamp,
        Signature: &tokenPassportSignature,
    }
    return tokenPassport
}
//----------------------------------------------------------------------
