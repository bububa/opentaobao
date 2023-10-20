package tbuser

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbuser"
)

// TaobaoUserAvatarGet 淘宝用户头像查询
// taobao.user.avatar.get
//
// 根据混淆nick查询用户头像
func TaobaoUserAvatarGet(clt *core.SDKClient, req *tbuser.TaobaoUserAvatarGetAPIRequest, resp *tbuser.TaobaoUserAvatarGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
