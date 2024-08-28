package idleisv

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleUserPermitRevoke 删除服务商与卖家之间的订单消息绑定关系
// alibaba.idle.user.permit.revoke
//
// 删除服务商与卖家之间的订单消息绑定关系，删除后不再发送消息
func AlibabaIdleUserPermitRevoke(ctx context.Context, clt *core.SDKClient, req *idleisv.AlibabaIdleUserPermitRevokeAPIRequest, resp *idleisv.AlibabaIdleUserPermitRevokeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
