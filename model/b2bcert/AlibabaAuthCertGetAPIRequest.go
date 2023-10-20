package b2bcert

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAuthCertGetAPIRequest 获取证书数据 API请求
// alibaba.auth.cert.get
//
// 获取证书数据
type AlibabaAuthCertGetAPIRequest struct {
	model.Params
	// 认证商
	_provider string
	// 证书数据
	_receiveInfo string
}

// NewAlibabaAuthCertGetRequest 初始化AlibabaAuthCertGetAPIRequest对象
func NewAlibabaAuthCertGetRequest() *AlibabaAuthCertGetAPIRequest {
	return &AlibabaAuthCertGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAuthCertGetAPIRequest) Reset() {
	r._provider = ""
	r._receiveInfo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAuthCertGetAPIRequest) GetApiMethodName() string {
	return "alibaba.auth.cert.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAuthCertGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAuthCertGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProvider is Provider Setter
// 认证商
func (r *AlibabaAuthCertGetAPIRequest) SetProvider(_provider string) error {
	r._provider = _provider
	r.Set("provider", _provider)
	return nil
}

// GetProvider Provider Getter
func (r AlibabaAuthCertGetAPIRequest) GetProvider() string {
	return r._provider
}

// SetReceiveInfo is ReceiveInfo Setter
// 证书数据
func (r *AlibabaAuthCertGetAPIRequest) SetReceiveInfo(_receiveInfo string) error {
	r._receiveInfo = _receiveInfo
	r.Set("receive_info", _receiveInfo)
	return nil
}

// GetReceiveInfo ReceiveInfo Getter
func (r AlibabaAuthCertGetAPIRequest) GetReceiveInfo() string {
	return r._receiveInfo
}

var poolAlibabaAuthCertGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAuthCertGetRequest()
	},
}

// GetAlibabaAuthCertGetRequest 从 sync.Pool 获取 AlibabaAuthCertGetAPIRequest
func GetAlibabaAuthCertGetAPIRequest() *AlibabaAuthCertGetAPIRequest {
	return poolAlibabaAuthCertGetAPIRequest.Get().(*AlibabaAuthCertGetAPIRequest)
}

// ReleaseAlibabaAuthCertGetAPIRequest 将 AlibabaAuthCertGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAuthCertGetAPIRequest(v *AlibabaAuthCertGetAPIRequest) {
	v.Reset()
	poolAlibabaAuthCertGetAPIRequest.Put(v)
}
