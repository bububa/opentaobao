package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询品牌下的会员价规则 API请求
alibaba.alsc.crm.rule.querympricerule

查询品牌下的会员价规则
*/
type AlibabaAlscCrmRuleQuerympriceruleAPIRequest struct {
    model.Params
    // 系统自动生成
    _paramPlanRuleQueryOpenReq   *PlanRuleQueryOpenReq
}

// 初始化AlibabaAlscCrmRuleQuerympriceruleAPIRequest对象
func NewAlibabaAlscCrmRuleQuerympriceruleRequest() *AlibabaAlscCrmRuleQuerympriceruleAPIRequest{
    return &AlibabaAlscCrmRuleQuerympriceruleAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRuleQuerympriceruleAPIRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.rule.querympricerule"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRuleQuerympriceruleAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamPlanRuleQueryOpenReq Setter
// 系统自动生成
func (r *AlibabaAlscCrmRuleQuerympriceruleAPIRequest) SetParamPlanRuleQueryOpenReq(_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
    r._paramPlanRuleQueryOpenReq = _paramPlanRuleQueryOpenReq
    r.Set("param_plan_rule_query_open_req", _paramPlanRuleQueryOpenReq)
    return nil
}

// ParamPlanRuleQueryOpenReq Getter
func (r AlibabaAlscCrmRuleQuerympriceruleAPIRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
    return r._paramPlanRuleQueryOpenReq
}
