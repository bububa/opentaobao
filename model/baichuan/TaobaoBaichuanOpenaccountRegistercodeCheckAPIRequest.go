package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest
百川检查注册验证码 API请求
taobao.baichuan.openaccount.registercode.check

百川检查注册验证码 */
type TaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest struct {
	model.Params
	// name
	_name string
}

// New
