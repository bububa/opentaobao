package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询运营计划 APIResponse
alibaba.alsc.crm.rule.queryoptplan

查询运营计划
*/
type AlibabaAlscCrmRuleQueryoptplanAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmRuleQueryoptplanResponse `json:"alibaba_alsc_crm_rule_queryoptplan_response,omitempty"` 
    AlibabaAlscCrmRuleQueryoptplanResponse
}

/* model for simplify = false
type AlibabaAlscCrmRuleQueryoptplanResponse struct {

    // 请求参数
    
    Result  *struct {
        CommonPageResult  *CommonPageResult `json:"common_page_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmRuleQueryoptplanResponse struct {

    // 请求参数
    Result   *CommonPageResult `json:"result,omitempty"`

}
