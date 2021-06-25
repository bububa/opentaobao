package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询品牌下的会员成长规则 APIRequest
alibaba.alsc.crm.rule.querygrowrule

查询品牌下的会员成长规则
*/
type AlibabaAlscCrmRuleQuerygrowruleRequest struct {
    model.Params

    // 系统自动生成
    paramPlanRuleQueryOpenReq   *PlanRuleQueryOpenReq 

}

func NewAlibabaAlscCrmRuleQuerygrowruleRequest() *AlibabaAlscCrmRuleQuerygrowruleRequest{
    return &AlibabaAlscCrmRuleQuerygrowruleRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmRuleQuerygrowruleRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.rule.querygrowrule"
}

func (r AlibabaAlscCrmRuleQuerygrowruleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmRuleQuerygrowruleRequest) SetParamPlanRuleQueryOpenReq(paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
    r.paramPlanRuleQueryOpenReq = paramPlanRuleQueryOpenReq
    r.Set("param_plan_rule_query_open_req", paramPlanRuleQueryOpenReq)
    return nil
}

func (r AlibabaAlscCrmRuleQuerygrowruleRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
    return r.paramPlanRuleQueryOpenReq
}

