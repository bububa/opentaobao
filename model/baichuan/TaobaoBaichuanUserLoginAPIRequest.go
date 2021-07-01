package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanUserLoginAPIRequest
百川H5登录 API请求
taobao.baichuan.user.login

百川H5登录 */
type TaobaoBaichuanUserLoginAPIRequest struct {
	model.Params
	// name
	_name string
}

// New
