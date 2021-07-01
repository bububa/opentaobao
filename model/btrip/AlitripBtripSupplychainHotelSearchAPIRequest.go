package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripSupplychainHotelSearchAPIRequest
【商旅】酒店订单查询 API请求
alitrip.btrip.supplychain.hotel.search

【商旅】酒店订单查询 */
type AlitripBtripSupplychainHotelSearchAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenApiSearchRq
}

// New
