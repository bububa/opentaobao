package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询品牌下的会员成长规则 APIResponse
alibaba.alsc.crm.rule.querygrowrule

查询品牌下的会员成长规则
*/
type AlibabaAlscCrmRuleQuerygrowruleAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alsc_crm_rule_querygrowrule_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"