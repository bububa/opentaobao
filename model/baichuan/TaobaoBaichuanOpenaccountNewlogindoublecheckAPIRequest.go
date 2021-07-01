package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanOpenaccountNewlogindoublecheckAPIRequest
百川新登录二次验证 API请求
taobao.baichuan.openaccount.newlogindoublecheck

百川新登录二次验证 */
type TaobaoBaichuanOpenaccountNewlogindoublecheckAPIRequest struct {
	model.Params
	// name
	_name string
}

// New
