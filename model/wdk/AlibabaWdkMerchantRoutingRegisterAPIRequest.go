package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMerchantRoutingRegisterAPIRequest 商家注册更新路由信息 API请求
// alibaba.wdk.merchant.routing.register
//
// 商家注册更新路由信息
type AlibabaWdkMerchantRoutingRegisterAPIRequest struct {
	model.Params
	// 接口入参
	_merchantRoutingInfoRegister *MerchantRoutingInfoRegisterDo
}

// NewAlibabaWdkMerchantRoutingRegisterRequest 初始化AlibabaWdkMerchantRoutingRegisterAPIRequest对象
func NewAlibabaWdkMerchantRoutingRegisterRequest() *AlibabaWdkMerchantRoutingRegisterAPIRequest {
	return &AlibabaWdkMerchantRoutingRegisterAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMerchantRoutingRegisterAPIRequest) Reset() {
	r._merchantRoutingInfoRegister = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMerchantRoutingRegisterAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.merchant.routing.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMerchantRoutingRegisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMerchantRoutingRegisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMerchantRoutingInfoRegister is MerchantRoutingInfoRegister Setter
// 接口入参
func (r *AlibabaWdkMerchantRoutingRegisterAPIRequest) SetMerchantRoutingInfoRegister(_merchantRoutingInfoRegister *MerchantRoutingInfoRegisterDo) error {
	r._merchantRoutingInfoRegister = _merchantRoutingInfoRegister
	r.Set("merchant_routing_info_register", _merchantRoutingInfoRegister)
	return nil
}

// GetMerchantRoutingInfoRegister MerchantRoutingInfoRegister Getter
func (r AlibabaWdkMerchantRoutingRegisterAPIRequest) GetMerchantRoutingInfoRegister() *MerchantRoutingInfoRegisterDo {
	return r._merchantRoutingInfoRegister
}

var poolAlibabaWdkMerchantRoutingRegisterAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMerchantRoutingRegisterRequest()
	},
}

// GetAlibabaWdkMerchantRoutingRegisterRequest 从 sync.Pool 获取 AlibabaWdkMerchantRoutingRegisterAPIRequest
func GetAlibabaWdkMerchantRoutingRegisterAPIRequest() *AlibabaWdkMerchantRoutingRegisterAPIRequest {
	return poolAlibabaWdkMerchantRoutingRegisterAPIRequest.Get().(*AlibabaWdkMerchantRoutingRegisterAPIRequest)
}

// ReleaseAlibabaWdkMerchantRoutingRegisterAPIRequest 将 AlibabaWdkMerchantRoutingRegisterAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMerchantRoutingRegisterAPIRequest(v *AlibabaWdkMerchantRoutingRegisterAPIRequest) {
	v.Reset()
	poolAlibabaWdkMerchantRoutingRegisterAPIRequest.Put(v)
}
