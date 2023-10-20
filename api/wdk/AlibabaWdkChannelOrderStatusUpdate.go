package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkChannelOrderStatusUpdate 订单状态变更
// alibaba.wdk.channel.order.status.update
//
// 订单状态变更
func AlibabaWdkChannelOrderStatusUpdate(clt *core.SDKClient, req *wdk.AlibabaWdkChannelOrderStatusUpdateAPIRequest, resp *wdk.AlibabaWdkChannelOrderStatusUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
