package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询品牌下的成为会员规则 APIResponse
alibaba.alsc.crm.rule.queryjoinmrule

查询品牌下的成为会员规则
*/
type AlibabaAlscCrmRuleQueryjoinmruleAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmRuleQueryjoinmruleResponse `json:"alibaba_alsc_crm_rule_queryjoinmrule_response,omitempty"`
}

type AlibabaAlscCrmRuleQueryjoinmruleResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
