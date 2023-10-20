package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRuleQuerympriceruleAPIRequest 查询品牌下的会员价规则 API请求
// alibaba.alsc.crm.rule.querympricerule
//
// 查询品牌下的会员价规则
type AlibabaAlscCrmRuleQuerympriceruleAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq
}

// NewAlibabaAlscCrmRuleQuerympriceruleRequest 初始化AlibabaAlscCrmRuleQuerympriceruleAPIRequest对象
func NewAlibabaAlscCrmRuleQuerympriceruleRequest() *AlibabaAlscCrmRuleQuerympriceruleAPIRequest {
	return &AlibabaAlscCrmRuleQuerympriceruleAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmRuleQuerympriceruleAPIRequest) Reset() {
	r._paramPlanRuleQueryOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRuleQuerympriceruleAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.rule.querympricerule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRuleQuerympriceruleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmRuleQuerympriceruleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPlanRuleQueryOpenReq is ParamPlanRuleQueryOpenReq Setter
// 系统自动生成
func (r *AlibabaAlscCrmRuleQuerympriceruleAPIRequest) SetParamPlanRuleQueryOpenReq(_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
	r._paramPlanRuleQueryOpenReq = _paramPlanRuleQueryOpenReq
	r.Set("param_plan_rule_query_open_req", _paramPlanRuleQueryOpenReq)
	return nil
}

// GetParamPlanRuleQueryOpenReq ParamPlanRuleQueryOpenReq Getter
func (r AlibabaAlscCrmRuleQuerympriceruleAPIRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
	return r._paramPlanRuleQueryOpenReq
}

var poolAlibabaAlscCrmRuleQuerympriceruleAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmRuleQuerympriceruleRequest()
	},
}

// GetAlibabaAlscCrmRuleQuerympriceruleRequest 从 sync.Pool 获取 AlibabaAlscCrmRuleQuerympriceruleAPIRequest
func GetAlibabaAlscCrmRuleQuerympriceruleAPIRequest() *AlibabaAlscCrmRuleQuerympriceruleAPIRequest {
	return poolAlibabaAlscCrmRuleQuerympriceruleAPIRequest.Get().(*AlibabaAlscCrmRuleQuerympriceruleAPIRequest)
}

// ReleaseAlibabaAlscCrmRuleQuerympriceruleAPIRequest 将 AlibabaAlscCrmRuleQuerympriceruleAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmRuleQuerympriceruleAPIRequest(v *AlibabaAlscCrmRuleQuerympriceruleAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmRuleQuerympriceruleAPIRequest.Put(v)
}
