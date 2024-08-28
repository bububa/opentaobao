package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkTradeOrderSuccessCreate 五道口终态订单创建
// alibaba.wdk.trade.order.success.create
//
// 五道口终态订单创建
func AlibabaWdkTradeOrderSuccessCreate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkTradeOrderSuccessCreateAPIRequest, resp *wdk.AlibabaWdkTradeOrderSuccessCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
