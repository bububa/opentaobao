package nazca

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNazcaTokenChangeauthapplyGetAPIRequest 根据token获取变更认证申请信息 API请求
// alibaba.nazca.token.changeauthapply.get
//
// 根据token获取变更认证申请信息
type AlibabaNazcaTokenChangeauthapplyGetAPIRequest struct {
	model.Params
	// token
	_token string
}

// NewAlibabaNazcaTokenChangeauthapplyGetRequest 初始化AlibabaNazcaTokenChangeauthapplyGetAPIRequest对象
func NewAlibabaNazcaTokenChangeauthapplyGetRequest() *AlibabaNazcaTokenChangeauthapplyGetAPIRequest {
	return &AlibabaNazcaTokenChangeauthapplyGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNazcaTokenChangeauthapplyGetAPIRequest) GetApiMethodName() string {
	return "alibaba.nazca.token.changeauthapply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNazcaTokenChangeauthapplyGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Token Setter
// token
func (r *AlibabaNazcaTokenChangeauthapplyGetAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// Get Token Getter
func (r AlibabaNazcaTokenChangeauthapplyGetAPIRequest) GetToken() string {
	return r._token
}
