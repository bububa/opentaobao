package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询品牌下的会员日规则 APIResponse
alibaba.alsc.crm.rule.querymdayerule

查询品牌下的会员日规则
*/
type AlibabaAlscCrmRuleQuerymdayeruleAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmRuleQuerymdayeruleResponse `json:"alibaba_alsc_crm_rule_querymdayerule_response,omitempty"` 
    AlibabaAlscCrmRuleQuerymdayeruleResponse
}

/* model for simplify = false
type AlibabaAlscCrmRuleQuerymdayeruleResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmRuleQuerymdayeruleResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
