package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量激活卡 APIResponse
alibaba.alsc.crm.card.batch.active

批量激活卡
*/
type AlibabaAlscCrmCardBatchActiveAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmCardBatchActiveResponse `json:"alibaba_alsc_crm_card_batch_active_response,omitempty"` 
    AlibabaAlscCrmCardBatchActiveResponse
}

/* model for simplify = false
type AlibabaAlscCrmCardBatchActiveResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmCardBatchActiveResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
