package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询标签列表 APIRequest
alibaba.alsc.crm.rule.querytaglist

查询标签列表
*/
type AlibabaAlscCrmRuleQuerytaglistRequest struct {
    model.Params

    // 请求参数
    paramPlanRuleQueryOpenReq   *PlanRuleQueryOpenReq 

}

func NewAlibabaAlscCrmRuleQuerytaglistRequest() *AlibabaAlscCrmRuleQuerytaglistRequest{
    return &AlibabaAlscCrmRuleQuerytaglistRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmRuleQuerytaglistRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.rule.querytaglist"
}

func (r AlibabaAlscCrmRuleQuerytaglistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmRuleQuerytaglistRequest) SetParamPlanRuleQueryOpenReq(paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
    r.paramPlanRuleQueryOpenReq = paramPlanRuleQueryOpenReq
    r.Set("param_plan_rule_query_open_req", paramPlanRuleQueryOpenReq)
    return nil
}

func (r AlibabaAlscCrmRuleQuerytaglistRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
    return r.paramPlanRuleQueryOpenReq
}

