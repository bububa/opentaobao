package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmrulequerymdayeruleAPIRequest 查询品牌下的会员日规则 API请求
// alibaba.alsc.crm.rule.querymdayerule
//
// 查询品牌下的会员日规则
type AlibabaalsccrmrulequerymdayeruleAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq
}

// NewAlibabaalsccrmrulequerymdayeruleRequest 初始化AlibabaalsccrmrulequerymdayeruleAPIRequest对象
func NewAlibabaalsccrmrulequerymdayeruleRequest() *AlibabaalsccrmrulequerymdayeruleAPIRequest {
	return &AlibabaalsccrmrulequerymdayeruleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmrulequerymdayeruleAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.rule.querymdayerule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmrulequerymdayeruleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmrulequerymdayeruleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPlanRuleQueryOpenReq is ParamPlanRuleQueryOpenReq Setter
// 系统自动生成
func (r *AlibabaalsccrmrulequerymdayeruleAPIRequest) SetParamPlanRuleQueryOpenReq(_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
	r._paramPlanRuleQueryOpenReq = _paramPlanRuleQueryOpenReq
	r.Set("param_plan_rule_query_open_req", _paramPlanRuleQueryOpenReq)
	return nil
}

// GetParamPlanRuleQueryOpenReq ParamPlanRuleQueryOpenReq Getter
func (r AlibabaalsccrmrulequerymdayeruleAPIRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
	return r._paramPlanRuleQueryOpenReq
}
