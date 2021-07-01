package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUserAvatarGetAPIRequest
淘宝用户头像查询 API请求
taobao.user.avatar.get

根据混淆nick查询用户头像 */
type TaobaoUserAvatarGetAPIRequest struct {
	model.Params
	// 混淆nick
	_nick string
}

// New
