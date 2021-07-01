package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanUserLogindoublecheckAPIRequest
百川H5登录二次验证 API请求
taobao.baichuan.user.logindoublecheck

百川H5登录二次验证 */
type TaobaoBaichuanUserLogindoublecheckAPIRequest struct {
	model.Params
	// name
	_name string
}

// New
