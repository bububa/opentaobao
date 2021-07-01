package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanOpenaccountLoginAPIRequest
百川用户名密码登录 API请求
taobao.baichuan.openaccount.login

百川用户名密码登录 */
type TaobaoBaichuanOpenaccountLoginAPIRequest struct {
	model.Params
	// name
	_name string
}

// New
