package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmrulequeryjoinmruleAPIRequest 查询品牌下的成为会员规则 API请求
// alibaba.alsc.crm.rule.queryjoinmrule
//
// 查询品牌下的成为会员规则
type AlibabaalsccrmrulequeryjoinmruleAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq
}

// NewAlibabaalsccrmrulequeryjoinmruleRequest 初始化AlibabaalsccrmrulequeryjoinmruleAPIRequest对象
func NewAlibabaalsccrmrulequeryjoinmruleRequest() *AlibabaalsccrmrulequeryjoinmruleAPIRequest {
	return &AlibabaalsccrmrulequeryjoinmruleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmrulequeryjoinmruleAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.rule.queryjoinmrule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmrulequeryjoinmruleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmrulequeryjoinmruleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPlanRuleQueryOpenReq is ParamPlanRuleQueryOpenReq Setter
// 系统自动生成
func (r *AlibabaalsccrmrulequeryjoinmruleAPIRequest) SetParamPlanRuleQueryOpenReq(_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
	r._paramPlanRuleQueryOpenReq = _paramPlanRuleQueryOpenReq
	r.Set("param_plan_rule_query_open_req", _paramPlanRuleQueryOpenReq)
	return nil
}

// GetParamPlanRuleQueryOpenReq ParamPlanRuleQueryOpenReq Getter
func (r AlibabaalsccrmrulequeryjoinmruleAPIRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
	return r._paramPlanRuleQueryOpenReq
}
