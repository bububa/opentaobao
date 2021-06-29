package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询品牌下的成为会员规则 APIResponse
alibaba.alsc.crm.rule.queryjoinmrule

查询品牌下的成为会员规则
*/
type AlibabaAlscCrmRuleQueryjoinmruleAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmRuleQueryjoinmruleResponse
}

type AlibabaAlscCrmRuleQueryjoinmruleResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_rule_queryjoinmrule_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
