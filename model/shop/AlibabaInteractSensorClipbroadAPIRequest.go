package shop

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorClipbroadAPIRequest Weex页面设置或读取剪切板 API请求
// alibaba.interact.sensor.clipbroad
//
// Weex页面设置或读取剪切板
type AlibabaInteractSensorClipbroadAPIRequest struct {
	model.Params
	// 客户端鉴权使用，实际不会发送或接收数据
	_unNamed string
}

// NewAlibabaInteractSensorClipbroadRequest 初始化AlibabaInteractSensorClipbroadAPIRequest对象
func NewAlibabaInteractSensorClipbroadRequest() *AlibabaInteractSensorClipbroadAPIRequest {
	return &AlibabaInteractSensorClipbroadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorClipbroadAPIRequest) Reset() {
	r._unNamed = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorClipbroadAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.clipbroad"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorClipbroadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorClipbroadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUnNamed is UnNamed Setter
// 客户端鉴权使用，实际不会发送或接收数据
func (r *AlibabaInteractSensorClipbroadAPIRequest) SetUnNamed(_unNamed string) error {
	r._unNamed = _unNamed
	r.Set("un_named", _unNamed)
	return nil
}

// GetUnNamed UnNamed Getter
func (r AlibabaInteractSensorClipbroadAPIRequest) GetUnNamed() string {
	return r._unNamed
}

var poolAlibabaInteractSensorClipbroadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorClipbroadRequest()
	},
}

// GetAlibabaInteractSensorClipbroadRequest 从 sync.Pool 获取 AlibabaInteractSensorClipbroadAPIRequest
func GetAlibabaInteractSensorClipbroadAPIRequest() *AlibabaInteractSensorClipbroadAPIRequest {
	return poolAlibabaInteractSensorClipbroadAPIRequest.Get().(*AlibabaInteractSensorClipbroadAPIRequest)
}

// ReleaseAlibabaInteractSensorClipbroadAPIRequest 将 AlibabaInteractSensorClipbroadAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorClipbroadAPIRequest(v *AlibabaInteractSensorClipbroadAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorClipbroadAPIRequest.Put(v)
}
