package security

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpFetchmaterialAPIRequest 聚安全实人认证获取结果接口 API请求
// alibaba.security.jaq.rp.fetchmaterial
//
// 聚安全实人认证获取结果接口
type AlibabaSecurityJaqRpFetchmaterialAPIRequest struct {
	model.Params
	// 消息服务推送的key
	_securityKey string
}

// NewAlibabaSecurityJaqRpFetchmaterialRequest 初始化AlibabaSecurityJaqRpFetchmaterialAPIRequest对象
func NewAlibabaSecurityJaqRpFetchmaterialRequest() *AlibabaSecurityJaqRpFetchmaterialAPIRequest {
	return &AlibabaSecurityJaqRpFetchmaterialAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSecurityJaqRpFetchmaterialAPIRequest) Reset() {
	r._securityKey = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpFetchmaterialAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.fetchmaterial"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpFetchmaterialAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqRpFetchmaterialAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSecurityKey is SecurityKey Setter
// 消息服务推送的key
func (r *AlibabaSecurityJaqRpFetchmaterialAPIRequest) SetSecurityKey(_securityKey string) error {
	r._securityKey = _securityKey
	r.Set("security_key", _securityKey)
	return nil
}

// GetSecurityKey SecurityKey Getter
func (r AlibabaSecurityJaqRpFetchmaterialAPIRequest) GetSecurityKey() string {
	return r._securityKey
}

var poolAlibabaSecurityJaqRpFetchmaterialAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSecurityJaqRpFetchmaterialRequest()
	},
}

// GetAlibabaSecurityJaqRpFetchmaterialRequest 从 sync.Pool 获取 AlibabaSecurityJaqRpFetchmaterialAPIRequest
func GetAlibabaSecurityJaqRpFetchmaterialAPIRequest() *AlibabaSecurityJaqRpFetchmaterialAPIRequest {
	return poolAlibabaSecurityJaqRpFetchmaterialAPIRequest.Get().(*AlibabaSecurityJaqRpFetchmaterialAPIRequest)
}

// ReleaseAlibabaSecurityJaqRpFetchmaterialAPIRequest 将 AlibabaSecurityJaqRpFetchmaterialAPIRequest 放入 sync.Pool
func ReleaseAlibabaSecurityJaqRpFetchmaterialAPIRequest(v *AlibabaSecurityJaqRpFetchmaterialAPIRequest) {
	v.Reset()
	poolAlibabaSecurityJaqRpFetchmaterialAPIRequest.Put(v)
}
