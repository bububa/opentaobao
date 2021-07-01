package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest
百川发送注册验证码 API请求
taobao.baichuan.openaccount.registercode.send

百川发送注册验证码 */
type TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest struct {
	model.Params
	// name
	_name string
}

// New
