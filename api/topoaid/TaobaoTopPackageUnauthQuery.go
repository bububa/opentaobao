package topoaid

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/topoaid"
)

// TaobaoTopPackageUnauthQuery 查询某手机号下的包裹
// taobao.top.package.unauth.query
//
// 查询某手机号下的包裹
func TaobaoTopPackageUnauthQuery(clt *core.SDKClient, req *topoaid.TaobaoTopPackageUnauthQueryAPIRequest, resp *topoaid.TaobaoTopPackageUnauthQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
