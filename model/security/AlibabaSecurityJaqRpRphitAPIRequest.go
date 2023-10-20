package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqrprphitAPIRequest 聚安全-实人认证日志打点接口 API请求
// alibaba.security.jaq.rp.rphit
//
// 聚安全实人认证日志打点接口
type AlibabasecurityjaqrprphitAPIRequest struct {
	model.Params
	// xxx
	_content string
}

// NewAlibabasecurityjaqrprphitRequest 初始化AlibabasecurityjaqrprphitAPIRequest对象
func NewAlibabasecurityjaqrprphitRequest() *AlibabasecurityjaqrprphitAPIRequest {
	return &AlibabasecurityjaqrprphitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqrprphitAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.rphit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqrprphitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqrprphitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContent is Content Setter
// xxx
func (r *AlibabasecurityjaqrprphitAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r AlibabasecurityjaqrprphitAPIRequest) GetContent() string {
	return r._content
}
