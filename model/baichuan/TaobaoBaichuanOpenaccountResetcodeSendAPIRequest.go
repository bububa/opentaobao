package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanOpenaccountResetcodeSendAPIRequest
百川发送找回密码验证码 API请求
taobao.baichuan.openaccount.resetcode.send

百川发送找回密码验证码 */
type TaobaoBaichuanOpenaccountResetcodeSendAPIRequest struct {
	model.Params
	// name
	_name string
}

// New
