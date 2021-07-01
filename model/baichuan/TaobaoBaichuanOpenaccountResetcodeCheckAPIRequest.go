package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest
百川验证找回密码验证码 API请求
taobao.baichuan.openaccount.resetcode.check

百川验证找回密码验证码 */
type TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest struct {
	model.Params
	// name
	_name string
}

// New
