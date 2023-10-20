package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorAuthorizeAPIRequest 客户端授权页 API请求
// alibaba.interact.sensor.authorize
//
// 客户端授权页
type AlibabaInteractSensorAuthorizeAPIRequest struct {
	model.Params
}

// NewAlibabaInteractSensorAuthorizeRequest 初始化AlibabaInteractSensorAuthorizeAPIRequest对象
func NewAlibabaInteractSensorAuthorizeRequest() *AlibabaInteractSensorAuthorizeAPIRequest {
	return &AlibabaInteractSensorAuthorizeAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorAuthorizeAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorAuthorizeAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.authorize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorAuthorizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorAuthorizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractSensorAuthorizeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorAuthorizeRequest()
	},
}

// GetAlibabaInteractSensorAuthorizeRequest 从 sync.Pool 获取 AlibabaInteractSensorAuthorizeAPIRequest
func GetAlibabaInteractSensorAuthorizeAPIRequest() *AlibabaInteractSensorAuthorizeAPIRequest {
	return poolAlibabaInteractSensorAuthorizeAPIRequest.Get().(*AlibabaInteractSensorAuthorizeAPIRequest)
}

// ReleaseAlibabaInteractSensorAuthorizeAPIRequest 将 AlibabaInteractSensorAuthorizeAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorAuthorizeAPIRequest(v *AlibabaInteractSensorAuthorizeAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorAuthorizeAPIRequest.Put(v)
}
