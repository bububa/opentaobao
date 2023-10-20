package topoaid

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/topoaid"
)

// Taobaotoppackagequery 淘系包裹查询
// taobao.top.package.query
//
// 淘系包裹查询
func Taobaotoppackagequery(clt *core.SDKClient, req *topoaid.TaobaotoppackagequeryAPIRequest, session string) (*topoaid.TaobaotoppackagequeryAPIResponse, error) {
	var resp topoaid.TaobaotoppackagequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
