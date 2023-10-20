package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmerchantroutingregisterAPIRequest 商家注册更新路由信息 API请求
// alibaba.wdk.merchant.routing.register
//
// 商家注册更新路由信息
type AlibabawdkmerchantroutingregisterAPIRequest struct {
	model.Params
	// 接口入参
	_merchantRoutingInfoRegister *MerchantRoutingInfoRegisterDo
}

// NewAlibabawdkmerchantroutingregisterRequest 初始化AlibabawdkmerchantroutingregisterAPIRequest对象
func NewAlibabawdkmerchantroutingregisterRequest() *AlibabawdkmerchantroutingregisterAPIRequest {
	return &AlibabawdkmerchantroutingregisterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmerchantroutingregisterAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.merchant.routing.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmerchantroutingregisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmerchantroutingregisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMerchantRoutingInfoRegister is MerchantRoutingInfoRegister Setter
// 接口入参
func (r *AlibabawdkmerchantroutingregisterAPIRequest) SetMerchantRoutingInfoRegister(_merchantRoutingInfoRegister *MerchantRoutingInfoRegisterDo) error {
	r._merchantRoutingInfoRegister = _merchantRoutingInfoRegister
	r.Set("merchant_routing_info_register", _merchantRoutingInfoRegister)
	return nil
}

// GetMerchantRoutingInfoRegister MerchantRoutingInfoRegister Getter
func (r AlibabawdkmerchantroutingregisterAPIRequest) GetMerchantRoutingInfoRegister() *MerchantRoutingInfoRegisterDo {
	return r._merchantRoutingInfoRegister
}
