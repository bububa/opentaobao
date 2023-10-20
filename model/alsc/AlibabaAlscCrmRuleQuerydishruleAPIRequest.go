package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRuleQuerydishruleAPIRequest 查询品牌下的入会菜品规则 API请求
// alibaba.alsc.crm.rule.querydishrule
//
// 查询品牌下的入会菜品规则
type AlibabaAlscCrmRuleQuerydishruleAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq
}

// NewAlibabaAlscCrmRuleQuerydishruleRequest 初始化AlibabaAlscCrmRuleQuerydishruleAPIRequest对象
func NewAlibabaAlscCrmRuleQuerydishruleRequest() *AlibabaAlscCrmRuleQuerydishruleAPIRequest {
	return &AlibabaAlscCrmRuleQuerydishruleAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmRuleQuerydishruleAPIRequest) Reset() {
	r._paramPlanRuleQueryOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRuleQuerydishruleAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.rule.querydishrule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRuleQuerydishruleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmRuleQuerydishruleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPlanRuleQueryOpenReq is ParamPlanRuleQueryOpenReq Setter
// 系统自动生成
func (r *AlibabaAlscCrmRuleQuerydishruleAPIRequest) SetParamPlanRuleQueryOpenReq(_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
	r._paramPlanRuleQueryOpenReq = _paramPlanRuleQueryOpenReq
	r.Set("param_plan_rule_query_open_req", _paramPlanRuleQueryOpenReq)
	return nil
}

// GetParamPlanRuleQueryOpenReq ParamPlanRuleQueryOpenReq Getter
func (r AlibabaAlscCrmRuleQuerydishruleAPIRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
	return r._paramPlanRuleQueryOpenReq
}

var poolAlibabaAlscCrmRuleQuerydishruleAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmRuleQuerydishruleRequest()
	},
}

// GetAlibabaAlscCrmRuleQuerydishruleRequest 从 sync.Pool 获取 AlibabaAlscCrmRuleQuerydishruleAPIRequest
func GetAlibabaAlscCrmRuleQuerydishruleAPIRequest() *AlibabaAlscCrmRuleQuerydishruleAPIRequest {
	return poolAlibabaAlscCrmRuleQuerydishruleAPIRequest.Get().(*AlibabaAlscCrmRuleQuerydishruleAPIRequest)
}

// ReleaseAlibabaAlscCrmRuleQuerydishruleAPIRequest 将 AlibabaAlscCrmRuleQuerydishruleAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmRuleQuerydishruleAPIRequest(v *AlibabaAlscCrmRuleQuerydishruleAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmRuleQuerydishruleAPIRequest.Put(v)
}
