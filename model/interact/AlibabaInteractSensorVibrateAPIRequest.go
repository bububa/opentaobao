package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorVibrateAPIRequest 客户端震动 API请求
// alibaba.interact.sensor.vibrate
//
// 客户端震动
type AlibabaInteractSensorVibrateAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorVibrateRequest 初始化AlibabaInteractSensorVibrateAPIRequest对象
func NewAlibabaInteractSensorVibrateRequest() *AlibabaInteractSensorVibrateAPIRequest {
	return &AlibabaInteractSensorVibrateAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorVibrateAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorVibrateAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.vibrate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorVibrateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorVibrateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractSensorVibrateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorVibrateRequest()
	},
}

// GetAlibabaInteractSensorVibrateRequest 从 sync.Pool 获取 AlibabaInteractSensorVibrateAPIRequest
func GetAlibabaInteractSensorVibrateAPIRequest() *AlibabaInteractSensorVibrateAPIRequest {
	return poolAlibabaInteractSensorVibrateAPIRequest.Get().(*AlibabaInteractSensorVibrateAPIRequest)
}

// ReleaseAlibabaInteractSensorVibrateAPIRequest 将 AlibabaInteractSensorVibrateAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorVibrateAPIRequest(v *AlibabaInteractSensorVibrateAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorVibrateAPIRequest.Put(v)
}
