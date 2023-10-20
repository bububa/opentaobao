package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRuleQuerygrowruleAPIRequest 查询品牌下的会员成长规则 API请求
// alibaba.alsc.crm.rule.querygrowrule
//
// 查询品牌下的会员成长规则
type AlibabaAlscCrmRuleQuerygrowruleAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq
}

// NewAlibabaAlscCrmRuleQuerygrowruleRequest 初始化AlibabaAlscCrmRuleQuerygrowruleAPIRequest对象
func NewAlibabaAlscCrmRuleQuerygrowruleRequest() *AlibabaAlscCrmRuleQuerygrowruleAPIRequest {
	return &AlibabaAlscCrmRuleQuerygrowruleAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmRuleQuerygrowruleAPIRequest) Reset() {
	r._paramPlanRuleQueryOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRuleQuerygrowruleAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.rule.querygrowrule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRuleQuerygrowruleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmRuleQuerygrowruleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPlanRuleQueryOpenReq is ParamPlanRuleQueryOpenReq Setter
// 系统自动生成
func (r *AlibabaAlscCrmRuleQuerygrowruleAPIRequest) SetParamPlanRuleQueryOpenReq(_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
	r._paramPlanRuleQueryOpenReq = _paramPlanRuleQueryOpenReq
	r.Set("param_plan_rule_query_open_req", _paramPlanRuleQueryOpenReq)
	return nil
}

// GetParamPlanRuleQueryOpenReq ParamPlanRuleQueryOpenReq Getter
func (r AlibabaAlscCrmRuleQuerygrowruleAPIRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
	return r._paramPlanRuleQueryOpenReq
}

var poolAlibabaAlscCrmRuleQuerygrowruleAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmRuleQuerygrowruleRequest()
	},
}

// GetAlibabaAlscCrmRuleQuerygrowruleRequest 从 sync.Pool 获取 AlibabaAlscCrmRuleQuerygrowruleAPIRequest
func GetAlibabaAlscCrmRuleQuerygrowruleAPIRequest() *AlibabaAlscCrmRuleQuerygrowruleAPIRequest {
	return poolAlibabaAlscCrmRuleQuerygrowruleAPIRequest.Get().(*AlibabaAlscCrmRuleQuerygrowruleAPIRequest)
}

// ReleaseAlibabaAlscCrmRuleQuerygrowruleAPIRequest 将 AlibabaAlscCrmRuleQuerygrowruleAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmRuleQuerygrowruleAPIRequest(v *AlibabaAlscCrmRuleQuerygrowruleAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmRuleQuerygrowruleAPIRequest.Put(v)
}
