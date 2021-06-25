package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建分账规则 APIRequest
taobao.oc.ap.rule.create

OC分账业务功能支持，用于创建分账规则
*/
type TaobaoOcApRuleCreateRequest struct {
    model.Params

    // 传入比例数组后者金额数组
    divisionRule   []Number 

    // 规则描述相关扩展信息，divisonRule的值包含（"byAmount" 或者 "byPercentage"），excutionPeriod的值包含（ "month" 或者 "day" 或者 "now"）
    extAttributes   string 

}

func NewTaobaoOcApRuleCreateRequest() *TaobaoOcApRuleCreateRequest{
    return &TaobaoOcApRuleCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOcApRuleCreateRequest) GetApiMethodName() string {
    return "taobao.oc.ap.rule.create"
}

func (r TaobaoOcApRuleCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOcApRuleCreateRequest) SetDivisionRule(divisionRule []Number) error {
    r.divisionRule = divisionRule
    r.Set("division_rule", divisionRule)
    return nil
}

func (r TaobaoOcApRuleCreateRequest) GetDivisionRule() []Number {
    return r.divisionRule
}

func (r *TaobaoOcApRuleCreateRequest) SetExtAttributes(extAttributes string) error {
    r.extAttributes = extAttributes
    r.Set("ext_attributes", extAttributes)
    return nil
}

func (r TaobaoOcApRuleCreateRequest) GetExtAttributes() string {
    return r.extAttributes
}

