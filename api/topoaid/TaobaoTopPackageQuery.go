package topoaid

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/topoaid"
)

// TaobaoTopPackageQuery 淘系包裹查询
// taobao.top.package.query
//
// 淘系包裹查询
func TaobaoTopPackageQuery(clt *core.SDKClient, req *topoaid.TaobaoTopPackageQueryAPIRequest, resp *topoaid.TaobaoTopPackageQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
