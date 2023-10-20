package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorGcanvasAPIRequest gcanvas API请求
// alibaba.interact.sensor.gcanvas
//
// gcanvas 功能
type AlibabaInteractSensorGcanvasAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorGcanvasRequest 初始化AlibabaInteractSensorGcanvasAPIRequest对象
func NewAlibabaInteractSensorGcanvasRequest() *AlibabaInteractSensorGcanvasAPIRequest {
	return &AlibabaInteractSensorGcanvasAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorGcanvasAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorGcanvasAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.gcanvas"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorGcanvasAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorGcanvasAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractSensorGcanvasAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorGcanvasRequest()
	},
}

// GetAlibabaInteractSensorGcanvasRequest 从 sync.Pool 获取 AlibabaInteractSensorGcanvasAPIRequest
func GetAlibabaInteractSensorGcanvasAPIRequest() *AlibabaInteractSensorGcanvasAPIRequest {
	return poolAlibabaInteractSensorGcanvasAPIRequest.Get().(*AlibabaInteractSensorGcanvasAPIRequest)
}

// ReleaseAlibabaInteractSensorGcanvasAPIRequest 将 AlibabaInteractSensorGcanvasAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorGcanvasAPIRequest(v *AlibabaInteractSensorGcanvasAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorGcanvasAPIRequest.Put(v)
}
