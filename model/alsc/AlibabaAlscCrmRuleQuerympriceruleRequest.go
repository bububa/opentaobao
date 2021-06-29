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
type AlibabaAlscCrmRuleQuerympriceruleRequest struct {
    model.Params
    // 系统自动生成
    _paramPlanRuleQueryOpenReq   *PlanRuleQueryOpenReq
}

// 初始化AlibabaAlscCrmRuleQuerympriceruleRequest对象
func NewAlibabaAlscCrmRuleQuerympriceruleRequest() *AlibabaAlscCrmRuleQuerympriceruleRequest{
    return &AlibabaAlscCrmRuleQuerympriceruleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRuleQuerympriceruleRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.rule.querympricerule"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRuleQuerympriceruleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamPlanRuleQueryOpenReq Setter
// 系统自动生成
func (r *AlibabaAlscCrmRuleQuerympriceruleRequest) SetParamPlanRuleQueryOpenReq(_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
    r._paramPlanRuleQueryOpenReq = _paramPlanRuleQueryOpenReq
    r.Set("param_plan_rule_query_open_req", _paramPlanRuleQueryOpenReq)
    return nil
}

// ParamPlanRuleQueryOpenReq Getter
func (r AlibabaAlscCrmRuleQuerympriceruleRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
    return r._paramPlanRuleQueryOpenReq
}
