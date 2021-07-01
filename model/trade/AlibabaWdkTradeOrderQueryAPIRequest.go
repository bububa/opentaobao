package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkTradeOrderQueryAPIRequest
查询外部交易订单接口 API请求
alibaba.wdk.trade.order.query

通过该接口可以在盒马查询交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理 */
type AlibabaWdkTradeOrderQueryAPIRequest struct {
	model.Params
	// 订单查询
	_query *TradeOrderQuery
}

// New
