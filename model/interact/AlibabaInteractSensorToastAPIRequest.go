package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorToastAPIRequest toast API请求
// alibaba.interact.sensor.toast
//
// toast提示
type AlibabaInteractSensorToastAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorToastRequest 初始化AlibabaInteractSensorToastAPIRequest对象
func NewAlibabaInteractSensorToastRequest() *AlibabaInteractSensorToastAPIRequest {
	return &AlibabaInteractSensorToastAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorToastAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorToastAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.toast"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorToastAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorToastAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractSensorToastAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorToastRequest()
	},
}

// GetAlibabaInteractSensorToastRequest 从 sync.Pool 获取 AlibabaInteractSensorToastAPIRequest
func GetAlibabaInteractSensorToastAPIRequest() *AlibabaInteractSensorToastAPIRequest {
	return poolAlibabaInteractSensorToastAPIRequest.Get().(*AlibabaInteractSensorToastAPIRequest)
}

// ReleaseAlibabaInteractSensorToastAPIRequest 将 AlibabaInteractSensorToastAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorToastAPIRequest(v *AlibabaInteractSensorToastAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorToastAPIRequest.Put(v)
}
