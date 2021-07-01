package nazca

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNazcaTokenAuthapplyGetAPIRequest
根据token获取认证申请信息 API请求
alibaba.nazca.token.authapply.get

根据token获取认证申请信息 */
type AlibabaNazcaTokenAuthapplyGetAPIRequest struct {
	model.Params
	// token
	_token string
}

// New
