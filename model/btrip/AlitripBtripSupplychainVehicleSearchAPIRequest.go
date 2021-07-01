package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripSupplychainVehicleSearchAPIRequest
【商旅】用车订单搜索 API请求
alitrip.btrip.supplychain.vehicle.search

【商旅】用车订单搜索 */
type AlitripBtripSupplychainVehicleSearchAPIRequest struct {
	model.Params
	// 出参
	_rq *OpenApiSearchRq
}

// New
