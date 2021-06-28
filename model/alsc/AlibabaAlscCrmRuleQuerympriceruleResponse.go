package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询品牌下的会员价规则 APIResponse
alibaba.alsc.crm.rule.querympricerule

查询品牌下的会员价规则
*/
type AlibabaAlscCrmRuleQuerympriceruleAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmRuleQuerympriceruleResponse `json:"alibaba_alsc_crm_rule_querympricerule_response,omitempty"` 
    AlibabaAlscCrmRuleQuerympriceruleResponse
}

/* model for simplify = false
type AlibabaAlscCrmRuleQuerympriceruleResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmRuleQuerympriceruleResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
