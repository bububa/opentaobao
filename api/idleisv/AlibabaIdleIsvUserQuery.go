package idleisv

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleIsvUserQuery 服务商ISV闲鱼用户信息查询
// alibaba.idle.isv.user.query
//
// 服务商ISV闲鱼用户信息查询
func AlibabaIdleIsvUserQuery(ctx context.Context, clt *core.SDKClient, req *idleisv.AlibabaIdleIsvUserQueryAPIRequest, resp *idleisv.AlibabaIdleIsvUserQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
