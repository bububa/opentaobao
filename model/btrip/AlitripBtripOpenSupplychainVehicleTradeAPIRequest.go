package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripOpenSupplychainVehicleTradeAPIRequest
商旅用车交易流水接口 API请求
alitrip.btrip.open.supplychain.vehicle.trade

商旅用车交易流水接口 */
type AlitripBtripOpenSupplychainVehicleTradeAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenApiZzdSearchRq
}

// New
