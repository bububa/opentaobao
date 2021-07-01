package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionOrderInfoGetAPIRequest
get order detail info API请求
aliexpress.solution.order.info.get

get order detail info */
type AliexpressSolutionOrderInfoGetAPIRequest struct {
	model.Params
	// param
	_param1 *OrderDetailQuery
}

// New
