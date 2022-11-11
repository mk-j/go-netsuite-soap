#!/bin/bash
echo "step 2: use gowsdl to build go code"
gowsdl -p netsuite -o netsuite.go https://webservices.netsuite.com/wsdl/v2022_1_0/netsuite.wsdl

