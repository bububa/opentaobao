package nazca

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNazcaTokenIssuecertapplyGetAPIRequest 根据token获取出证申请信息 API请求
// alibaba.nazca.token.issuecertapply.get
//
// 根据token获取出证申请信息
type AlibabaNazcaTokenIssuecertapplyGetAPIRequest struct {
	model.Params
	// token
	_token string
}

// NewAlibabaNazcaTokenIssuecertapplyGetRequest 初始化AlibabaNazcaTokenIssuecertapplyGetAPIRequest对象
func NewAlibabaNazcaTokenIssuecertapplyGetRequest() *AlibabaNazcaTokenIssuecertapplyGetAPIRequest {
	return &AlibabaNazcaTokenIssuecertapplyGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNazcaTokenIssuecertapplyGetAPIRequest) GetApiMethodName() string {
	return "alibaba.nazca.token.issuecertapply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNazcaTokenIssuecertapplyGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetToken is Token Setter
// token
func (r *AlibabaNazcaTokenIssuecertapplyGetAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaNazcaTokenIssuecertapplyGetAPIRequest) GetToken() string {
	return r._token
}
