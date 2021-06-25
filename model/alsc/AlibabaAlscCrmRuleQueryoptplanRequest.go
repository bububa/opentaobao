package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询运营计划 APIRequest
alibaba.alsc.crm.rule.queryoptplan

查询运营计划
*/
type AlibabaAlscCrmRuleQueryoptplanRequest struct {
    model.Params

    // 请求参数
    planRuleQueryOpenRequest   *PlanRuleQueryOpenReq 

}

func NewAlibabaAlscCrmRuleQueryoptplanRequest() *AlibabaAlscCrmRuleQueryoptplanRequest{
    return &AlibabaAlscCrmRuleQueryoptplanRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmRuleQueryoptplanRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.rule.queryoptplan"
}

func (r AlibabaAlscCrmRuleQueryoptplanRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmRuleQueryoptplanRequest) SetPlanRuleQueryOpenRequest(planRuleQueryOpenRequest *PlanRuleQueryOpenReq) error {
    r.planRuleQueryOpenRequest = planRuleQueryOpenRequest
    r.Set("plan_rule_query_open_request", planRuleQueryOpenRequest)
    return nil
}

func (r AlibabaAlscCrmRuleQueryoptplanRequest) GetPlanRuleQueryOpenRequest() *PlanRuleQueryOpenReq {
    return r.planRuleQueryOpenRequest
}

