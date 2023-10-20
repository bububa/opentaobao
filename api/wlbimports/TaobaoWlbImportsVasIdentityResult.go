package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// TaobaoWlbImportsVasIdentityResult 集货鉴定结果
// taobao.wlb.imports.vas.identity.result
//
// 集货鉴定结果查询
func TaobaoWlbImportsVasIdentityResult(clt *core.SDKClient, req *wlbimports.TaobaoWlbImportsVasIdentityResultAPIRequest, resp *wlbimports.TaobaoWlbImportsVasIdentityResultAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
