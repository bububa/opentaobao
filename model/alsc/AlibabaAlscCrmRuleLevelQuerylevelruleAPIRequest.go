package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmrulelevelquerylevelruleAPIRequest 查询会员等级规则 API请求
// alibaba.alsc.crm.rule.level.querylevelrule
//
// 查询会员等级规则
type AlibabaalsccrmrulelevelquerylevelruleAPIRequest struct {
	model.Params
	// 请求参数
	_planRuleQueryRequest *PlanRuleQueryOpenReq
}

// NewAlibabaalsccrmrulelevelquerylevelruleRequest 初始化AlibabaalsccrmrulelevelquerylevelruleAPIRequest对象
func NewAlibabaalsccrmrulelevelquerylevelruleRequest() *AlibabaalsccrmrulelevelquerylevelruleAPIRequest {
	return &AlibabaalsccrmrulelevelquerylevelruleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmrulelevelquerylevelruleAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.rule.level.querylevelrule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmrulelevelquerylevelruleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmrulelevelquerylevelruleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPlanRuleQueryRequest is PlanRuleQueryRequest Setter
// 请求参数
func (r *AlibabaalsccrmrulelevelquerylevelruleAPIRequest) SetPlanRuleQueryRequest(_planRuleQueryRequest *PlanRuleQueryOpenReq) error {
	r._planRuleQueryRequest = _planRuleQueryRequest
	r.Set("plan_rule_query_request", _planRuleQueryRequest)
	return nil
}

// GetPlanRuleQueryRequest PlanRuleQueryRequest Getter
func (r AlibabaalsccrmrulelevelquerylevelruleAPIRequest) GetPlanRuleQueryRequest() *PlanRuleQueryOpenReq {
	return r._planRuleQueryRequest
}
