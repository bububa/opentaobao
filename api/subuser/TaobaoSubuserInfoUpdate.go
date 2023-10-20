package subuser

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/subuser"
)

// TaobaoSubuserInfoUpdate 修改指定账户子账号的基本信息
// taobao.subuser.info.update
//
// 修改指定账户子账号的基本信息（通过主账号登陆只能修改属于该主账号的子账号基本信息）
func TaobaoSubuserInfoUpdate(clt *core.SDKClient, req *subuser.TaobaoSubuserInfoUpdateAPIRequest, resp *subuser.TaobaoSubuserInfoUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
