package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询标签列表 API请求
alibaba.alsc.crm.rule.querytaglist

查询标签列表
*/
type AlibabaAlscCrmRuleQuerytaglistAPIRequest struct {
    model.Params
    // 请求参数
    _paramPlanRuleQueryOpenReq   *PlanRuleQueryOpenReq
}

// 初始化AlibabaAlscCrmRuleQuerytaglistAPIRequest对象
func NewAlibabaAlscCrmRuleQuerytaglistRequest() *AlibabaAlscCrmRuleQuerytaglistAPIRequest{
    return &AlibabaAlscCrmRuleQuerytaglistAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRuleQuerytaglistAPIRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.rule.querytaglist"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRuleQuerytaglistAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamPlanRuleQueryOpenReq Setter
// 请求参数
func (r *AlibabaAlscCrmRuleQuerytaglistAPIRequest) SetParamPlanRuleQueryOpenReq(_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
    r._paramPlanRuleQueryOpenReq = _paramPlanRuleQueryOpenReq
    r.Set("param_plan_rule_query_open_req", _paramPlanRuleQueryOpenReq)
    return nil
}

// ParamPlanRuleQueryOpenReq Getter
func (r AlibabaAlscCrmRuleQuerytaglistAPIRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
    return r._paramPlanRuleQueryOpenReq
}
