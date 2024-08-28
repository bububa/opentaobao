package idleisv

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleUserPermitQuery 查询服务商与卖家之间的订单消息绑定关系
// alibaba.idle.user.permit.query
//
// 查询服务商与卖家之间的订单消息绑定关系
func AlibabaIdleUserPermitQuery(ctx context.Context, clt *core.SDKClient, req *idleisv.AlibabaIdleUserPermitQueryAPIRequest, resp *idleisv.AlibabaIdleUserPermitQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
