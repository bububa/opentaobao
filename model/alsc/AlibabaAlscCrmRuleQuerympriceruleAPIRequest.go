package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmrulequerympriceruleAPIRequest 查询品牌下的会员价规则 API请求
// alibaba.alsc.crm.rule.querympricerule
//
// 查询品牌下的会员价规则
type AlibabaalsccrmrulequerympriceruleAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq
}

// NewAlibabaalsccrmrulequerympriceruleRequest 初始化AlibabaalsccrmrulequerympriceruleAPIRequest对象
func NewAlibabaalsccrmrulequerympriceruleRequest() *AlibabaalsccrmrulequerympriceruleAPIRequest {
	return &AlibabaalsccrmrulequerympriceruleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmrulequerympriceruleAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.rule.querympricerule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmrulequerympriceruleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmrulequerympriceruleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPlanRuleQueryOpenReq is ParamPlanRuleQueryOpenReq Setter
// 系统自动生成
func (r *AlibabaalsccrmrulequerympriceruleAPIRequest) SetParamPlanRuleQueryOpenReq(_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
	r._paramPlanRuleQueryOpenReq = _paramPlanRuleQueryOpenReq
	r.Set("param_plan_rule_query_open_req", _paramPlanRuleQueryOpenReq)
	return nil
}

// GetParamPlanRuleQueryOpenReq ParamPlanRuleQueryOpenReq Getter
func (r AlibabaalsccrmrulequerympriceruleAPIRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
	return r._paramPlanRuleQueryOpenReq
}
