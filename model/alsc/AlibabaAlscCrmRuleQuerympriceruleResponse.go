package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询品牌下的会员价规则 APIResponse
alibaba.alsc.crm.rule.querympricerule

查询品牌下的会员价规则
*/
type AlibabaAlscCrmRuleQuerympriceruleAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmRuleQuerympriceruleResponse
}

type AlibabaAlscCrmRuleQuerympriceruleResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_rule_querympricerule_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
