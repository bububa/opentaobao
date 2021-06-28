package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询品牌下的入会菜品规则 APIResponse
alibaba.alsc.crm.rule.querydishrule

查询品牌下的入会菜品规则
*/
type AlibabaAlscCrmRuleQuerydishruleAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmRuleQuerydishruleResponse `json:"alibaba_alsc_crm_rule_querydishrule_response,omitempty"` 
    AlibabaAlscCrmRuleQuerydishruleResponse
}

/* model for simplify = false
type AlibabaAlscCrmRuleQuerydishruleResponse struct {

    // 分页返回模型
    
    Result  *struct {
        CommonPageResult  *CommonPageResult `json:"common_page_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmRuleQuerydishruleResponse struct {

    // 分页返回模型
    Result   *CommonPageResult `json:"result,omitempty"`

}
