package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询品牌下的成为会员规则 APIRequest
alibaba.alsc.crm.rule.queryjoinmrule

查询品牌下的成为会员规则
*/
type AlibabaAlscCrmRuleQueryjoinmruleRequest struct {
    model.Params

    // 系统自动生成
    paramPlanRuleQueryOpenReq   *PlanRuleQueryOpenReq 

}

func NewAlibabaAlscCrmRuleQueryjoinmruleRequest() *AlibabaAlscCrmRuleQueryjoinmruleRequest{
    return &AlibabaAlscCrmRuleQueryjoinmruleRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmRuleQueryjoinmruleRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.rule.queryjoinmrule"
}

func (r AlibabaAlscCrmRuleQueryjoinmruleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmRuleQueryjoinmruleRequest) SetParamPlanRuleQueryOpenReq(paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
    r.paramPlanRuleQueryOpenReq = paramPlanRuleQueryOpenReq
    r.Set("param_plan_rule_query_open_req", paramPlanRuleQueryOpenReq)
    return nil
}

func (r AlibabaAlscCrmRuleQueryjoinmruleRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
    return r.paramPlanRuleQueryOpenReq
}

