package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoUserOpenidGet 查询用户openId
// taobao.user.openid.get
//
// 查询用户openId
func TaobaoUserOpenidGet(clt *core.SDKClient, req *user.TaobaoUserOpenidGetAPIRequest, resp *user.TaobaoUserOpenidGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
