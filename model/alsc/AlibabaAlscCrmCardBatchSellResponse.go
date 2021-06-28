package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量开卡（售卡） APIResponse
alibaba.alsc.crm.card.batch.sell

批量开卡（售卡）
*/
type AlibabaAlscCrmCardBatchSellAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmCardBatchSellResponse `json:"alibaba_alsc_crm_card_batch_sell_response,omitempty"` 
    AlibabaAlscCrmCardBatchSellResponse
}

/* model for simplify = false
type AlibabaAlscCrmCardBatchSellResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmCardBatchSellResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
