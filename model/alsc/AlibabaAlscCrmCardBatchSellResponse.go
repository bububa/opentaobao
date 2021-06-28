package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量开卡（售卡） APIResponse
alibaba.alsc.crm.card.batch.sell

批量开卡（售卡）
*/
type AlibabaAlscCrmCardBatchSellAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmCardBatchSellResponse
}

type AlibabaAlscCrmCardBatchSellResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_card_batch_sell_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
