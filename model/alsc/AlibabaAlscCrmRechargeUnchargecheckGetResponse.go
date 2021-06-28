package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
储值账户退充值校验 APIResponse
alibaba.alsc.crm.recharge.unchargecheck.get

储值账户退充值校验接口
*/
type AlibabaAlscCrmRechargeUnchargecheckGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmRechargeUnchargecheckGetResponse `json:"alibaba_alsc_crm_recharge_unchargecheck_get_response,omitempty"` 
    AlibabaAlscCrmRechargeUnchargecheckGetResponse
}

/* model for simplify = false
type AlibabaAlscCrmRechargeUnchargecheckGetResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmRechargeUnchargecheckGetResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
