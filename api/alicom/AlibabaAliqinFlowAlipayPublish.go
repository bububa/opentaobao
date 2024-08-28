package alicom

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinFlowAlipayPublish 流量钱包流量发放-面向支付宝用户
// alibaba.aliqin.flow.alipay.publish
//
// 用户淘宝流量钱包商家给支付宝用户发放流量
func AlibabaAliqinFlowAlipayPublish(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaAliqinFlowAlipayPublishAPIRequest, resp *alicom.AlibabaAliqinFlowAlipayPublishAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
