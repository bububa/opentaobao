package util

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorUiAPIRequest) Reset() {
	r._unNamed = ""
	r.Params.ToZero()
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

var poolAlibabaInteractSensorUiAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorUiRequest()
	},
}

// GetAlibabaInteractSensorUiRequest 从 sync.Pool 获取 AlibabaInteractSensorUiAPIRequest
func GetAlibabaInteractSensorUiAPIRequest() *AlibabaInteractSensorUiAPIRequest {
	return poolAlibabaInteractSensorUiAPIRequest.Get().(*AlibabaInteractSensorUiAPIRequest)
}

// ReleaseAlibabaInteractSensorUiAPIRequest 将 AlibabaInteractSensorUiAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorUiAPIRequest(v *AlibabaInteractSensorUiAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorUiAPIRequest.Put(v)
}
