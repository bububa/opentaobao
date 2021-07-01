package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkTradeOrderCreateAPIRequest
外部交易订单创单接口 API请求
alibaba.wdk.trade.order.create

通过该接口可以再盒马创建交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理 */
type AlibabaWdkTradeOrderCreateAPIRequest struct {
	model.Params
	// 待创建的订单
	_trade *TradeOrder
}

// New
