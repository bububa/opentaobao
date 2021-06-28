package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询物理卡 APIResponse
alibaba.alsc.crm.card.qryphysical

查询物理卡
*/
type AlibabaAlscCrmCardQryphysicalAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmCardQryphysicalResponse `json:"alibaba_alsc_crm_card_qryphysical_response,omitempty"` 
    AlibabaAlscCrmCardQryphysicalResponse
}

/* model for simplify = false
type AlibabaAlscCrmCardQryphysicalResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmCardQryphysicalResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
