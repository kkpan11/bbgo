package coinbase

import "github.com/c9s/requestgen"

type CancelOrderResponse string

//go:generate requestgen -method DELETE -url "/orders/:order_id" -type CancelOrderRequest -responseType .CancelOrderResponse
type CancelOrderRequest struct {
	client requestgen.AuthenticatedAPIClient

	orderID string `param:"order_id,slug,required"`

	profileID *string `param:"profile_id"`
	productID *string `param:"product_id"`
}

func (c *RestAPIClient) NewCancelOrderRequest() *CancelOrderRequest {
	return &CancelOrderRequest{client: c}
}

type CancelAllOrdersResponse []string

//go:generate requestgen -method DELETE -url "/orders" -type CancelAllOrdersRequest -responseType .CancelAllOrdersResponse
type CancelAllOrdersRequest struct {
	client requestgen.AuthenticatedAPIClient

	profileID *string `param:"profile_id"`
	productID *string `param:"product_id"`
}

func (c *RestAPIClient) NewCancelAllOrdersRequest() *CancelAllOrdersRequest {
	return &CancelAllOrdersRequest{client: c}
}
