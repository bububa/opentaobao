package topoaid

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/topoaid"
)

// TaobaoTopPackageQuery 淘系包裹查询
// taobao.top.package.query
//
// 淘系包裹查询
func TaobaoTopPackageQuery(ctx context.Context, clt *core.SDKClient, req *topoaid.TaobaoTopPackageQueryAPIRequest, resp *topoaid.TaobaoTopPackageQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
