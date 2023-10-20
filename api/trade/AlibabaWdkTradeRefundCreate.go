package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// AlibabaWdkTradeRefundCreate 外部渠道逆向订单创建
// alibaba.wdk.trade.refund.create
//
// 该接口是创建退货订单的服务。当外部渠道发起退款后，调用此接口可以完成五道口底层交易、履约、配送等一系列流程进行退货。
func AlibabaWdkTradeRefundCreate(clt *core.SDKClient, req *trade.AlibabaWdkTradeRefundCreateAPIRequest, resp *trade.AlibabaWdkTradeRefundCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
