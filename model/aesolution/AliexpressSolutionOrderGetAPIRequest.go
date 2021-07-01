package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionOrderGetAPIRequest
get order list API请求
aliexpress.solution.order.get

Get Order List from AliExpress */
type AliexpressSolutionOrderGetAPIRequest struct {
	model.Params
	// param
	_param0 *OrderQuery
}

// New
