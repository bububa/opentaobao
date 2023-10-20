package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractIsvGatewayAPIRequest isv调用gateway API请求
// alibaba.interact.isv.gateway
//
// isv能够调用jae本身的server
type AlibabaInteractIsvGatewayAPIRequest struct {
	model.Params
}

// NewAlibabaInteractIsvGatewayRequest 初始化AlibabaInteractIsvGatewayAPIRequest对象
func NewAlibabaInteractIsvGatewayRequest() *AlibabaInteractIsvGatewayAPIRequest {
	return &AlibabaInteractIsvGatewayAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractIsvGatewayAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractIsvGatewayAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.isv.gateway"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractIsvGatewayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractIsvGatewayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractIsvGatewayAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractIsvGatewayRequest()
	},
}

// GetAlibabaInteractIsvGatewayRequest 从 sync.Pool 获取 AlibabaInteractIsvGatewayAPIRequest
func GetAlibabaInteractIsvGatewayAPIRequest() *AlibabaInteractIsvGatewayAPIRequest {
	return poolAlibabaInteractIsvGatewayAPIRequest.Get().(*AlibabaInteractIsvGatewayAPIRequest)
}

// ReleaseAlibabaInteractIsvGatewayAPIRequest 将 AlibabaInteractIsvGatewayAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractIsvGatewayAPIRequest(v *AlibabaInteractIsvGatewayAPIRequest) {
	v.Reset()
	poolAlibabaInteractIsvGatewayAPIRequest.Put(v)
}
