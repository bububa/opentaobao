package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询积分规则 APIResponse
alibaba.alsc.crm.point.rule.get

新增积分规则查询接口,传入includeLogicalDelete和maxUpdateTime时走同步下行逻辑不然则走普通积分规则查询接口
*/
type AlibabaAlscCrmPointRuleGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmPointRuleGetResponse `json:"alibaba_alsc_crm_point_rule_get_response,omitempty"`
}

type AlibabaAlscCrmPointRuleGetResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
