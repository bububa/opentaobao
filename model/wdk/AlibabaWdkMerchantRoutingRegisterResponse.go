package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家注册更新路由信息 APIResponse
alibaba.wdk.merchant.routing.register

商家注册更新路由信息
*/
type AlibabaWdkMerchantRoutingRegisterAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMerchantRoutingRegisterResponse `json:"alibaba_wdk_merchant_routing_register_response,omitempty"` 
    AlibabaWdkMerchantRoutingRegisterResponse
}

/* model for simplify = false
type AlibabaWdkMerchantRoutingRegisterResponse struct {

    // 返回结果
    
    Result  *struct {
        AlibabaWdkMerchantRoutingRegisterApiResult  *AlibabaWdkMerchantRoutingRegisterApiResult `json:"alibaba_wdk_merchant_routing_register_api_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMerchantRoutingRegisterResponse struct {

    // 返回结果
    Result   *AlibabaWdkMerchantRoutingRegisterApiResult `json:"result,omitempty"`

}
