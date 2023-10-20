package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// AlibabaWdkTradeOrderCancel 外部交易订单取消接口
// alibaba.wdk.trade.order.cancel
//
// 通过该接口可以再盒马取消交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
func AlibabaWdkTradeOrderCancel(clt *core.SDKClient, req *trade.AlibabaWdkTradeOrderCancelAPIRequest, resp *trade.AlibabaWdkTradeOrderCancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
