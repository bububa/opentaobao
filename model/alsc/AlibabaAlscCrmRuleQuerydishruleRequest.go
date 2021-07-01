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
type AlibabaAlscCrmRuleQuerydishruleAPIRequest struct {
    model.Params
    // 系统自动生成
    _paramPlanRuleQueryOpenReq   *PlanRuleQueryOpenReq
}

// 初始化AlibabaAlscCrmRuleQuerydishruleAPIRequest对象
func NewAlibabaAlscCrmRuleQuerydishruleRequest() *AlibabaAlscCrmRuleQuerydishruleAPIRequest{
    return &AlibabaAlscCrmRuleQuerydishruleAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRuleQuerydishruleAPIRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.rule.querydishrule"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRuleQuerydishruleAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamPlanRuleQueryOpenReq Setter
// 系统自动生成
func (r *AlibabaAlscCrmRuleQuerydishruleAPIRequest) SetParamPlanRuleQueryOpenReq(_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
    r._paramPlanRuleQueryOpenReq = _paramPlanRuleQueryOpenReq
    r.Set("param_plan_rule_query_open_req", _paramPlanRuleQueryOpenReq)
    return nil
}

// ParamPlanRuleQueryOpenReq Getter
func (r AlibabaAlscCrmRuleQuerydishruleAPIRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
    return r._paramPlanRuleQueryOpenReq
}
