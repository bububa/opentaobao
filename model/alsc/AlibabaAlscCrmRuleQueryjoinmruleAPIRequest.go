package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmRuleQueryjoinmruleAPIRequest
查询品牌下的成为会员规则 API请求
alibaba.alsc.crm.rule.queryjoinmrule

查询品牌下的成为会员规则 */
type AlibabaAlscCrmRuleQueryjoinmruleAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq
}

// NewAlibabaAlscCrmRuleQueryjoinmruleRequest 初始化AlibabaAlscCrmRuleQueryjoinmruleAPIRequest对象
func NewAlibabaAlscCrmRuleQueryjoinmruleRequest() *AlibabaAlscCrmRuleQueryjoinmruleAPIRequest {
	return &AlibabaAlscCrmRuleQueryjoinmruleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRuleQueryjoinmruleAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.rule.queryjoinmrule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRuleQueryjoinmruleAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamPlanRuleQueryOpenReq Setter
// 系统自动生成
func (r *AlibabaAlscCrmRuleQueryjoinmruleAPIRequest) SetParamPlanRuleQueryOpenReq(_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
	r._paramPlanRuleQueryOpenReq = _paramPlanRuleQueryOpenReq
	r.Set("param_plan_rule_query_open_req", _paramPlanRuleQueryOpenReq)
	return nil
}

// Get ParamPlanRuleQueryOpenReq Getter
func (r AlibabaAlscCrmRuleQueryjoinmruleAPIRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
	return r._paramPlanRuleQueryOpenReq
}
