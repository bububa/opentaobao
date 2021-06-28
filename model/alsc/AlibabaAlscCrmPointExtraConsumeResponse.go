package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
积分补扣 APIResponse
alibaba.alsc.crm.point.extra.consume

积分补扣
*/
type AlibabaAlscCrmPointExtraConsumeAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmPointExtraConsumeResponse `json:"alibaba_alsc_crm_point_extra_consume_response,omitempty"` 
    AlibabaAlscCrmPointExtraConsumeResponse
}

/* model for simplify = false
type AlibabaAlscCrmPointExtraConsumeResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmPointExtraConsumeResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
