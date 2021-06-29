package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询品牌下的成为会员规则 API请求
alibaba.alsc.crm.rule.queryjoinmrule

查询品牌下的成为会员规则
*/
type AlibabaAlscCrmRuleQueryjoinmruleRequest struct {
    model.Params
    // 系统自动生成
    paramPlanRuleQueryOpenReq   *PlanRuleQueryOpenReq
}

// 初始化AlibabaAlscCrmRuleQueryjoinmruleRequest对象
func NewAlibabaAlscCrmRuleQueryjoinmruleRequest() *AlibabaAlscCrmRuleQueryjoinmruleRequest{
    return &AlibabaAlscCrmRuleQueryjoinmruleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRuleQueryjoinmruleRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.rule.queryjoinmrule"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRuleQueryjoinmruleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamPlanRuleQueryOpenReq Setter
// 系统自动生成
func (r *AlibabaAlscCrmRuleQueryjoinmruleRequest) SetParamPlanRuleQueryOpenReq(paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
    r.paramPlanRuleQueryOpenReq = paramPlanRuleQueryOpenReq
    r.Set("param_plan_rule_query_open_req", paramPlanRuleQueryOpenReq)
    return nil
}

// ParamPlanRuleQueryOpenReq Getter
func (r AlibabaAlscCrmRuleQueryjoinmruleRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
    return r.paramPlanRuleQueryOpenReq
}
