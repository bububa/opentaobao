package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripOpenSupplychainHotelTradeAPIRequest
【商旅】酒店交易查询流水接口 API请求
alitrip.btrip.open.supplychain.hotel.trade

【商旅】酒店交易查询流水接口——杭州市政府 */
type AlitripBtripOpenSupplychainHotelTradeAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenApiZzdSearchRq
}

// New
