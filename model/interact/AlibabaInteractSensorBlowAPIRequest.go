package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorBlowAPIRequest 吹气 API请求
// alibaba.interact.sensor.blow
//
// 客户端吹气
type AlibabaInteractSensorBlowAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorBlowRequest 初始化AlibabaInteractSensorBlowAPIRequest对象
func NewAlibabaInteractSensorBlowRequest() *AlibabaInteractSensorBlowAPIRequest {
	return &AlibabaInteractSensorBlowAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorBlowAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorBlowAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.blow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorBlowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorBlowAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractSensorBlowAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorBlowRequest()
	},
}

// GetAlibabaInteractSensorBlowRequest 从 sync.Pool 获取 AlibabaInteractSensorBlowAPIRequest
func GetAlibabaInteractSensorBlowAPIRequest() *AlibabaInteractSensorBlowAPIRequest {
	return poolAlibabaInteractSensorBlowAPIRequest.Get().(*AlibabaInteractSensorBlowAPIRequest)
}

// ReleaseAlibabaInteractSensorBlowAPIRequest 将 AlibabaInteractSensorBlowAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorBlowAPIRequest(v *AlibabaInteractSensorBlowAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorBlowAPIRequest.Put(v)
}
