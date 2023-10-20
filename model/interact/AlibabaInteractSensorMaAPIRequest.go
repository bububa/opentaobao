package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorMaAPIRequest 码相关API API请求
// alibaba.interact.sensor.ma
//
// 码相关API
type AlibabaInteractSensorMaAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorMaRequest 初始化AlibabaInteractSensorMaAPIRequest对象
func NewAlibabaInteractSensorMaRequest() *AlibabaInteractSensorMaAPIRequest {
	return &AlibabaInteractSensorMaAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorMaAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorMaAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.ma"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorMaAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorMaAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractSensorMaAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorMaRequest()
	},
}

// GetAlibabaInteractSensorMaRequest 从 sync.Pool 获取 AlibabaInteractSensorMaAPIRequest
func GetAlibabaInteractSensorMaAPIRequest() *AlibabaInteractSensorMaAPIRequest {
	return poolAlibabaInteractSensorMaAPIRequest.Get().(*AlibabaInteractSensorMaAPIRequest)
}

// ReleaseAlibabaInteractSensorMaAPIRequest 将 AlibabaInteractSensorMaAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorMaAPIRequest(v *AlibabaInteractSensorMaAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorMaAPIRequest.Put(v)
}
