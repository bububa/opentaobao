package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询日消耗预算 APIResponse
alibaba.scbp.account.budget.get

查询日消耗预算
*/
type AlibabaScbpAccountBudgetGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAccountBudgetGetResponse `json:"alibaba_scbp_account_budget_get_response,omitempty"` 
    AlibabaScbpAccountBudgetGetResponse
}

/* model for simplify = false
type AlibabaScbpAccountBudgetGetResponse struct {

    // 返回今日预算
    
    Budget   string `json:"budget,omitempty"`
    

}
*/

type AlibabaScbpAccountBudgetGetResponse struct {

    // 返回今日预算
    Budget   string `json:"budget,omitempty"`

}
