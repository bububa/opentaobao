package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorOpenwindowAPIRequest 客户端打开新页面 API请求
// alibaba.interact.sensor.openwindow
//
// 客户端打开新页面
type AlibabaInteractSensorOpenwindowAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorOpenwindowRequest 初始化AlibabaInteractSensorOpenwindowAPIRequest对象
func NewAlibabaInteractSensorOpenwindowRequest() *AlibabaInteractSensorOpenwindowAPIRequest {
	return &AlibabaInteractSensorOpenwindowAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorOpenwindowAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorOpenwindowAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.openwindow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorOpenwindowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorOpenwindowAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractSensorOpenwindowAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorOpenwindowRequest()
	},
}

// GetAlibabaInteractSensorOpenwindowRequest 从 sync.Pool 获取 AlibabaInteractSensorOpenwindowAPIRequest
func GetAlibabaInteractSensorOpenwindowAPIRequest() *AlibabaInteractSensorOpenwindowAPIRequest {
	return poolAlibabaInteractSensorOpenwindowAPIRequest.Get().(*AlibabaInteractSensorOpenwindowAPIRequest)
}

// ReleaseAlibabaInteractSensorOpenwindowAPIRequest 将 AlibabaInteractSensorOpenwindowAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorOpenwindowAPIRequest(v *AlibabaInteractSensorOpenwindowAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorOpenwindowAPIRequest.Put(v)
}
