package nrt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtAssetAuthorizationAdd 增加数据权限授权
// tmall.nrt.asset.authorization.add
//
// 增加数据权限授权
func TmallNrtAssetAuthorizationAdd(ctx context.Context, clt *core.SDKClient, req *nrt.TmallNrtAssetAuthorizationAddAPIRequest, resp *nrt.TmallNrtAssetAuthorizationAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
