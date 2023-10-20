package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorShakeAPIRequest 摇一摇 API请求
// alibaba.interact.sensor.shake
//
// 摇一摇
type AlibabaInteractSensorShakeAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorShakeRequest 初始化AlibabaInteractSensorShakeAPIRequest对象
func NewAlibabaInteractSensorShakeRequest() *AlibabaInteractSensorShakeAPIRequest {
	return &AlibabaInteractSensorShakeAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorShakeAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorShakeAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.shake"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorShakeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorShakeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractSensorShakeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorShakeRequest()
	},
}

// GetAlibabaInteractSensorShakeRequest 从 sync.Pool 获取 AlibabaInteractSensorShakeAPIRequest
func GetAlibabaInteractSensorShakeAPIRequest() *AlibabaInteractSensorShakeAPIRequest {
	return poolAlibabaInteractSensorShakeAPIRequest.Get().(*AlibabaInteractSensorShakeAPIRequest)
}

// ReleaseAlibabaInteractSensorShakeAPIRequest 将 AlibabaInteractSensorShakeAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorShakeAPIRequest(v *AlibabaInteractSensorShakeAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorShakeAPIRequest.Put(v)
}
