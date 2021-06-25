package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询会员等级规则 APIRequest
alibaba.alsc.crm.rule.level.querylevelrule

查询会员等级规则
*/
type AlibabaAlscCrmRuleLevelQuerylevelruleRequest struct {
    model.Params

    // 请求参数
    planRuleQueryRequest   *PlanRuleQueryOpenReq 

}

func NewAlibabaAlscCrmRuleLevelQuerylevelruleRequest() *AlibabaAlscCrmRuleLevelQuerylevelruleRequest{
    return &AlibabaAlscCrmRuleLevelQuerylevelruleRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmRuleLevelQuerylevelruleRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.rule.level.querylevelrule"
}

func (r AlibabaAlscCrmRuleLevelQuerylevelruleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmRuleLevelQuerylevelruleRequest) SetPlanRuleQueryRequest(planRuleQueryRequest *PlanRuleQueryOpenReq) error {
    r.planRuleQueryRequest = planRuleQueryRequest
    r.Set("plan_rule_query_request", planRuleQueryRequest)
    return nil
}

func (r AlibabaAlscCrmRuleLevelQuerylevelruleRequest) GetPlanRuleQueryRequest() *PlanRuleQueryOpenReq {
    return r.planRuleQueryRequest
}

