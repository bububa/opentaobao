package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkTradeOrderCancelAPIRequest
外部交易订单取消接口 API请求
alibaba.wdk.trade.order.cancel

通过该接口可以再盒马取消交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理 */
type AlibabaWdkTradeOrderCancelAPIRequest struct {
	model.Params
	// 待取消的订单
	_trade *TradeOrder
}

// New
