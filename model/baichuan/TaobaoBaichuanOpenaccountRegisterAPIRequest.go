package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanOpenaccountRegisterAPIRequest
百川账号注册 API请求
taobao.baichuan.openaccount.register

百川账号注册 */
type TaobaoBaichuanOpenaccountRegisterAPIRequest struct {
	model.Params
	// name
	_name string
}

// New
