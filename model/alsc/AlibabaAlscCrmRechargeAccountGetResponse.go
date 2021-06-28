package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询储值账户信息 APIResponse
alibaba.alsc.crm.recharge.account.get

查询储值账户信息接口
*/
type AlibabaAlscCrmRechargeAccountGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmRechargeAccountGetResponse `json:"alibaba_alsc_crm_recharge_account_get_response,omitempty"` 
    AlibabaAlscCrmRechargeAccountGetResponse
}

/* model for simplify = false
type AlibabaAlscCrmRechargeAccountGetResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmRechargeAccountGetResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
