package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AlibabaEleFengniaoShippingorderEvent 查询运单事件信息
// alibaba.ele.fengniao.shippingorder.event
//
// 查询运单事件信息
func AlibabaEleFengniaoShippingorderEvent(ctx context.Context, clt *core.SDKClient, req *logistic.AlibabaEleFengniaoShippingorderEventAPIRequest, resp *logistic.AlibabaEleFengniaoShippingorderEventAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
