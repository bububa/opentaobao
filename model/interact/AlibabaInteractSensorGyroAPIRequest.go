package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorGyroAPIRequest 陀螺仪 API请求
// alibaba.interact.sensor.gyro
//
// 客户端陀螺仪
type AlibabaInteractSensorGyroAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorGyroRequest 初始化AlibabaInteractSensorGyroAPIRequest对象
func NewAlibabaInteractSensorGyroRequest() *AlibabaInteractSensorGyroAPIRequest {
	return &AlibabaInteractSensorGyroAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorGyroAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorGyroAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.gyro"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorGyroAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorGyroAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractSensorGyroAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorGyroRequest()
	},
}

// GetAlibabaInteractSensorGyroRequest 从 sync.Pool 获取 AlibabaInteractSensorGyroAPIRequest
func GetAlibabaInteractSensorGyroAPIRequest() *AlibabaInteractSensorGyroAPIRequest {
	return poolAlibabaInteractSensorGyroAPIRequest.Get().(*AlibabaInteractSensorGyroAPIRequest)
}

// ReleaseAlibabaInteractSensorGyroAPIRequest 将 AlibabaInteractSensorGyroAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorGyroAPIRequest(v *AlibabaInteractSensorGyroAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorGyroAPIRequest.Put(v)
}
