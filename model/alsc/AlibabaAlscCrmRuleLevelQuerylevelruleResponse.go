package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询会员等级规则 APIResponse
alibaba.alsc.crm.rule.level.querylevelrule

查询会员等级规则
*/
type AlibabaAlscCrmRuleLevelQuerylevelruleAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmRuleLevelQuerylevelruleResponse `json:"alibaba_alsc_crm_rule_level_querylevelrule_response,omitempty"` 
    AlibabaAlscCrmRuleLevelQuerylevelruleResponse
}

/* model for simplify = false
type AlibabaAlscCrmRuleLevelQuerylevelruleResponse struct {

    // 会员等级规则
    
    Result  *struct {
        CommonPageResult  *CommonPageResult `json:"common_page_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmRuleLevelQuerylevelruleResponse struct {

    // 会员等级规则
    Result   *CommonPageResult `json:"result,omitempty"`

}
