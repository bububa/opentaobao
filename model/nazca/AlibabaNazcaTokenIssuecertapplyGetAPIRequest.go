package nazca

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNazcaTokenIssuecertapplyGetAPIRequest
根据token获取出证申请信息 API请求
alibaba.nazca.token.issuecertapply.get

根据token获取出证申请信息 */
type AlibabaNazcaTokenIssuecertapplyGetAPIRequest struct {
	model.Params
	// token
	_token string
}

// New
