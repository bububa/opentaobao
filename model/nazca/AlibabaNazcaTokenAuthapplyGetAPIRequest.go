package nazca

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNazcaTokenAuthapplyGetAPIRequest 根据token获取认证申请信息 API请求
// alibaba.nazca.token.authapply.get
//
// 根据token获取认证申请信息
type AlibabaNazcaTokenAuthapplyGetAPIRequest struct {
	model.Params
	// token
	_token string
}

// NewAlibabaNazcaTokenAuthapplyGetRequest 初始化AlibabaNazcaTokenAuthapplyGetAPIRequest对象
func NewAlibabaNazcaTokenAuthapplyGetRequest() *AlibabaNazcaTokenAuthapplyGetAPIRequest {
	return &AlibabaNazcaTokenAuthapplyGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaNazcaTokenAuthapplyGetAPIRequest) Reset() {
	r._token = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNazcaTokenAuthapplyGetAPIRequest) GetApiMethodName() string {
	return "alibaba.nazca.token.authapply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNazcaTokenAuthapplyGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaNazcaTokenAuthapplyGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// token
func (r *AlibabaNazcaTokenAuthapplyGetAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaNazcaTokenAuthapplyGetAPIRequest) GetToken() string {
	return r._token
}

var poolAlibabaNazcaTokenAuthapplyGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaNazcaTokenAuthapplyGetRequest()
	},
}

// GetAlibabaNazcaTokenAuthapplyGetRequest 从 sync.Pool 获取 AlibabaNazcaTokenAuthapplyGetAPIRequest
func GetAlibabaNazcaTokenAuthapplyGetAPIRequest() *AlibabaNazcaTokenAuthapplyGetAPIRequest {
	return poolAlibabaNazcaTokenAuthapplyGetAPIRequest.Get().(*AlibabaNazcaTokenAuthapplyGetAPIRequest)
}

// ReleaseAlibabaNazcaTokenAuthapplyGetAPIRequest 将 AlibabaNazcaTokenAuthapplyGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaNazcaTokenAuthapplyGetAPIRequest(v *AlibabaNazcaTokenAuthapplyGetAPIRequest) {
	v.Reset()
	poolAlibabaNazcaTokenAuthapplyGetAPIRequest.Put(v)
}
