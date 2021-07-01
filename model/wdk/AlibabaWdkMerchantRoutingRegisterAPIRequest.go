package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家注册更新路由信息 API请求
alibaba.wdk.merchant.routing.register

商家注册更新路由信息
*/
type AlibabaWdkMerchantRoutingRegisterAPIRequest struct {
    model.Params
    // 接口入参
    _merchantRoutingInfoRegister   *MerchantRoutingInfoRegisterDo
}

// 初始化AlibabaWdkMerchantRoutingRegisterAPIRequest对象
func NewAlibabaWdkMerchantRoutingRegisterRequest() *AlibabaWdkMerchantRoutingRegisterAPIRequest{
    return &AlibabaWdkMerchantRoutingRegisterAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMerchantRoutingRegisterAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.merchant.routing.register"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMerchantRoutingRegisterAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MerchantRoutingInfoRegister Setter
// 接口入参
func (r *AlibabaWdkMerchantRoutingRegisterAPIRequest) SetMerchantRoutingInfoRegister(_merchantRoutingInfoRegister *MerchantRoutingInfoRegisterDo) error {
    r._merchantRoutingInfoRegister = _merchantRoutingInfoRegister
    r.Set("merchant_routing_info_register", _merchantRoutingInfoRegister)
    return nil
}

// MerchantRoutingInfoRegister Getter
func (r AlibabaWdkMerchantRoutingRegisterAPIRequest) GetMerchantRoutingInfoRegister() *MerchantRoutingInfoRegisterDo {
    return r._merchantRoutingInfoRegister
}
