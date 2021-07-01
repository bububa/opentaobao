package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripOpenSupplychainTrainTradeAPIRequest
商旅火车票交易流水接口 API请求
alitrip.btrip.open.supplychain.train.trade

商旅火车票交易流水接口 */
type AlitripBtripOpenSupplychainTrainTradeAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenApiZzdSearchRq
}

// New
