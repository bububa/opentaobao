package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
标准开卡流程 APIResponse
alibaba.alsc.crm.card.open

标准开卡流程
*/
type AlibabaAlscCrmCardOpenAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmCardOpenResponse `json:"alibaba_alsc_crm_card_open_response,omitempty"` 
    AlibabaAlscCrmCardOpenResponse
}

/* model for simplify = false
type AlibabaAlscCrmCardOpenResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmCardOpenResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
