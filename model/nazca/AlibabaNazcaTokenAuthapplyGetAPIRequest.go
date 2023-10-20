package nazca

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabanazcatokenauthapplygetAPIRequest 根据token获取认证申请信息 API请求
// alibaba.nazca.token.authapply.get
//
// 根据token获取认证申请信息
type AlibabanazcatokenauthapplygetAPIRequest struct {
	model.Params
	// token
	_token string
}

// NewAlibabanazcatokenauthapplygetRequest 初始化AlibabanazcatokenauthapplygetAPIRequest对象
func NewAlibabanazcatokenauthapplygetRequest() *AlibabanazcatokenauthapplygetAPIRequest {
	return &AlibabanazcatokenauthapplygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabanazcatokenauthapplygetAPIRequest) GetApiMethodName() string {
	return "alibaba.nazca.token.authapply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabanazcatokenauthapplygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabanazcatokenauthapplygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// token
func (r *AlibabanazcatokenauthapplygetAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabanazcatokenauthapplygetAPIRequest) GetToken() string {
	return r._token
}
