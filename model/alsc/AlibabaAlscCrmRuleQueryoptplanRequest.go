package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询运营计划 API请求
alibaba.alsc.crm.rule.queryoptplan

查询运营计划
*/
type AlibabaAlscCrmRuleQueryoptplanAPIRequest struct {
    model.Params
    // 请求参数
    _planRuleQueryOpenRequest   *PlanRuleQueryOpenReq
}

// 初始化AlibabaAlscCrmRuleQueryoptplanAPIRequest对象
func NewAlibabaAlscCrmRuleQueryoptplanRequest() *AlibabaAlscCrmRuleQueryoptplanAPIRequest{
    return &AlibabaAlscCrmRuleQueryoptplanAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRuleQueryoptplanAPIRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.rule.queryoptplan"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRuleQueryoptplanAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PlanRuleQueryOpenRequest Setter
// 请求参数
func (r *AlibabaAlscCrmRuleQueryoptplanAPIRequest) SetPlanRuleQueryOpenRequest(_planRuleQueryOpenRequest *PlanRuleQueryOpenReq) error {
    r._planRuleQueryOpenRequest = _planRuleQueryOpenRequest
    r.Set("plan_rule_query_open_request", _planRuleQueryOpenRequest)
    return nil
}

// PlanRuleQueryOpenRequest Getter
func (r AlibabaAlscCrmRuleQueryoptplanAPIRequest) GetPlanRuleQueryOpenRequest() *PlanRuleQueryOpenReq {
    return r._planRuleQueryOpenRequest
}
