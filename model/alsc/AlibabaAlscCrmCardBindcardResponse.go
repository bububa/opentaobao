package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
绑定物理卡 APIResponse
alibaba.alsc.crm.card.bindcard

将会员卡和实例物理卡绑定在一起
*/
type AlibabaAlscCrmCardBindcardAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmCardBindcardResponse `json:"alibaba_alsc_crm_card_bindcard_response,omitempty"` 
    AlibabaAlscCrmCardBindcardResponse
}

/* model for simplify = false
type AlibabaAlscCrmCardBindcardResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmCardBindcardResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
