package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripRailTradeRefundAPIRequest
退票接口 API请求
alitrip.rail.trade.refund

退票接口 */
type AlitripRailTradeRefundAPIRequest struct {
	model.Params
	// 入参
	_refundParam *RefundReq
}

// New
