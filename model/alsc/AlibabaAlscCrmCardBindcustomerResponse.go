package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
卡号绑定顾客 APIResponse
alibaba.alsc.crm.card.bindcustomer

为卡号绑定顾客
*/
type AlibabaAlscCrmCardBindcustomerAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmCardBindcustomerResponse `json:"alibaba_alsc_crm_card_bindcustomer_response,omitempty"` 
    AlibabaAlscCrmCardBindcustomerResponse
}

/* model for simplify = false
type AlibabaAlscCrmCardBindcustomerResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmCardBindcustomerResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
