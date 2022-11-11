#!/bin/bash


sed -i 's/type ChangePassword ChangePasswordRequest/type ChangePasswordType ChangePasswordRequest/' netsuite/netsuite.go
sed -i 's/type ChangeEmail ChangeEmailRequest/type ChangeEmailType ChangeEmailRequest/' netsuite/netsuite.go
sed -i 's/ChangePassword *ChangePassword /ChangePassword *ChangePasswordType /' netsuite/netsuite.go
sed -i 's/ChangeEmail *ChangeEmail /ChangeEmail *ChangeEmailType /' netsuite/netsuite.go
sed -i 's/[^e]Record \*RecordRef/RelatedRecord \*RecordRef/' netsuite/netsuite.go
sed -i '0,/type CustomerSalesTeam struct/{s/type CustomerSalesTeam struct/type CustomerSalesTeamZZZ struct/}'  netsuite/netsuite.go 
sed -i '/^type CustomerSalesTeamZZZ struct /,/^}/ s/.*//'  netsuite/netsuite.go 

sed -i '0,/type CurrencyRate struct/{s/type CurrencyRate struct/type CurrencyRateItem struct/}'  netsuite/netsuite.go 
sed -i 's/CurrencyRate \[\]\*CurrencyRate/CurrencyRate []*CurrencyRateItem/' netsuite/netsuite.go

sed -i '0,/type SiteCategory struct/{s/type SiteCategory struct/type SiteCategoryItem struct/}'  netsuite/netsuite.go 
sed -i 's/SiteCategory \[\]\*SiteCategory/SiteCategory []*SiteCategoryItem/' netsuite/netsuite.go

sed -i 's/type CampaignResponse string/type CampaignResponseType string/' netsuite/netsuite.go
sed -i 's/CampaignResponse =/CampaignResponseType =/' netsuite/netsuite.go

cat << EOF >> netsuite/netsuite.go
type TimeBillTimeType string 
const (
    TimeBillTimeType_actualTime TimeBillTimeType = "_actualTime"
    TimeBillTimeType_allocatedTime TimeBillTimeType = "_allocatedTime"
    TimeBillTimeType_plannedTime TimeBillTimeType = "_plannedTime"
)
EOF

sed -i 's/urn:support_2022_1.lists.webservices.netsuite.com //'  netsuite/netsuite.go
sed -i 's/urn:website_2022_1.lists.webservices.netsuite.com //'  netsuite/netsuite.go
sed -i 's/urn:core_2022_1.platform.webservices.netsuite.com //'  netsuite/netsuite.go
sed -i 's/urn:employees_2022_1.lists.webservices.netsuite.com //'  netsuite/netsuite.go
sed -i 's/urn:faults_2022_1.platform.webservices.netsuite.com //'  netsuite/netsuite.go
sed -i 's/urn:marketing_2022_1.lists.webservices.netsuite.com //'  netsuite/netsuite.go
sed -i 's/urn:common_2022_1.platform.webservices.netsuite.com //'  netsuite/netsuite.go
sed -i 's/urn:accounting_2022_1.lists.webservices.netsuite.com //'  netsuite/netsuite.go
sed -i 's/urn:messages_2022_1.platform.webservices.netsuite.com //'  netsuite/netsuite.go
sed -i 's/urn:bank_2022_1.transactions.webservices.netsuite.com //'  netsuite/netsuite.go
sed -i 's/urn:supplychain_2022_1.lists.webservices.netsuite.com //'  netsuite/netsuite.go
sed -i 's/urn:sales_2022_1.transactions.webservices.netsuite.com //'  netsuite/netsuite.go
sed -i 's/urn:customization_2022_1.setup.webservices.netsuite.com //'  netsuite/netsuite.go
sed -i 's/urn:relationships_2022_1.lists.webservices.netsuite.com //'  netsuite/netsuite.go
sed -i 's/urn:general_2022_1.transactions.webservices.netsuite.com //'  netsuite/netsuite.go
sed -i 's/urn:scheduling_2022_1.activities.webservices.netsuite.com //'  netsuite/netsuite.go
sed -i 's/urn:filecabinet_2022_1.documents.webservices.netsuite.com //'  netsuite/netsuite.go
sed -i 's/urn:communication_2022_1.general.webservices.netsuite.com //'  netsuite/netsuite.go
sed -i 's/urn:financial_2022_1.transactions.webservices.netsuite.com //'  netsuite/netsuite.go
sed -i 's/urn:inventory_2022_1.transactions.webservices.netsuite.com //'  netsuite/netsuite.go
sed -i 's/urn:purchases_2022_1.transactions.webservices.netsuite.com //'  netsuite/netsuite.go
sed -i 's/urn:employees_2022_1.transactions.webservices.netsuite.com //'  netsuite/netsuite.go
sed -i 's/urn:customers_2022_1.transactions.webservices.netsuite.com //'  netsuite/netsuite.go
sed -i 's/urn:demandplanning_2022_1.transactions.webservices.netsuite.com //'  netsuite/netsuite.go
