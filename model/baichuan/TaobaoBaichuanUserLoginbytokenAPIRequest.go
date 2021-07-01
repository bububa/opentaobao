package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanUserLoginbytokenAPIRequest
百川手淘信任登录 API请求
taobao.baichuan.user.loginbytoken

百川手淘信任登录 */
type TaobaoBaichuanUserLoginbytokenAPIRequest struct {
	model.Params
	// name
	_name string
}

// New
