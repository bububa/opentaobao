package nazca

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabanazcatokenchangeauthapplygetAPIRequest 根据token获取变更认证申请信息 API请求
// alibaba.nazca.token.changeauthapply.get
//
// 根据token获取变更认证申请信息
type AlibabanazcatokenchangeauthapplygetAPIRequest struct {
	model.Params
	// token
	_token string
}

// NewAlibabanazcatokenchangeauthapplygetRequest 初始化AlibabanazcatokenchangeauthapplygetAPIRequest对象
func NewAlibabanazcatokenchangeauthapplygetRequest() *AlibabanazcatokenchangeauthapplygetAPIRequest {
	return &AlibabanazcatokenchangeauthapplygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabanazcatokenchangeauthapplygetAPIRequest) GetApiMethodName() string {
	return "alibaba.nazca.token.changeauthapply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabanazcatokenchangeauthapplygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabanazcatokenchangeauthapplygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// token
func (r *AlibabanazcatokenchangeauthapplygetAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabanazcatokenchangeauthapplygetAPIRequest) GetToken() string {
	return r._token
}
