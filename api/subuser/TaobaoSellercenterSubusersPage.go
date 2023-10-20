package subuser

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/subuser"
)

// TaobaoSellercenterSubusersPage 通过主账号登陆态分页查询子账号列表
// taobao.sellercenter.subusers.page
//
// 通过主账号登陆态分页查询子账号列表
func TaobaoSellercenterSubusersPage(clt *core.SDKClient, req *subuser.TaobaoSellercenterSubusersPageAPIRequest, resp *subuser.TaobaoSellercenterSubusersPageAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
