package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripOpenSupplychainFlightTradeAPIRequest
【商旅】机票交易流水查询接口 API请求
alitrip.btrip.open.supplychain.flight.trade

【商旅】杭州市政府机票交易流水接口查询 */
type AlitripBtripOpenSupplychainFlightTradeAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenApiZzdSearchRq
}

// New
