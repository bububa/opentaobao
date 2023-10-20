package nazca

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaNazcaTokenIssuecertapplyGetAPIRequest) Reset() {
	r._token = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNazcaTokenIssuecertapplyGetAPIRequest) GetApiMethodName() string {
	return "alibaba.nazca.token.issuecertapply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNazcaTokenIssuecertapplyGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaNazcaTokenIssuecertapplyGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaNazcaTokenIssuecertapplyGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaNazcaTokenIssuecertapplyGetRequest()
	},
}

// GetAlibabaNazcaTokenIssuecertapplyGetRequest 从 sync.Pool 获取 AlibabaNazcaTokenIssuecertapplyGetAPIRequest
func GetAlibabaNazcaTokenIssuecertapplyGetAPIRequest() *AlibabaNazcaTokenIssuecertapplyGetAPIRequest {
	return poolAlibabaNazcaTokenIssuecertapplyGetAPIRequest.Get().(*AlibabaNazcaTokenIssuecertapplyGetAPIRequest)
}

// ReleaseAlibabaNazcaTokenIssuecertapplyGetAPIRequest 将 AlibabaNazcaTokenIssuecertapplyGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaNazcaTokenIssuecertapplyGetAPIRequest(v *AlibabaNazcaTokenIssuecertapplyGetAPIRequest) {
	v.Reset()
	poolAlibabaNazcaTokenIssuecertapplyGetAPIRequest.Put(v)
}
