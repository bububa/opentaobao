package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询品牌下的会员日规则 APIRequest
alibaba.alsc.crm.rule.querymdayerule

查询品牌下的会员日规则
*/
type AlibabaAlscCrmRuleQuerymdayeruleRequest struct {
    model.Params

    // 系统自动生成
    paramPlanRuleQueryOpenReq   *PlanRuleQueryOpenReq 

}

func NewAlibabaAlscCrmRuleQuerymdayeruleRequest() *AlibabaAlscCrmRuleQuerymdayeruleRequest{
    return &AlibabaAlscCrmRuleQuerymdayeruleRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmRuleQuerymdayeruleRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.rule.querymdayerule"
}

func (r AlibabaAlscCrmRuleQuerymdayeruleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmRuleQuerymdayeruleRequest) SetParamPlanRuleQueryOpenReq(paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
    r.paramPlanRuleQueryOpenReq = paramPlanRuleQueryOpenReq
    r.Set("param_plan_rule_query_open_req", paramPlanRuleQueryOpenReq)
    return nil
}

func (r AlibabaAlscCrmRuleQuerymdayeruleRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
    return r.paramPlanRuleQueryOpenReq
}

