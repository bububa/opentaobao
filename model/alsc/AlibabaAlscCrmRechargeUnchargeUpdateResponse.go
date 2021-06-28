package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
充值退款 APIResponse
alibaba.alsc.crm.recharge.uncharge.update

充值退款
*/
type AlibabaAlscCrmRechargeUnchargeUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmRechargeUnchargeUpdateResponse `json:"alibaba_alsc_crm_recharge_uncharge_update_response,omitempty"` 
    AlibabaAlscCrmRechargeUnchargeUpdateResponse
}

/* model for simplify = false
type AlibabaAlscCrmRechargeUnchargeUpdateResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmRechargeUnchargeUpdateResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
