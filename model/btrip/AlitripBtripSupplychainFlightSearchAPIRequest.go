package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripSupplychainFlightSearchAPIRequest
【商旅】机票订单查询 API请求
alitrip.btrip.supplychain.flight.search

【商旅】机票订单查询 */
type AlitripBtripSupplychainFlightSearchAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenApiSearchRq
}

// New
