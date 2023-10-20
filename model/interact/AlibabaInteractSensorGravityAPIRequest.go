package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorGravityAPIRequest 重力感应 API请求
// alibaba.interact.sensor.gravity
//
// native获取重力感应
type AlibabaInteractSensorGravityAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorGravityRequest 初始化AlibabaInteractSensorGravityAPIRequest对象
func NewAlibabaInteractSensorGravityRequest() *AlibabaInteractSensorGravityAPIRequest {
	return &AlibabaInteractSensorGravityAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorGravityAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorGravityAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.gravity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorGravityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorGravityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractSensorGravityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorGravityRequest()
	},
}

// GetAlibabaInteractSensorGravityRequest 从 sync.Pool 获取 AlibabaInteractSensorGravityAPIRequest
func GetAlibabaInteractSensorGravityAPIRequest() *AlibabaInteractSensorGravityAPIRequest {
	return poolAlibabaInteractSensorGravityAPIRequest.Get().(*AlibabaInteractSensorGravityAPIRequest)
}

// ReleaseAlibabaInteractSensorGravityAPIRequest 将 AlibabaInteractSensorGravityAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorGravityAPIRequest(v *AlibabaInteractSensorGravityAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorGravityAPIRequest.Put(v)
}
