package topoaid

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/topoaid"
)

// TaobaoTopPackageUnauthQuery 查询某手机号下的包裹
// taobao.top.package.unauth.query
//
// 查询某手机号下的包裹
func TaobaoTopPackageUnauthQuery(clt *core.SDKClient, req *topoaid.TaobaoTopPackageUnauthQueryAPIRequest, session string) (*topoaid.TaobaoTopPackageUnauthQueryAPIResponse, error) {
	var resp topoaid.TaobaoTopPackageUnauthQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
