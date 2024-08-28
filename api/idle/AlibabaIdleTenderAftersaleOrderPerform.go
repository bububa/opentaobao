package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTenderAftersaleOrderPerform 闲鱼帮卖售后订单履约
// alibaba.idle.tender.aftersale.order.perform
//
// 闲鱼帮卖售后订单履约
func AlibabaIdleTenderAftersaleOrderPerform(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleTenderAftersaleOrderPerformAPIRequest, resp *idle.AlibabaIdleTenderAftersaleOrderPerformAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
