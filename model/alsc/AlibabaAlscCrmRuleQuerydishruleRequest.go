package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询品牌下的入会菜品规则 APIRequest
alibaba.alsc.crm.rule.querydishrule

查询品牌下的入会菜品规则
*/
type AlibabaAlscCrmRuleQuerydishruleRequest struct {
    model.Params

    // 系统自动生成
    paramPlanRuleQueryOpenReq   *PlanRuleQueryOpenReq 

}

func NewAlibabaAlscCrmRuleQuerydishruleRequest() *AlibabaAlscCrmRuleQuerydishruleRequest{
    return &AlibabaAlscCrmRuleQuerydishruleRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmRuleQuerydishruleRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.rule.querydishrule"
}

func (r AlibabaAlscCrmRuleQuerydishruleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmRuleQuerydishruleRequest) SetParamPlanRuleQueryOpenReq(paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
    r.paramPlanRuleQueryOpenReq = paramPlanRuleQueryOpenReq
    r.Set("param_plan_rule_query_open_req", paramPlanRuleQueryOpenReq)
    return nil
}

func (r AlibabaAlscCrmRuleQuerydishruleRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
    return r.paramPlanRuleQueryOpenReq
}

