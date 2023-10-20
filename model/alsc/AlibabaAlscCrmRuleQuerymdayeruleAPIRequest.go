package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRuleQuerymdayeruleAPIRequest 查询品牌下的会员日规则 API请求
// alibaba.alsc.crm.rule.querymdayerule
//
// 查询品牌下的会员日规则
type AlibabaAlscCrmRuleQuerymdayeruleAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq
}

// NewAlibabaAlscCrmRuleQuerymdayeruleRequest 初始化AlibabaAlscCrmRuleQuerymdayeruleAPIRequest对象
func NewAlibabaAlscCrmRuleQuerymdayeruleRequest() *AlibabaAlscCrmRuleQuerymdayeruleAPIRequest {
	return &AlibabaAlscCrmRuleQuerymdayeruleAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmRuleQuerymdayeruleAPIRequest) Reset() {
	r._paramPlanRuleQueryOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRuleQuerymdayeruleAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.rule.querymdayerule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRuleQuerymdayeruleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmRuleQuerymdayeruleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPlanRuleQueryOpenReq is ParamPlanRuleQueryOpenReq Setter
// 系统自动生成
func (r *AlibabaAlscCrmRuleQuerymdayeruleAPIRequest) SetParamPlanRuleQueryOpenReq(_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
	r._paramPlanRuleQueryOpenReq = _paramPlanRuleQueryOpenReq
	r.Set("param_plan_rule_query_open_req", _paramPlanRuleQueryOpenReq)
	return nil
}

// GetParamPlanRuleQueryOpenReq ParamPlanRuleQueryOpenReq Getter
func (r AlibabaAlscCrmRuleQuerymdayeruleAPIRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
	return r._paramPlanRuleQueryOpenReq
}

var poolAlibabaAlscCrmRuleQuerymdayeruleAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmRuleQuerymdayeruleRequest()
	},
}

// GetAlibabaAlscCrmRuleQuerymdayeruleRequest 从 sync.Pool 获取 AlibabaAlscCrmRuleQuerymdayeruleAPIRequest
func GetAlibabaAlscCrmRuleQuerymdayeruleAPIRequest() *AlibabaAlscCrmRuleQuerymdayeruleAPIRequest {
	return poolAlibabaAlscCrmRuleQuerymdayeruleAPIRequest.Get().(*AlibabaAlscCrmRuleQuerymdayeruleAPIRequest)
}

// ReleaseAlibabaAlscCrmRuleQuerymdayeruleAPIRequest 将 AlibabaAlscCrmRuleQuerymdayeruleAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmRuleQuerymdayeruleAPIRequest(v *AlibabaAlscCrmRuleQuerymdayeruleAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmRuleQuerymdayeruleAPIRequest.Put(v)
}
