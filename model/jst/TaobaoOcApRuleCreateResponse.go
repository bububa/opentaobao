package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建分账规则 API返回值 
taobao.oc.ap.rule.create

OC分账业务功能支持，用于创建分账规则
*/
type TaobaoOcApRuleCreateAPIResponse struct {
    model.CommonResponse
    TaobaoOcApRuleCreateResponse
}

// 创建分账规则 成功返回结果
type TaobaoOcApRuleCreateResponse struct {
    XMLName xml.Name `xml:"oc_ap_rule_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 规则id
    RuleId   int64 `json:"rule_id,omitempty" xml:"rule_id,omitempty"`
    // 错误描述
    ErrorDescription   string `json:"error_description,omitempty" xml:"error_description,omitempty"`
    // 表示方法是否正常执行成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
