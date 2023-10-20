package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqrpcloudrphitAPIRequest 实人认证云服务日志打点 API请求
// alibaba.security.jaq.rp.cloud.rphit
//
// 聚安全实人认证日志打点接口
type AlibabasecurityjaqrpcloudrphitAPIRequest struct {
	model.Params
	// xxx
	_content string
}

// NewAlibabasecurityjaqrpcloudrphitRequest 初始化AlibabasecurityjaqrpcloudrphitAPIRequest对象
func NewAlibabasecurityjaqrpcloudrphitRequest() *AlibabasecurityjaqrpcloudrphitAPIRequest {
	return &AlibabasecurityjaqrpcloudrphitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqrpcloudrphitAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.cloud.rphit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqrpcloudrphitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqrpcloudrphitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContent is Content Setter
// xxx
func (r *AlibabasecurityjaqrpcloudrphitAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r AlibabasecurityjaqrpcloudrphitAPIRequest) GetContent() string {
	return r._content
}
