package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
储值流水详细 APIResponse
alibaba.alsc.crm.recharge.account.flowdetail.get

查询储值流水详细接口
*/
type AlibabaAlscCrmRechargeAccountFlowdetailGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmRechargeAccountFlowdetailGetResponse `json:"alibaba_alsc_crm_recharge_account_flowdetail_get_response,omitempty"` 
    AlibabaAlscCrmRechargeAccountFlowdetailGetResponse
}

/* model for simplify = false
type AlibabaAlscCrmRechargeAccountFlowdetailGetResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmRechargeAccountFlowdetailGetResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
