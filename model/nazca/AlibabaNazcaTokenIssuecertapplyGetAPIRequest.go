package nazca

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabanazcatokenissuecertapplygetAPIRequest 根据token获取出证申请信息 API请求
// alibaba.nazca.token.issuecertapply.get
//
// 根据token获取出证申请信息
type AlibabanazcatokenissuecertapplygetAPIRequest struct {
	model.Params
	// token
	_token string
}

// NewAlibabanazcatokenissuecertapplygetRequest 初始化AlibabanazcatokenissuecertapplygetAPIRequest对象
func NewAlibabanazcatokenissuecertapplygetRequest() *AlibabanazcatokenissuecertapplygetAPIRequest {
	return &AlibabanazcatokenissuecertapplygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabanazcatokenissuecertapplygetAPIRequest) GetApiMethodName() string {
	return "alibaba.nazca.token.issuecertapply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabanazcatokenissuecertapplygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabanazcatokenissuecertapplygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// token
func (r *AlibabanazcatokenissuecertapplygetAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabanazcatokenissuecertapplygetAPIRequest) GetToken() string {
	return r._token
}
