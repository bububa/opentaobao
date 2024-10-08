package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkChannelOrderUserrefund 用户发起售后退款(整单/部分)
// alibaba.wdk.channel.order.userrefund
//
// 用户发起售后退款(整单/部分)
func AlibabaWdkChannelOrderUserrefund(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkChannelOrderUserrefundAPIRequest, resp *wdk.AlibabaWdkChannelOrderUserrefundAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
