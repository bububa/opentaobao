package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分页查询储值流水 APIResponse
alibaba.alsc.crm.recharge.accountflows.get

增加分页查询储值流水接口
*/
type AlibabaAlscCrmRechargeAccountflowsGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmRechargeAccountflowsGetResponse `json:"alibaba_alsc_crm_recharge_accountflows_get_response,omitempty"` 
    AlibabaAlscCrmRechargeAccountflowsGetResponse
}

/* model for simplify = false
type AlibabaAlscCrmRechargeAccountflowsGetResponse struct {

    // 分页返回模型
    
    Result  *struct {
        CommonPageResult  *CommonPageResult `json:"common_page_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmRechargeAccountflowsGetResponse struct {

    // 分页返回模型
    Result   *CommonPageResult `json:"result,omitempty"`

}
