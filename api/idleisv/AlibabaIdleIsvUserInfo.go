package idleisv

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleIsvUserInfo 闲鱼用户信息查询接口
// alibaba.idle.isv.user.info
//
// 闲鱼用户信息查询接口
func AlibabaIdleIsvUserInfo(ctx context.Context, clt *core.SDKClient, req *idleisv.AlibabaIdleIsvUserInfoAPIRequest, resp *idleisv.AlibabaIdleIsvUserInfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
