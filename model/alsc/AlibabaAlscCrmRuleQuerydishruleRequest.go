package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询品牌下的入会菜品规则 API请求
alibaba.alsc.crm.rule.querydishrule

查询品牌下的入会菜品规则
*/
type AlibabaAlscCrmRuleQuerydishruleRequest struct {
    model.Params
    // 系统自动生成
    paramPlanRuleQueryOpenReq   *PlanRuleQueryOpenReq
}

// 初始化AlibabaAlscCrmRuleQuerydishruleRequest对象
func NewAlibabaAlscCrmRuleQuerydishruleRequest() *AlibabaAlscCrmRuleQuerydishruleRequest{
    return &AlibabaAlscCrmRuleQuerydishruleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRuleQuerydishruleRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.rule.querydishrule"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRuleQuerydishruleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamPlanRuleQueryOpenReq Setter
// 系统自动生成
func (r *AlibabaAlscCrmRuleQuerydishruleRequest) SetParamPlanRuleQueryOpenReq(paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
    r.paramPlanRuleQueryOpenReq = paramPlanRuleQueryOpenReq
    r.Set("param_plan_rule_query_open_req", paramPlanRuleQueryOpenReq)
    return nil
}

// ParamPlanRuleQueryOpenReq Getter
func (r AlibabaAlscCrmRuleQuerydishruleRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
    return r.paramPlanRuleQueryOpenReq
}
