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
    Response *AlibabaAlscCrmCardBatchActiveResponse `json:"alibaba_alsc_crm_card_batch_active_response,omitempty"`
}

type AlibabaAlscCrmCardBatchActiveResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
