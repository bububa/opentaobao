package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaLoginAuthsignGetAPIRequest
获取登陆权限签名 API请求
taobao.simba.login.authsign.get

获取登陆权限签名 */
type TaobaoSimbaLoginAuthsignGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
}

// New
