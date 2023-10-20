package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorGmediaAPIRequest gmedia API请求
// alibaba.interact.sensor.gmedia
//
// 媒体功能
type AlibabaInteractSensorGmediaAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorGmediaRequest 初始化AlibabaInteractSensorGmediaAPIRequest对象
func NewAlibabaInteractSensorGmediaRequest() *AlibabaInteractSensorGmediaAPIRequest {
	return &AlibabaInteractSensorGmediaAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorGmediaAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorGmediaAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.gmedia"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorGmediaAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorGmediaAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractSensorGmediaAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorGmediaRequest()
	},
}

// GetAlibabaInteractSensorGmediaRequest 从 sync.Pool 获取 AlibabaInteractSensorGmediaAPIRequest
func GetAlibabaInteractSensorGmediaAPIRequest() *AlibabaInteractSensorGmediaAPIRequest {
	return poolAlibabaInteractSensorGmediaAPIRequest.Get().(*AlibabaInteractSensorGmediaAPIRequest)
}

// ReleaseAlibabaInteractSensorGmediaAPIRequest 将 AlibabaInteractSensorGmediaAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorGmediaAPIRequest(v *AlibabaInteractSensorGmediaAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorGmediaAPIRequest.Put(v)
}
