package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询品牌下的入会菜品规则 APIResponse
alibaba.alsc.crm.rule.querydishrule

查询品牌下的入会菜品规则
*/
type AlibabaAlscCrmRuleQuerydishruleAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmRuleQuerydishruleResponse
}

type AlibabaAlscCrmRuleQuerydishruleResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_rule_querydishrule_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 分页返回模型
    
    Result   *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
