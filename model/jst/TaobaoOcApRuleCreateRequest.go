package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建分账规则 API请求
taobao.oc.ap.rule.create

OC分账业务功能支持，用于创建分账规则
*/
type TaobaoOcApRuleCreateRequest struct {
    model.Params
    // 传入比例数组后者金额数组
    divisionRule   []int64
    // 规则描述相关扩展信息，divisonRule的值包含（"byAmount" 或者 "byPercentage"），excutionPeriod的值包含（ "month" 或者 "day" 或者 "now"）
    extAttributes   string
}

// 初始化TaobaoOcApRuleCreateRequest对象
func NewTaobaoOcApRuleCreateRequest() *TaobaoOcApRuleCreateRequest{
    return &TaobaoOcApRuleCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOcApRuleCreateRequest) GetApiMethodName() string {
    return "taobao.oc.ap.rule.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOcApRuleCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DivisionRule Setter
// 传入比例数组后者金额数组
func (r *TaobaoOcApRuleCreateRequest) SetDivisionRule(divisionRule []int64) error {
    r.divisionRule = divisionRule
    r.Set("division_rule", divisionRule)
    return nil
}

// DivisionRule Getter
func (r TaobaoOcApRuleCreateRequest) GetDivisionRule() []int64 {
    return r.divisionRule
}
// ExtAttributes Setter
// 规则描述相关扩展信息，divisonRule的值包含（"byAmount" 或者 "byPercentage"），excutionPeriod的值包含（ "month" 或者 "day" 或者 "now"）
func (r *TaobaoOcApRuleCreateRequest) SetExtAttributes(extAttributes string) error {
    r.extAttributes = extAttributes
    r.Set("ext_attributes", extAttributes)
    return nil
}

// ExtAttributes Getter
func (r TaobaoOcApRuleCreateRequest) GetExtAttributes() string {
    return r.extAttributes
}
