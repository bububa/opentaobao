package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTopOnceTokenGetAPIRequest
网关一次性token获取 API请求
taobao.top.once.token.get

网关一次性token获取 */
type TaobaoTopOnceTokenGetAPIRequest struct {
	model.Params
	// sec_token
	_secToken string
}

// New
