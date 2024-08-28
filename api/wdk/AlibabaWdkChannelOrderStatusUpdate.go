package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkChannelOrderStatusUpdate 订单状态变更
// alibaba.wdk.channel.order.status.update
//
// 订单状态变更
func AlibabaWdkChannelOrderStatusUpdate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkChannelOrderStatusUpdateAPIRequest, resp *wdk.AlibabaWdkChannelOrderStatusUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
