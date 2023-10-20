package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// TaobaoWlbImportsResourceGet 获取所有服务列表
// taobao.wlb.imports.resource.get
//
// 一般进口TOP接口，获取所有服务列表。
func TaobaoWlbImportsResourceGet(clt *core.SDKClient, req *wlbimports.TaobaoWlbImportsResourceGetAPIRequest, resp *wlbimports.TaobaoWlbImportsResourceGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
