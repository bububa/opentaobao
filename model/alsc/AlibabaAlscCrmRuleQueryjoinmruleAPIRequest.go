package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRuleQueryjoinmruleAPIRequest 查询品牌下的成为会员规则 API请求
// alibaba.alsc.crm.rule.queryjoinmrule
//
// 查询品牌下的成为会员规则
type AlibabaAlscCrmRuleQueryjoinmruleAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq
}

// NewAlibabaAlscCrmRuleQueryjoinmruleRequest 初始化AlibabaAlscCrmRuleQueryjoinmruleAPIRequest对象
func NewAlibabaAlscCrmRuleQueryjoinmruleRequest() *AlibabaAlscCrmRuleQueryjoinmruleAPIRequest {
	return &AlibabaAlscCrmRuleQueryjoinmruleAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmRuleQueryjoinmruleAPIRequest) Reset() {
	r._paramPlanRuleQueryOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRuleQueryjoinmruleAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.rule.queryjoinmrule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRuleQueryjoinmruleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmRuleQueryjoinmruleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPlanRuleQueryOpenReq is ParamPlanRuleQueryOpenReq Setter
// 系统自动生成
func (r *AlibabaAlscCrmRuleQueryjoinmruleAPIRequest) SetParamPlanRuleQueryOpenReq(_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
	r._paramPlanRuleQueryOpenReq = _paramPlanRuleQueryOpenReq
	r.Set("param_plan_rule_query_open_req", _paramPlanRuleQueryOpenReq)
	return nil
}

// GetParamPlanRuleQueryOpenReq ParamPlanRuleQueryOpenReq Getter
func (r AlibabaAlscCrmRuleQueryjoinmruleAPIRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
	return r._paramPlanRuleQueryOpenReq
}

var poolAlibabaAlscCrmRuleQueryjoinmruleAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmRuleQueryjoinmruleRequest()
	},
}

// GetAlibabaAlscCrmRuleQueryjoinmruleRequest 从 sync.Pool 获取 AlibabaAlscCrmRuleQueryjoinmruleAPIRequest
func GetAlibabaAlscCrmRuleQueryjoinmruleAPIRequest() *AlibabaAlscCrmRuleQueryjoinmruleAPIRequest {
	return poolAlibabaAlscCrmRuleQueryjoinmruleAPIRequest.Get().(*AlibabaAlscCrmRuleQueryjoinmruleAPIRequest)
}

// ReleaseAlibabaAlscCrmRuleQueryjoinmruleAPIRequest 将 AlibabaAlscCrmRuleQueryjoinmruleAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmRuleQueryjoinmruleAPIRequest(v *AlibabaAlscCrmRuleQueryjoinmruleAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmRuleQueryjoinmruleAPIRequest.Put(v)
}
