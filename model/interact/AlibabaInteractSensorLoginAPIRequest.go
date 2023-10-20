package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorLoginAPIRequest 获取登陆页面 API请求
// alibaba.interact.sensor.login
//
// 获取登陆页面
type AlibabaInteractSensorLoginAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorLoginRequest 初始化AlibabaInteractSensorLoginAPIRequest对象
func NewAlibabaInteractSensorLoginRequest() *AlibabaInteractSensorLoginAPIRequest {
	return &AlibabaInteractSensorLoginAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorLoginAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorLoginAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.login"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorLoginAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorLoginAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractSensorLoginAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorLoginRequest()
	},
}

// GetAlibabaInteractSensorLoginRequest 从 sync.Pool 获取 AlibabaInteractSensorLoginAPIRequest
func GetAlibabaInteractSensorLoginAPIRequest() *AlibabaInteractSensorLoginAPIRequest {
	return poolAlibabaInteractSensorLoginAPIRequest.Get().(*AlibabaInteractSensorLoginAPIRequest)
}

// ReleaseAlibabaInteractSensorLoginAPIRequest 将 AlibabaInteractSensorLoginAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorLoginAPIRequest(v *AlibabaInteractSensorLoginAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorLoginAPIRequest.Put(v)
}
