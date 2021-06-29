package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询运营计划 API请求
alibaba.alsc.crm.rule.queryoptplan

查询运营计划
*/
type AlibabaAlscCrmRuleQueryoptplanRequest struct {
    model.Params
    // 请求参数
    planRuleQueryOpenRequest   *PlanRuleQueryOpenReq
}

// 初始化AlibabaAlscCrmRuleQueryoptplanRequest对象
func NewAlibabaAlscCrmRuleQueryoptplanRequest() *AlibabaAlscCrmRuleQueryoptplanRequest{
    return &AlibabaAlscCrmRuleQueryoptplanRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRuleQueryoptplanRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.rule.queryoptplan"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRuleQueryoptplanRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PlanRuleQueryOpenRequest Setter
// 请求参数
func (r *AlibabaAlscCrmRuleQueryoptplanRequest) SetPlanRuleQueryOpenRequest(planRuleQueryOpenRequest *PlanRuleQueryOpenReq) error {
    r.planRuleQueryOpenRequest = planRuleQueryOpenRequest
    r.Set("plan_rule_query_open_request", planRuleQueryOpenRequest)
    return nil
}

// PlanRuleQueryOpenRequest Getter
func (r AlibabaAlscCrmRuleQueryoptplanRequest) GetPlanRuleQueryOpenRequest() *PlanRuleQueryOpenReq {
    return r.planRuleQueryOpenRequest
}
