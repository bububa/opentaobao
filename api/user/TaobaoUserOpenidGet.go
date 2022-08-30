package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoUserOpenidGet 查询用户openId
// taobao.user.openid.get
//
// 查询用户openId
func TaobaoUserOpenidGet(clt *core.SDKClient, req *user.TaobaoUserOpenidGetAPIRequest, session string) (*user.TaobaoUserOpenidGetAPIResponse, error) {
	var resp user.TaobaoUserOpenidGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
