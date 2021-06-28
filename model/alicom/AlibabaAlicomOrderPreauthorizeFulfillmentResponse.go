package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
履约结果 APIResponse
alibaba.alicom.order.preauthorize.fulfillment

预授权-履约结果
*/
type AlibabaAlicomOrderPreauthorizeFulfillmentAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlicomOrderPreauthorizeFulfillmentResponse `json:"alibaba_alicom_order_preauthorize_fulfillment_response,omitempty"` 
    AlibabaAlicomOrderPreauthorizeFulfillmentResponse
}

/* model for simplify = false
type AlibabaAlicomOrderPreauthorizeFulfillmentResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlicomOrderPreauthorizeFulfillmentResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
