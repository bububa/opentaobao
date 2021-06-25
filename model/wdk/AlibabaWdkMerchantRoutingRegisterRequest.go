package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家注册更新路由信息 APIRequest
alibaba.wdk.merchant.routing.register

商家注册更新路由信息
*/
type AlibabaWdkMerchantRoutingRegisterRequest struct {
    model.Params

    // 接口入参
    merchantRoutingInfoRegister   *MerchantRoutingInfoRegisterDo 

}

func NewAlibabaWdkMerchantRoutingRegisterRequest() *AlibabaWdkMerchantRoutingRegisterRequest{
    return &AlibabaWdkMerchantRoutingRegisterRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMerchantRoutingRegisterRequest) GetApiMethodName() string {
    return "alibaba.wdk.merchant.routing.register"
}

func (r AlibabaWdkMerchantRoutingRegisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMerchantRoutingRegisterRequest) SetMerchantRoutingInfoRegister(merchantRoutingInfoRegister *MerchantRoutingInfoRegisterDo) error {
    r.merchantRoutingInfoRegister = merchantRoutingInfoRegister
    r.Set("merchant_routing_info_register", merchantRoutingInfoRegister)
    return nil
}

func (r AlibabaWdkMerchantRoutingRegisterRequest) GetMerchantRoutingInfoRegister() *MerchantRoutingInfoRegisterDo {
    return r.merchantRoutingInfoRegister
}

