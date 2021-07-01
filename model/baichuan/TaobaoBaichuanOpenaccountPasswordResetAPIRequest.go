package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanOpenaccountPasswordResetAPIRequest
百川找回密码 API请求
taobao.baichuan.openaccount.password.reset

百川找回密码 */
type TaobaoBaichuanOpenaccountPasswordResetAPIRequest struct {
	model.Params
	// name
	_name string
}

// New
