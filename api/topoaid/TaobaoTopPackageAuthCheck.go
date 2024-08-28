package topoaid

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/topoaid"
)

// TaobaoTopPackageAuthCheck 校验用户授权关系
// taobao.top.package.auth.check
//
// 校验用户授权关系
func TaobaoTopPackageAuthCheck(ctx context.Context, clt *core.SDKClient, req *topoaid.TaobaoTopPackageAuthCheckAPIRequest, resp *topoaid.TaobaoTopPackageAuthCheckAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
