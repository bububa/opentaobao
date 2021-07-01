package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanOpenaccountLoginbytokenAPIRequest
百川TOKEN 登录 API请求
taobao.baichuan.openaccount.loginbytoken

百川TOKEN 登录 */
type TaobaoBaichuanOpenaccountLoginbytokenAPIRequest struct {
	model.Params
	// name
	_name string
}

// New
