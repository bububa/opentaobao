package security

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpCloudRphitAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.cloud.rphit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpCloudRphitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
