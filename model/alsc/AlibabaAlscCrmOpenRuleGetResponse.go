package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询规则 APIResponse
alibaba.alsc.crm.open.rule.get

查询会员规则
*/
type AlibabaAlscCrmOpenRuleGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmOpenRuleGetResponse `json:"alibaba_alsc_crm_open_rule_get_response,omitempty"`
}

type AlibabaAlscCrmOpenRuleGetResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}