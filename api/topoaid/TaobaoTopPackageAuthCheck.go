package topoaid

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/topoaid"
)

// Taobaotoppackageauthcheck 校验用户授权关系
// taobao.top.package.auth.check
//
// 校验用户授权关系
func Taobaotoppackageauthcheck(clt *core.SDKClient, req *topoaid.TaobaotoppackageauthcheckAPIRequest, session string) (*topoaid.TaobaotoppackageauthcheckAPIResponse, error) {
	var resp topoaid.TaobaotoppackageauthcheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
