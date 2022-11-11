package netsuite

import (
	"context"
	//"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
	//"time"
)

type NetSuitePortType interface {
	// Errors below can be either of the following types:
	//
	//   - InsufficientPermissionFault
	//   - InvalidAccountFault
	//   - InvalidCredentialsFault
	//   - InvalidSessionFault
	//   - InvalidVersionFault
	//   - ExceededRequestLimitFault
	//   - UnexpectedErrorFault


	Add(request *AddRequest) (*AddResponse, error)
	AddContext(ctx context.Context, request *AddRequest) (*AddResponse, error)

	Get(request *GetRequest) (*GetResponse, error)
	GetContext(ctx context.Context, request *GetRequest) (*GetResponse, error)

	Delete(request *DeleteRequest) (*DeleteResponse, error)
	DeleteContext(ctx context.Context, request *DeleteRequest) (*DeleteResponse, error)

	Terms_Search(request *SearchRequest) (*SearchResponse, error)
	Terms_SearchContext(ctx context.Context, request *SearchRequest) (*SearchResponse, error)

	Category_Get(request *Category_GetRequest) (*Category_GetResponse, error)
	Category_GetContext(ctx context.Context, request *Category_GetRequest) (*Category_GetResponse, error)

	//---
	Customer_Add(request *Customer_AddRequest) (*Customer_AddResponse, error)
	Customer_AddContext(ctx context.Context, request *Customer_AddRequest) (*Customer_AddResponse, error)

	Customer_Get(request *Customer_GetRequest) (*Customer_GetResponse, error)
	Customer_GetContext(ctx context.Context, request *Customer_GetRequest) (*Customer_GetResponse, error)

	//---
	CreditMemo_Add(request *CreditMemo_AddRequest) (*CreditMemo_AddResponse, error)
	CreditMemo_AddContext(ctx context.Context, request *CreditMemo_AddRequest) (*CreditMemo_AddResponse, error)

	CreditMemo_Get(request *CreditMemo_GetRequest) (*CreditMemo_GetResponse, error)
	CreditMemo_GetContext(ctx context.Context, request *CreditMemo_GetRequest) (*CreditMemo_GetResponse, error)

	//---
	CashSale_Add(request *CashSale_AddRequest) (*CashSale_AddResponse, error)
	CashSale_AddContext(ctx context.Context, request *CashSale_AddRequest) (*CashSale_AddResponse, error)

	CashSale_Get(request *CashSale_GetRequest) (*CashSale_GetResponse, error)
	CashSale_GetContext(ctx context.Context, request *CashSale_GetRequest) (*CashSale_GetResponse, error)

	//---
	Invoice_Get(request *Invoice_GetRequest) (*Invoice_GetResponse, error)
	Invoice_GetContext(ctx context.Context, request *Invoice_GetRequest) (*Invoice_GetResponse, error)

	Invoice_Add(request *Invoice_AddRequest) (*Invoice_AddResponse, error)
	Invoice_AddContext(ctx context.Context, request *Invoice_AddRequest) (*Invoice_AddResponse, error)

	//---
	SalesOrder_Get(request *SalesOrder_GetRequest) (*SalesOrder_GetResponse, error)
	SalesOrder_GetContext(ctx context.Context, request *SalesOrder_GetRequest) (*SalesOrder_GetResponse, error)

	SalesOrder_Add(request *SalesOrder_AddRequest) (*SalesOrder_AddResponse, error)
	SalesOrder_AddContext(ctx context.Context, request *SalesOrder_AddRequest) (*SalesOrder_AddResponse, error)

	//---
	Currency_GetAll(request *Currency_GetAllRequest) (*Currency_GetAllResponse, error)
	Currency_GetAllContext(ctx context.Context, request *Currency_GetAllRequest) (*Currency_GetAllResponse, error)

}

type netSuitePortType struct {
	client *soap.Client
}

func NewNetSuitePortType(client *soap.Client) NetSuitePortType {
	return &netSuitePortType{
		client: client,
	}
}
//----------------------------------------------------------------------
func (service *netSuitePortType) AddContext(ctx context.Context, request *AddRequest) (*AddResponse, error) {
	response := new(AddResponse)
	err := service.client.CallContext(ctx, "add", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *netSuitePortType) Add(request *AddRequest) (*AddResponse, error) {
	return service.AddContext(
		context.Background(),
		request,
	)
}

//----------------------------------------------------------------------

func (service *netSuitePortType) DeleteContext(ctx context.Context, request *DeleteRequest) (*DeleteResponse, error) {
	response := new(DeleteResponse)
	err := service.client.CallContext(ctx, "delete", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *netSuitePortType) Delete(request *DeleteRequest) (*DeleteResponse, error) {
	return service.DeleteContext(
		context.Background(),
		request,
	)
}
//----------------------------------------------------------------------

func (service *netSuitePortType) GetContext(ctx context.Context, request *GetRequest) (*GetResponse, error) {
	response := new(GetResponse)
	err := service.client.CallContext(ctx, "get", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *netSuitePortType) Get(request *GetRequest) (*GetResponse, error) {
	return service.GetContext(
		context.Background(),
		request,
	)
}

//----------------------------------------------------------------------

func (service *netSuitePortType) Category_GetContext(ctx context.Context, request *Category_GetRequest) (*Category_GetResponse, error) {
	response := new(Category_GetResponse)
	err := service.client.CallContext(ctx, "get", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *netSuitePortType) Category_Get(request *Category_GetRequest) (*Category_GetResponse, error) {
	return service.Category_GetContext(
		context.Background(),
		request,
	)
}


//----------------------------------------------------------------------

func (service *netSuitePortType) Customer_AddContext(ctx context.Context, request *Customer_AddRequest) (*Customer_AddResponse, error) {
	response := new(Customer_AddResponse)
	err := service.client.CallContext(ctx, "add", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *netSuitePortType) Customer_Add(request *Customer_AddRequest) (*Customer_AddResponse, error) {
	return service.Customer_AddContext(
		context.Background(),
		request,
	)
}

func (service *netSuitePortType) Customer_GetContext(ctx context.Context, request *Customer_GetRequest) (*Customer_GetResponse, error) {
	response := new(Customer_GetResponse)
	err := service.client.CallContext(ctx, "get", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *netSuitePortType) Customer_Get(request *Customer_GetRequest) (*Customer_GetResponse, error) {
	return service.Customer_GetContext(
		context.Background(),
		request,
	)
}

//----------------------------------------------------------------------

func (service *netSuitePortType) CashSale_AddContext(ctx context.Context, request *CashSale_AddRequest) (*CashSale_AddResponse, error) {
	response := new(CashSale_AddResponse)
	err := service.client.CallContext(ctx, "add", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *netSuitePortType) CashSale_Add(request *CashSale_AddRequest) (*CashSale_AddResponse, error) {
	return service.CashSale_AddContext(
		context.Background(),
		request,
	)
}

func (service *netSuitePortType) CashSale_GetContext(ctx context.Context, request *CashSale_GetRequest) (*CashSale_GetResponse, error) {
	response := new(CashSale_GetResponse)
	err := service.client.CallContext(ctx, "get", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *netSuitePortType) CashSale_Get(request *CashSale_GetRequest) (*CashSale_GetResponse, error) {
	return service.CashSale_GetContext(
		context.Background(),
		request,
	)
}
//----------------------------------------------------------------------

func (service *netSuitePortType) CreditMemo_AddContext(ctx context.Context, request *CreditMemo_AddRequest) (*CreditMemo_AddResponse, error) {
	response := new(CreditMemo_AddResponse)
	err := service.client.CallContext(ctx, "add", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *netSuitePortType) CreditMemo_Add(request *CreditMemo_AddRequest) (*CreditMemo_AddResponse, error) {
	return service.CreditMemo_AddContext(
		context.Background(),
		request,
	)
}

func (service *netSuitePortType) CreditMemo_GetContext(ctx context.Context, request *CreditMemo_GetRequest) (*CreditMemo_GetResponse, error) {
	response := new(CreditMemo_GetResponse)
	err := service.client.CallContext(ctx, "get", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *netSuitePortType) CreditMemo_Get(request *CreditMemo_GetRequest) (*CreditMemo_GetResponse, error) {
	return service.CreditMemo_GetContext(
		context.Background(),
		request,
	)
}

//----------------------------------------------------------------------

func (service *netSuitePortType) Invoice_AddContext(ctx context.Context, request *Invoice_AddRequest) (*Invoice_AddResponse, error) {
	response := new(Invoice_AddResponse)
	err := service.client.CallContext(ctx, "add", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *netSuitePortType) Invoice_Add(request *Invoice_AddRequest) (*Invoice_AddResponse, error) {
	return service.Invoice_AddContext(
		context.Background(),
		request,
	)
}

func (service *netSuitePortType) Invoice_GetContext(ctx context.Context, request *Invoice_GetRequest) (*Invoice_GetResponse, error) {
	response := new(Invoice_GetResponse)
	err := service.client.CallContext(ctx, "get", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *netSuitePortType) Invoice_Get(request *Invoice_GetRequest) (*Invoice_GetResponse, error) {
	return service.Invoice_GetContext(
		context.Background(),
		request,
	)
}

//----------------------------------------------------------------------

func (service *netSuitePortType) SalesOrder_AddContext(ctx context.Context, request *SalesOrder_AddRequest) (*SalesOrder_AddResponse, error) {
	response := new(SalesOrder_AddResponse)
	err := service.client.CallContext(ctx, "add", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *netSuitePortType) SalesOrder_Add(request *SalesOrder_AddRequest) (*SalesOrder_AddResponse, error) {
	return service.SalesOrder_AddContext(
		context.Background(),
		request,
	)
}

func (service *netSuitePortType) SalesOrder_GetContext(ctx context.Context, request *SalesOrder_GetRequest) (*SalesOrder_GetResponse, error) {
	response := new(SalesOrder_GetResponse)
	err := service.client.CallContext(ctx, "get", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *netSuitePortType) SalesOrder_Get(request *SalesOrder_GetRequest) (*SalesOrder_GetResponse, error) {
	return service.SalesOrder_GetContext(
		context.Background(),
		request,
	)
}

//----------------------------------------------------------------------

func (service *netSuitePortType) Currency_GetAllContext(ctx context.Context, request *Currency_GetAllRequest) (*Currency_GetAllResponse, error) {
	response := new(Currency_GetAllResponse)
	err := service.client.CallContext(ctx, "getAll", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *netSuitePortType) Currency_GetAll(request *Currency_GetAllRequest) (*Currency_GetAllResponse, error) {
	return service.Currency_GetAllContext(
		context.Background(),
		request,
	)
}

//----------------------------------------------------------------------

func (service *netSuitePortType) Terms_SearchContext(ctx context.Context, request *SearchRequest) (*SearchResponse, error) {
	response := new(SearchResponse)
	err := service.client.CallContext(ctx, "search", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *netSuitePortType) Terms_Search(request *SearchRequest) (*SearchResponse, error) {	
	return service.Terms_SearchContext(
		context.Background(),
		request,
	)
}

