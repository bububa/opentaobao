package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkTradeOrderSuccessCreate 五道口终态订单创建
// alibaba.wdk.trade.order.success.create
//
// 五道口终态订单创建
func AlibabaWdkTradeOrderSuccessCreate(clt *core.SDKClient, req *wdk.AlibabaWdkTradeOrderSuccessCreateAPIRequest, session string) (*wdk.AlibabaWdkTradeOrderSuccessCreateAPIResponse, error) {
	var resp wdk.AlibabaWdkTradeOrderSuccessCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
