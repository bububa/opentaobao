package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmrulequeryoptplanAPIRequest 查询运营计划 API请求
// alibaba.alsc.crm.rule.queryoptplan
//
// 查询运营计划
type AlibabaalsccrmrulequeryoptplanAPIRequest struct {
	model.Params
	// 请求参数
	_planRuleQueryOpenRequest *PlanRuleQueryOpenReq
}

// NewAlibabaalsccrmrulequeryoptplanRequest 初始化AlibabaalsccrmrulequeryoptplanAPIRequest对象
func NewAlibabaalsccrmrulequeryoptplanRequest() *AlibabaalsccrmrulequeryoptplanAPIRequest {
	return &AlibabaalsccrmrulequeryoptplanAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmrulequeryoptplanAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.rule.queryoptplan"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmrulequeryoptplanAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmrulequeryoptplanAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPlanRuleQueryOpenRequest is PlanRuleQueryOpenRequest Setter
// 请求参数
func (r *AlibabaalsccrmrulequeryoptplanAPIRequest) SetPlanRuleQueryOpenRequest(_planRuleQueryOpenRequest *PlanRuleQueryOpenReq) error {
	r._planRuleQueryOpenRequest = _planRuleQueryOpenRequest
	r.Set("plan_rule_query_open_request", _planRuleQueryOpenRequest)
	return nil
}

// GetPlanRuleQueryOpenRequest PlanRuleQueryOpenRequest Getter
func (r AlibabaalsccrmrulequeryoptplanAPIRequest) GetPlanRuleQueryOpenRequest() *PlanRuleQueryOpenReq {
	return r._planRuleQueryOpenRequest
}
