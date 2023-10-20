package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorGutilAPIRequest canvas工具包 API请求
// alibaba.interact.sensor.gutil
//
// canvas工具包
type AlibabaInteractSensorGutilAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorGutilRequest 初始化AlibabaInteractSensorGutilAPIRequest对象
func NewAlibabaInteractSensorGutilRequest() *AlibabaInteractSensorGutilAPIRequest {
	return &AlibabaInteractSensorGutilAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorGutilAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorGutilAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.gutil"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorGutilAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorGutilAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractSensorGutilAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorGutilRequest()
	},
}

// GetAlibabaInteractSensorGutilRequest 从 sync.Pool 获取 AlibabaInteractSensorGutilAPIRequest
func GetAlibabaInteractSensorGutilAPIRequest() *AlibabaInteractSensorGutilAPIRequest {
	return poolAlibabaInteractSensorGutilAPIRequest.Get().(*AlibabaInteractSensorGutilAPIRequest)
}

// ReleaseAlibabaInteractSensorGutilAPIRequest 将 AlibabaInteractSensorGutilAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorGutilAPIRequest(v *AlibabaInteractSensorGutilAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorGutilAPIRequest.Put(v)
}
