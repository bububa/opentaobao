package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleIsvGoosefishOrderCreate 闲鱼三方安康容器订单创建
// alibaba.idle.isv.goosefish.order.create
//
// 闲鱼三方安康容器订单创建
func AlibabaIdleIsvGoosefishOrderCreate(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleIsvGoosefishOrderCreateAPIRequest, resp *idle.AlibabaIdleIsvGoosefishOrderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
