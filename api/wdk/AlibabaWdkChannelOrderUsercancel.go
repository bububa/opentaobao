package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkChannelOrderUsercancel 用户发起售中取消
// alibaba.wdk.channel.order.usercancel
//
// 用户发起售中取消
func AlibabaWdkChannelOrderUsercancel(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkChannelOrderUsercancelAPIRequest, resp *wdk.AlibabaWdkChannelOrderUsercancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
