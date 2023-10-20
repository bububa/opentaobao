package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensoruiAPIRequest 基本ui操作 API请求
// alibaba.interact.sensor.ui
//
// Weex 基本UI操作
type AlibabainteractsensoruiAPIRequest struct {
	model.Params
	// 仅作客户端鉴权使用，不会发送接收请求
	_unNamed string
}

// NewAlibabainteractsensoruiRequest 初始化AlibabainteractsensoruiAPIRequest对象
func NewAlibabainteractsensoruiRequest() *AlibabainteractsensoruiAPIRequest {
	return &AlibabainteractsensoruiAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensoruiAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.ui"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensoruiAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensoruiAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUnNamed is UnNamed Setter
// 仅作客户端鉴权使用，不会发送接收请求
func (r *AlibabainteractsensoruiAPIRequest) SetUnNamed(_unNamed string) error {
	r._unNamed = _unNamed
	r.Set("un_named", _unNamed)
	return nil
}

// GetUnNamed UnNamed Getter
func (r AlibabainteractsensoruiAPIRequest) GetUnNamed() string {
	return r._unNamed
}
