package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorAudioAPIRequest 声音 API请求
// alibaba.interact.sensor.audio
//
// 客户端声音
type AlibabaInteractSensorAudioAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorAudioRequest 初始化AlibabaInteractSensorAudioAPIRequest对象
func NewAlibabaInteractSensorAudioRequest() *AlibabaInteractSensorAudioAPIRequest {
	return &AlibabaInteractSensorAudioAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorAudioAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorAudioAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.audio"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorAudioAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorAudioAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractSensorAudioAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorAudioRequest()
	},
}

// GetAlibabaInteractSensorAudioRequest 从 sync.Pool 获取 AlibabaInteractSensorAudioAPIRequest
func GetAlibabaInteractSensorAudioAPIRequest() *AlibabaInteractSensorAudioAPIRequest {
	return poolAlibabaInteractSensorAudioAPIRequest.Get().(*AlibabaInteractSensorAudioAPIRequest)
}

// ReleaseAlibabaInteractSensorAudioAPIRequest 将 AlibabaInteractSensorAudioAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorAudioAPIRequest(v *AlibabaInteractSensorAudioAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorAudioAPIRequest.Put(v)
}
