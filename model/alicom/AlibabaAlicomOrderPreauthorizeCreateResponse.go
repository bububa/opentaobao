package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
业务办理结果 APIResponse
alibaba.alicom.order.preauthorize.create

授授权:签约结果通知
*/
type AlibabaAlicomOrderPreauthorizeCreateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlicomOrderPreauthorizeCreateResponse `json:"alibaba_alicom_order_preauthorize_create_response,omitempty"` 
    AlibabaAlicomOrderPreauthorizeCreateResponse
}

/* model for simplify = false
type AlibabaAlicomOrderPreauthorizeCreateResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlicomOrderPreauthorizeCreateResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
