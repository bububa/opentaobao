package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorUiAPIRequest 基本ui操作 API请求
// alibaba.interact.sensor.ui
//
// Weex 基本UI操作
type AlibabaInteractSensorUiAPIRequest struct {
	model.Params
	// 仅作客户端鉴权使用，不会发送接收请求
	_unNamed string
}

// NewAlibabaInteractSensorUiRequest 初始化AlibabaInteractSensorUiAPIRequest对象
func NewAlibabaInteractSensorUiRequest() *AlibabaInteractSensorUiAPIRequest {
	return &AlibabaInteractSensorUiAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorUiAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.ui"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorUiAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorUiAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUnNamed is UnNamed Setter
// 仅作客户端鉴权使用，不会发送接收请求
func (r *AlibabaInteractSensorUiAPIRequest) SetUnNamed(_unNamed string) error {
	r._unNamed = _unNamed
	r.Set("un_named", _unNamed)
	return nil
}

// GetUnNamed UnNamed Getter
func (r AlibabaInteractSensorUiAPIRequest) GetUnNamed() string {
	return r._unNamed
}
