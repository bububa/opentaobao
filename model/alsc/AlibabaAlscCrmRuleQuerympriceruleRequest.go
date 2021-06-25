package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询品牌下的会员价规则 APIRequest
alibaba.alsc.crm.rule.querympricerule

查询品牌下的会员价规则
*/
type AlibabaAlscCrmRuleQuerympriceruleRequest struct {
    model.Params

    // 系统自动生成
    paramPlanRuleQueryOpenReq   *PlanRuleQueryOpenReq 

}

func NewAlibabaAlscCrmRuleQuerympriceruleRequest() *AlibabaAlscCrmRuleQuerympriceruleRequest{
    return &AlibabaAlscCrmRuleQuerympriceruleRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmRuleQuerympriceruleRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.rule.querympricerule"
}

func (r AlibabaAlscCrmRuleQuerympriceruleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmRuleQuerympriceruleRequest) SetParamPlanRuleQueryOpenReq(paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
    r.paramPlanRuleQueryOpenReq = paramPlanRuleQueryOpenReq
    r.Set("param_plan_rule_query_open_req", paramPlanRuleQueryOpenReq)
    return nil
}

func (r AlibabaAlscCrmRuleQuerympriceruleRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
    return r.paramPlanRuleQueryOpenReq
}

