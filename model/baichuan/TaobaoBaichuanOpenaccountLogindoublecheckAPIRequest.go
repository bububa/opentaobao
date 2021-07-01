package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanOpenaccountLogindoublecheckAPIRequest
百川登录二次验证 API请求
taobao.baichuan.openaccount.logindoublecheck

百川登录二次验证 */
type TaobaoBaichuanOpenaccountLogindoublecheckAPIRequest struct {
	model.Params
	// name
	_name string
}

// New
