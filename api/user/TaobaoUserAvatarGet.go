package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoUserAvatarGet 淘宝用户头像查询
// taobao.user.avatar.get
//
// 根据混淆nick查询用户头像
func TaobaoUserAvatarGet(clt *core.SDKClient, req *user.TaobaoUserAvatarGetAPIRequest, session string) (*user.TaobaoUserAvatarGetAPIResponse, error) {
	var resp user.TaobaoUserAvatarGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
