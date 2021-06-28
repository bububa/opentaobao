package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家注册更新路由信息 APIResponse
alibaba.wdk.merchant.routing.register

商家注册更新路由信息
*/
type AlibabaWdkMerchantRoutingRegisterAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_merchant_routing_register_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *AlibabaWdkMerchantRoutingRegisterApiResult `json:"result,omitempty" xml:"