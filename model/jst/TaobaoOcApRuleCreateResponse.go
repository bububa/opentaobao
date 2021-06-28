package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建分账规则 APIResponse
taobao.oc.ap.rule.create

OC分账业务功能支持，用于创建分账规则
*/
type TaobaoOcApRuleCreateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOcApRuleCreateResponse `json:"oc_ap_rule_create_response,omitempty"` 
    TaobaoOcApRuleCreateResponse
}

/* model for simplify = false
type TaobaoOcApRuleCreateResponse struct {

    // 规则id
    
    RuleId   int64 `json:"rule_id,omitempty"`
    

    // 错误描述
    
    ErrorDescription   string `json:"error_description,omitempty"`
    

    // 表示方法是否正常执行成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoOcApRuleCreateResponse struct {

    // 规则id
    RuleId   int64 `json:"rule_id,omitempty"`

    // 错误描述
    ErrorDescription   string `json:"error_description,omitempty"`

    // 表示方法是否正常执行成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
