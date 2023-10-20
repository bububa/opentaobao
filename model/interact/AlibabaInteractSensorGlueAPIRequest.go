package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorGlueAPIRequest 视频播放 API请求
// alibaba.interact.sensor.glue
//
// 视频播放
type AlibabaInteractSensorGlueAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorGlueRequest 初始化AlibabaInteractSensorGlueAPIRequest对象
func NewAlibabaInteractSensorGlueRequest() *AlibabaInteractSensorGlueAPIRequest {
	return &AlibabaInteractSensorGlueAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorGlueAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorGlueAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.glue"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorGlueAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorGlueAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractSensorGlueAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorGlueRequest()
	},
}

// GetAlibabaInteractSensorGlueRequest 从 sync.Pool 获取 AlibabaInteractSensorGlueAPIRequest
func GetAlibabaInteractSensorGlueAPIRequest() *AlibabaInteractSensorGlueAPIRequest {
	return poolAlibabaInteractSensorGlueAPIRequest.Get().(*AlibabaInteractSensorGlueAPIRequest)
}

// ReleaseAlibabaInteractSensorGlueAPIRequest 将 AlibabaInteractSensorGlueAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorGlueAPIRequest(v *AlibabaInteractSensorGlueAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorGlueAPIRequest.Put(v)
}
