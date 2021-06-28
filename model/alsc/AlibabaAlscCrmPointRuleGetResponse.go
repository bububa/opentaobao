package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询积分规则 APIResponse
alibaba.alsc.crm.point.rule.get

新增积分规则查询接口,传入includeLogicalDelete和maxUpdateTime时走同步下行逻辑不然则走普通积分规则查询接口
*/
type AlibabaAlscCrmPointRuleGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmPointRuleGetResponse
}

type AlibabaAlscCrmPointRuleGetResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_point_rule_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
