package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpRphitAPIRequest 聚安全-实人认证日志打点接口 API请求
// alibaba.security.jaq.rp.rphit
//
// 聚安全实人认证日志打点接口
type AlibabaSecurityJaqRpRphitAPIRequest struct {
	model.Params
	// xxx
	_content string
}

// NewAlibabaSecurityJaqRpRphitRequest 初始化AlibabaSecurityJaqRpRphitAPIRequest对象
func NewAlibabaSecurityJaqRpRphitRequest() *AlibabaSecurityJaqRpRphitAPIRequest {
	return &AlibabaSecurityJaqRpRphitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpRphitAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.rphit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpRphitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqRpRphitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContent is Content Setter
// xxx
func (r *AlibabaSecurityJaqRpRphitAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r AlibabaSecurityJaqRpRphitAPIRequest) GetContent() string {
	return r._content
}
