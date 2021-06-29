package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询品牌下的会员成长规则 API请求
alibaba.alsc.crm.rule.querygrowrule

查询品牌下的会员成长规则
*/
type AlibabaAlscCrmRuleQuerygrowruleRequest struct {
    model.Params
    // 系统自动生成
    _paramPlanRuleQueryOpenReq   *PlanRuleQueryOpenReq
}

// 初始化AlibabaAlscCrmRuleQuerygrowruleRequest对象
func NewAlibabaAlscCrmRuleQuerygrowruleRequest() *AlibabaAlscCrmRuleQuerygrowruleRequest{
    return &AlibabaAlscCrmRuleQuerygrowruleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRuleQuerygrowruleRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.rule.querygrowrule"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRuleQuerygrowruleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamPlanRuleQueryOpenReq Setter
// 系统自动生成
func (r *AlibabaAlscCrmRuleQuerygrowruleRequest) SetParamPlanRuleQueryOpenReq(_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
    r._paramPlanRuleQueryOpenReq = _paramPlanRuleQueryOpenReq
    r.Set("param_plan_rule_query_open_req", _paramPlanRuleQueryOpenReq)
    return nil
}

// ParamPlanRuleQueryOpenReq Getter
func (r AlibabaAlscCrmRuleQuerygrowruleRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
    return r._paramPlanRuleQueryOpenReq
}
