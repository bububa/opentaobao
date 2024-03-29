package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorTakephotoAPIRequest takePhoto API请求
// alibaba.interact.sensor.takephoto
//
// 客户端takePhoto
type AlibabaInteractSensorTakephotoAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorTakephotoRequest 初始化AlibabaInteractSensorTakephotoAPIRequest对象
func NewAlibabaInteractSensorTakephotoRequest() *AlibabaInteractSensorTakephotoAPIRequest {
	return &AlibabaInteractSensorTakephotoAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorTakephotoAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorTakephotoAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.takephoto"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorTakephotoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorTakephotoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractSensorTakephotoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorTakephotoRequest()
	},
}

// GetAlibabaInteractSensorTakephotoRequest 从 sync.Pool 获取 AlibabaInteractSensorTakephotoAPIRequest
func GetAlibabaInteractSensorTakephotoAPIRequest() *AlibabaInteractSensorTakephotoAPIRequest {
	return poolAlibabaInteractSensorTakephotoAPIRequest.Get().(*AlibabaInteractSensorTakephotoAPIRequest)
}

// ReleaseAlibabaInteractSensorTakephotoAPIRequest 将 AlibabaInteractSensorTakephotoAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorTakephotoAPIRequest(v *AlibabaInteractSensorTakephotoAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorTakephotoAPIRequest.Put(v)
}
