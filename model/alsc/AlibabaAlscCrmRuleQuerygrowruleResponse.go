package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询品牌下的会员成长规则 APIResponse
alibaba.alsc.crm.rule.querygrowrule

查询品牌下的会员成长规则
*/
type AlibabaAlscCrmRuleQuerygrowruleAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmRuleQuerygrowruleResponse `json:"alibaba_alsc_crm_rule_querygrowrule_response,omitempty"` 
    AlibabaAlscCrmRuleQuerygrowruleResponse
}

/* model for simplify = false
type AlibabaAlscCrmRuleQuerygrowruleResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmRuleQuerygrowruleResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
