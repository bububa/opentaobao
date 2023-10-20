package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRuleLevelQuerylevelruleAPIRequest 查询会员等级规则 API请求
// alibaba.alsc.crm.rule.level.querylevelrule
//
// 查询会员等级规则
type AlibabaAlscCrmRuleLevelQuerylevelruleAPIRequest struct {
	model.Params
	// 请求参数
	_planRuleQueryRequest *PlanRuleQueryOpenReq
}

// NewAlibabaAlscCrmRuleLevelQuerylevelruleRequest 初始化AlibabaAlscCrmRuleLevelQuerylevelruleAPIRequest对象
func NewAlibabaAlscCrmRuleLevelQuerylevelruleRequest() *AlibabaAlscCrmRuleLevelQuerylevelruleAPIRequest {
	return &AlibabaAlscCrmRuleLevelQuerylevelruleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRuleLevelQuerylevelruleAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.rule.level.querylevelrule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRuleLevelQuerylevelruleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmRuleLevelQuerylevelruleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPlanRuleQueryRequest is PlanRuleQueryRequest Setter
// 请求参数
func (r *AlibabaAlscCrmRuleLevelQuerylevelruleAPIRequest) SetPlanRuleQueryRequest(_planRuleQueryRequest *PlanRuleQueryOpenReq) error {
	r._planRuleQueryRequest = _planRuleQueryRequest
	r.Set("plan_rule_query_request", _planRuleQueryRequest)
	return nil
}

// GetPlanRuleQueryRequest PlanRuleQueryRequest Getter
func (r AlibabaAlscCrmRuleLevelQuerylevelruleAPIRequest) GetPlanRuleQueryRequest() *PlanRuleQueryOpenReq {
	return r._planRuleQueryRequest
}
