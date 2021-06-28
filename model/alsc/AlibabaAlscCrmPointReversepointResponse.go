package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
积分消费回退 APIResponse
alibaba.alsc.crm.point.reversepoint

积分消费回退
*/
type AlibabaAlscCrmPointReversepointAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmPointReversepointResponse `json:"alibaba_alsc_crm_point_reversepoint_response,omitempty"` 
    AlibabaAlscCrmPointReversepointResponse
}

/* model for simplify = false
type AlibabaAlscCrmPointReversepointResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmPointReversepointResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
