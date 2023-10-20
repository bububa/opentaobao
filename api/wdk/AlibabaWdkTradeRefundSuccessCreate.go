package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkTradeRefundSuccessCreate 五道口终态逆向订单创建
// alibaba.wdk.trade.refund.success.create
//
// 五道口终态逆向订单创建
func AlibabaWdkTradeRefundSuccessCreate(clt *core.SDKClient, req *wdk.AlibabaWdkTradeRefundSuccessCreateAPIRequest, resp *wdk.AlibabaWdkTradeRefundSuccessCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
