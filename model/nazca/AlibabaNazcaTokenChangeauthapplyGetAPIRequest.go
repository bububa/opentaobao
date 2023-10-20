package nazca

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaNazcaTokenChangeauthapplyGetAPIRequest) Reset() {
	r._token = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNazcaTokenChangeauthapplyGetAPIRequest) GetApiMethodName() string {
	return "alibaba.nazca.token.changeauthapply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNazcaTokenChangeauthapplyGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaNazcaTokenChangeauthapplyGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// token
func (r *AlibabaNazcaTokenChangeauthapplyGetAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaNazcaTokenChangeauthapplyGetAPIRequest) GetToken() string {
	return r._token
}

var poolAlibabaNazcaTokenChangeauthapplyGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaNazcaTokenChangeauthapplyGetRequest()
	},
}

// GetAlibabaNazcaTokenChangeauthapplyGetRequest 从 sync.Pool 获取 AlibabaNazcaTokenChangeauthapplyGetAPIRequest
func GetAlibabaNazcaTokenChangeauthapplyGetAPIRequest() *AlibabaNazcaTokenChangeauthapplyGetAPIRequest {
	return poolAlibabaNazcaTokenChangeauthapplyGetAPIRequest.Get().(*AlibabaNazcaTokenChangeauthapplyGetAPIRequest)
}

// ReleaseAlibabaNazcaTokenChangeauthapplyGetAPIRequest 将 AlibabaNazcaTokenChangeauthapplyGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaNazcaTokenChangeauthapplyGetAPIRequest(v *AlibabaNazcaTokenChangeauthapplyGetAPIRequest) {
	v.Reset()
	poolAlibabaNazcaTokenChangeauthapplyGetAPIRequest.Put(v)
}
