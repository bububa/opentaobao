package wlbimports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// TaobaoWlbImportsVasIdentityResult 集货鉴定结果
// taobao.wlb.imports.vas.identity.result
//
// 集货鉴定结果查询
func TaobaoWlbImportsVasIdentityResult(ctx context.Context, clt *core.SDKClient, req *wlbimports.TaobaoWlbImportsVasIdentityResultAPIRequest, resp *wlbimports.TaobaoWlbImportsVasIdentityResultAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
