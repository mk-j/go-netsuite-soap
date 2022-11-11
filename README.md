# go-netsuite
Go/Golang library for connecting to soap/suitetalk API

This requires a config/config.yml file in the following format:
```
netsuite:
    account: 1234567_SB1
    token: 0000111122223333444455556666777788889999aaaabbbbccccddddeeeeffff
    tokenSecret: 0000111122223333444455556666777788889999aaaabbbbccccddddeeeeffff
    consumerKey: 0000111122223333444455556666777788889999aaaabbbbccccddddeeeeffff
    consumerSecret: 0000111122223333444455556666777788889999aaaabbbbccccddddeeeeffff
    soapURL: https://1234567-sb1.suitetalk.api.netsuite.com/wsdl/v2022_1_0/netsuite.wsdl
```

```go
package main

import (
    "fmt"
    "github.com/mk-j/go-netsuite-soap"
)

func main() {
	conn := netsuite.NetSuiteConnector{}
	conn.ReadConfig("./config/config.yml")

	//SOAP API
	soapcustomer, err := conn.SoapService().GetCustomer("5220784")
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(JsonEncode(soapcustomer))
}

```

