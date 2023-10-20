package yunos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyuncosmogatewayinvokeAPIRequest alios cosmo服务调用 API请求
// aliyun.cosmo.gateway.invoke
//
// AliOS cosmo服务分发平台对外调用接口
type AliyuncosmogatewayinvokeAPIRequest struct {
	model.Params
	// 请求上下文参数
	_context *RdamContext
	// 请求对象
	_rdamRequest *RdamGenericRequest
}

// NewAliyuncosmogatewayinvokeRequest 初始化AliyuncosmogatewayinvokeAPIRequest对象
func NewAliyuncosmogatewayinvokeRequest() *AliyuncosmogatewayinvokeAPIRequest {
	return &AliyuncosmogatewayinvokeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyuncosmogatewayinvokeAPIRequest) GetApiMethodName() string {
	return "aliyun.cosmo.gateway.invoke"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyuncosmogatewayinvokeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyuncosmogatewayinvokeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContext is Context Setter
// 请求上下文参数
func (r *AliyuncosmogatewayinvokeAPIRequest) SetContext(_context *RdamContext) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r AliyuncosmogatewayinvokeAPIRequest) GetContext() *RdamContext {
	return r._context
}

// SetRdamRequest is RdamRequest Setter
// 请求对象
func (r *AliyuncosmogatewayinvokeAPIRequest) SetRdamRequest(_rdamRequest *RdamGenericRequest) error {
	r._rdamRequest = _rdamRequest
	r.Set("rdam_request", _rdamRequest)
	return nil
}

// GetRdamRequest RdamRequest Getter
func (r AliyuncosmogatewayinvokeAPIRequest) GetRdamRequest() *RdamGenericRequest {
	return r._rdamRequest
}
