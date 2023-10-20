package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkTradeRefundSuccessCreate 五道口终态逆向订单创建
// alibaba.wdk.trade.refund.success.create
//
// 五道口终态逆向订单创建
func AlibabaWdkTradeRefundSuccessCreate(clt *core.SDKClient, req *wdk.AlibabaWdkTradeRefundSuccessCreateAPIRequest, session string) (*wdk.AlibabaWdkTradeRefundSuccessCreateAPIResponse, error) {
	var resp wdk.AlibabaWdkTradeRefundSuccessCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
