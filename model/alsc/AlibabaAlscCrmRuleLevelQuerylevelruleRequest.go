package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询会员等级规则 API请求
alibaba.alsc.crm.rule.level.querylevelrule

查询会员等级规则
*/
type AlibabaAlscCrmRuleLevelQuerylevelruleRequest struct {
    model.Params
    // 请求参数
    planRuleQueryRequest   *PlanRuleQueryOpenReq
}

// 初始化AlibabaAlscCrmRuleLevelQuerylevelruleRequest对象
func NewAlibabaAlscCrmRuleLevelQuerylevelruleRequest() *AlibabaAlscCrmRuleLevelQuerylevelruleRequest{
    return &AlibabaAlscCrmRuleLevelQuerylevelruleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRuleLevelQuerylevelruleRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.rule.level.querylevelrule"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRuleLevelQuerylevelruleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PlanRuleQueryRequest Setter
// 请求参数
func (r *AlibabaAlscCrmRuleLevelQuerylevelruleRequest) SetPlanRuleQueryRequest(planRuleQueryRequest *PlanRuleQueryOpenReq) error {
    r.planRuleQueryRequest = planRuleQueryRequest
    r.Set("plan_rule_query_request", planRuleQueryRequest)
    return nil
}

// PlanRuleQueryRequest Getter
func (r AlibabaAlscCrmRuleLevelQuerylevelruleRequest) GetPlanRuleQueryRequest() *PlanRuleQueryOpenReq {
    return r.planRuleQueryRequest
}
