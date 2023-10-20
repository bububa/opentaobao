package topoaid

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/topoaid"
)

// Taobaotoppackageunauthquery 查询某手机号下的包裹
// taobao.top.package.unauth.query
//
// 查询某手机号下的包裹
func Taobaotoppackageunauthquery(clt *core.SDKClient, req *topoaid.TaobaotoppackageunauthqueryAPIRequest, session string) (*topoaid.TaobaotoppackageunauthqueryAPIResponse, error) {
	var resp topoaid.TaobaotoppackageunauthqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
