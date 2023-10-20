package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// AlibabaWdkTradeOrderCreate 外部交易订单创单接口
// alibaba.wdk.trade.order.create
//
// 通过该接口可以再盒马创建交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
func AlibabaWdkTradeOrderCreate(clt *core.SDKClient, req *trade.AlibabaWdkTradeOrderCreateAPIRequest, resp *trade.AlibabaWdkTradeOrderCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
