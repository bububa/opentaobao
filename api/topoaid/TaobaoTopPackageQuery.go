package topoaid

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/topoaid"
)

// TaobaoTopPackageQuery 淘系包裹查询
// taobao.top.package.query
//
// 淘系包裹查询
func TaobaoTopPackageQuery(clt *core.SDKClient, req *topoaid.TaobaoTopPackageQueryAPIRequest, session string) (*topoaid.TaobaoTopPackageQueryAPIResponse, error) {
	var resp topoaid.TaobaoTopPackageQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
