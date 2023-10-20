package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensorclipbroadAPIRequest Weex页面设置或读取剪切板 API请求
// alibaba.interact.sensor.clipbroad
//
// Weex页面设置或读取剪切板
type AlibabainteractsensorclipbroadAPIRequest struct {
	model.Params
	// 客户端鉴权使用，实际不会发送或接收数据
	_unNamed string
}

// NewAlibabainteractsensorclipbroadRequest 初始化AlibabainteractsensorclipbroadAPIRequest对象
func NewAlibabainteractsensorclipbroadRequest() *AlibabainteractsensorclipbroadAPIRequest {
	return &AlibabainteractsensorclipbroadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensorclipbroadAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.clipbroad"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensorclipbroadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensorclipbroadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUnNamed is UnNamed Setter
// 客户端鉴权使用，实际不会发送或接收数据
func (r *AlibabainteractsensorclipbroadAPIRequest) SetUnNamed(_unNamed string) error {
	r._unNamed = _unNamed
	r.Set("un_named", _unNamed)
	return nil
}

// GetUnNamed UnNamed Getter
func (r AlibabainteractsensorclipbroadAPIRequest) GetUnNamed() string {
	return r._unNamed
}
