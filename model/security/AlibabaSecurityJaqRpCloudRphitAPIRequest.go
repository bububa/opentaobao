package security

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpCloudRphitAPIRequest 实人认证云服务日志打点 API请求
// alibaba.security.jaq.rp.cloud.rphit
//
// 聚安全实人认证日志打点接口
type AlibabaSecurityJaqRpCloudRphitAPIRequest struct {
	model.Params
	// xxx
	_content string
}

// NewAlibabaSecurityJaqRpCloudRphitRequest 初始化AlibabaSecurityJaqRpCloudRphitAPIRequest对象
func NewAlibabaSecurityJaqRpCloudRphitRequest() *AlibabaSecurityJaqRpCloudRphitAPIRequest {
	return &AlibabaSecurityJaqRpCloudRphitAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSecurityJaqRpCloudRphitAPIRequest) Reset() {
	r._content = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpCloudRphitAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.cloud.rphit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpCloudRphitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqRpCloudRphitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContent is Content Setter
// xxx
func (r *AlibabaSecurityJaqRpCloudRphitAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r AlibabaSecurityJaqRpCloudRphitAPIRequest) GetContent() string {
	return r._content
}

var poolAlibabaSecurityJaqRpCloudRphitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSecurityJaqRpCloudRphitRequest()
	},
}

// GetAlibabaSecurityJaqRpCloudRphitRequest 从 sync.Pool 获取 AlibabaSecurityJaqRpCloudRphitAPIRequest
func GetAlibabaSecurityJaqRpCloudRphitAPIRequest() *AlibabaSecurityJaqRpCloudRphitAPIRequest {
	return poolAlibabaSecurityJaqRpCloudRphitAPIRequest.Get().(*AlibabaSecurityJaqRpCloudRphitAPIRequest)
}

// ReleaseAlibabaSecurityJaqRpCloudRphitAPIRequest 将 AlibabaSecurityJaqRpCloudRphitAPIRequest 放入 sync.Pool
func ReleaseAlibabaSecurityJaqRpCloudRphitAPIRequest(v *AlibabaSecurityJaqRpCloudRphitAPIRequest) {
	v.Reset()
	poolAlibabaSecurityJaqRpCloudRphitAPIRequest.Put(v)
}
