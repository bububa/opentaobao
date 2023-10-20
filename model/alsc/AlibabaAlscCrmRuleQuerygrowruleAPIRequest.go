package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmrulequerygrowruleAPIRequest 查询品牌下的会员成长规则 API请求
// alibaba.alsc.crm.rule.querygrowrule
//
// 查询品牌下的会员成长规则
type AlibabaalsccrmrulequerygrowruleAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq
}

// NewAlibabaalsccrmrulequerygrowruleRequest 初始化AlibabaalsccrmrulequerygrowruleAPIRequest对象
func NewAlibabaalsccrmrulequerygrowruleRequest() *AlibabaalsccrmrulequerygrowruleAPIRequest {
	return &AlibabaalsccrmrulequerygrowruleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmrulequerygrowruleAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.rule.querygrowrule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmrulequerygrowruleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmrulequerygrowruleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPlanRuleQueryOpenReq is ParamPlanRuleQueryOpenReq Setter
// 系统自动生成
func (r *AlibabaalsccrmrulequerygrowruleAPIRequest) SetParamPlanRuleQueryOpenReq(_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
	r._paramPlanRuleQueryOpenReq = _paramPlanRuleQueryOpenReq
	r.Set("param_plan_rule_query_open_req", _paramPlanRuleQueryOpenReq)
	return nil
}

// GetParamPlanRuleQueryOpenReq ParamPlanRuleQueryOpenReq Getter
func (r AlibabaalsccrmrulequerygrowruleAPIRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
	return r._paramPlanRuleQueryOpenReq
}
