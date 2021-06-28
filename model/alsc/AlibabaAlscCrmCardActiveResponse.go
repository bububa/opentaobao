package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
标准激活卡 APIResponse
alibaba.alsc.crm.card.active

激活卡
*/
type AlibabaAlscCrmCardActiveAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmCardActiveResponse `json:"alibaba_alsc_crm_card_active_response,omitempty"` 
    AlibabaAlscCrmCardActiveResponse
}

/* model for simplify = false
type AlibabaAlscCrmCardActiveResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmCardActiveResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
