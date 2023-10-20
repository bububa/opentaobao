package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorMakeupAPIRequest 美妆虚拟试装 API请求
// alibaba.interact.sensor.makeup
//
// 手机淘宝美妆类目虚拟试妆权限，客户端能力（JS－API）
type AlibabaInteractSensorMakeupAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorMakeupRequest 初始化AlibabaInteractSensorMakeupAPIRequest对象
func NewAlibabaInteractSensorMakeupRequest() *AlibabaInteractSensorMakeupAPIRequest {
	return &AlibabaInteractSensorMakeupAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorMakeupAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorMakeupAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.makeup"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorMakeupAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorMakeupAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractSensorMakeupAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorMakeupRequest()
	},
}

// GetAlibabaInteractSensorMakeupRequest 从 sync.Pool 获取 AlibabaInteractSensorMakeupAPIRequest
func GetAlibabaInteractSensorMakeupAPIRequest() *AlibabaInteractSensorMakeupAPIRequest {
	return poolAlibabaInteractSensorMakeupAPIRequest.Get().(*AlibabaInteractSensorMakeupAPIRequest)
}

// ReleaseAlibabaInteractSensorMakeupAPIRequest 将 AlibabaInteractSensorMakeupAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorMakeupAPIRequest(v *AlibabaInteractSensorMakeupAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorMakeupAPIRequest.Put(v)
}
