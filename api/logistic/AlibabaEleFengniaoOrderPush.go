package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AlibabaEleFengniaoOrderPush 推送订单
// alibaba.ele.fengniao.order.push
//
// 推送淘宝订单至蜂鸟开放平台配送
func AlibabaEleFengniaoOrderPush(ctx context.Context, clt *core.SDKClient, req *logistic.AlibabaEleFengniaoOrderPushAPIRequest, resp *logistic.AlibabaEleFengniaoOrderPushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
