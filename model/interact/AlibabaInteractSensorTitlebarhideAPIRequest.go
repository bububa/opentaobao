package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorTitlebarhideAPIRequest 隐藏titleBar API请求
// alibaba.interact.sensor.titlebarhide
//
// 隐藏titleBar
type AlibabaInteractSensorTitlebarhideAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorTitlebarhideRequest 初始化AlibabaInteractSensorTitlebarhideAPIRequest对象
func NewAlibabaInteractSensorTitlebarhideRequest() *AlibabaInteractSensorTitlebarhideAPIRequest {
	return &AlibabaInteractSensorTitlebarhideAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorTitlebarhideAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorTitlebarhideAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.titlebarhide"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorTitlebarhideAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorTitlebarhideAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractSensorTitlebarhideAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorTitlebarhideRequest()
	},
}

// GetAlibabaInteractSensorTitlebarhideRequest 从 sync.Pool 获取 AlibabaInteractSensorTitlebarhideAPIRequest
func GetAlibabaInteractSensorTitlebarhideAPIRequest() *AlibabaInteractSensorTitlebarhideAPIRequest {
	return poolAlibabaInteractSensorTitlebarhideAPIRequest.Get().(*AlibabaInteractSensorTitlebarhideAPIRequest)
}

// ReleaseAlibabaInteractSensorTitlebarhideAPIRequest 将 AlibabaInteractSensorTitlebarhideAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorTitlebarhideAPIRequest(v *AlibabaInteractSensorTitlebarhideAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorTitlebarhideAPIRequest.Put(v)
}
