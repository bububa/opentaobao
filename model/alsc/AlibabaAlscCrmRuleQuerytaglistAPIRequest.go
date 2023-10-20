package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmrulequerytaglistAPIRequest 查询标签列表 API请求
// alibaba.alsc.crm.rule.querytaglist
//
// 查询标签列表
type AlibabaalsccrmrulequerytaglistAPIRequest struct {
	model.Params
	// 请求参数
	_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq
}

// NewAlibabaalsccrmrulequerytaglistRequest 初始化AlibabaalsccrmrulequerytaglistAPIRequest对象
func NewAlibabaalsccrmrulequerytaglistRequest() *AlibabaalsccrmrulequerytaglistAPIRequest {
	return &AlibabaalsccrmrulequerytaglistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmrulequerytaglistAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.rule.querytaglist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmrulequerytaglistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmrulequerytaglistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPlanRuleQueryOpenReq is ParamPlanRuleQueryOpenReq Setter
// 请求参数
func (r *AlibabaalsccrmrulequerytaglistAPIRequest) SetParamPlanRuleQueryOpenReq(_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
	r._paramPlanRuleQueryOpenReq = _paramPlanRuleQueryOpenReq
	r.Set("param_plan_rule_query_open_req", _paramPlanRuleQueryOpenReq)
	return nil
}

// GetParamPlanRuleQueryOpenReq ParamPlanRuleQueryOpenReq Getter
func (r AlibabaalsccrmrulequerytaglistAPIRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
	return r._paramPlanRuleQueryOpenReq
}
