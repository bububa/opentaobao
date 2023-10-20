package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorPopwindowAPIRequest popwindow API请求
// alibaba.interact.sensor.popwindow
//
// popwindow
type AlibabaInteractSensorPopwindowAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorPopwindowRequest 初始化AlibabaInteractSensorPopwindowAPIRequest对象
func NewAlibabaInteractSensorPopwindowRequest() *AlibabaInteractSensorPopwindowAPIRequest {
	return &AlibabaInteractSensorPopwindowAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorPopwindowAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorPopwindowAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.popwindow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorPopwindowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorPopwindowAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractSensorPopwindowAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorPopwindowRequest()
	},
}

// GetAlibabaInteractSensorPopwindowRequest 从 sync.Pool 获取 AlibabaInteractSensorPopwindowAPIRequest
func GetAlibabaInteractSensorPopwindowAPIRequest() *AlibabaInteractSensorPopwindowAPIRequest {
	return poolAlibabaInteractSensorPopwindowAPIRequest.Get().(*AlibabaInteractSensorPopwindowAPIRequest)
}

// ReleaseAlibabaInteractSensorPopwindowAPIRequest 将 AlibabaInteractSensorPopwindowAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorPopwindowAPIRequest(v *AlibabaInteractSensorPopwindowAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorPopwindowAPIRequest.Put(v)
}
