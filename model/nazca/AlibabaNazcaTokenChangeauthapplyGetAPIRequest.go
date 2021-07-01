package nazca

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNazcaTokenChangeauthapplyGetAPIRequest
根据token获取变更认证申请信息 API请求
alibaba.nazca.token.changeauthapply.get

根据token获取变更认证申请信息 */
type AlibabaNazcaTokenChangeauthapplyGetAPIRequest struct {
	model.Params
	// token
	_token string
}

// New
