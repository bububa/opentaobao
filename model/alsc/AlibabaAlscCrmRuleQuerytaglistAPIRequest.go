package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRuleQuerytaglistAPIRequest 查询标签列表 API请求
// alibaba.alsc.crm.rule.querytaglist
//
// 查询标签列表
type AlibabaAlscCrmRuleQuerytaglistAPIRequest struct {
	model.Params
	// 请求参数
	_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq
}

// NewAlibabaAlscCrmRuleQuerytaglistRequest 初始化AlibabaAlscCrmRuleQuerytaglistAPIRequest对象
func NewAlibabaAlscCrmRuleQuerytaglistRequest() *AlibabaAlscCrmRuleQuerytaglistAPIRequest {
	return &AlibabaAlscCrmRuleQuerytaglistAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmRuleQuerytaglistAPIRequest) Reset() {
	r._paramPlanRuleQueryOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRuleQuerytaglistAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.rule.querytaglist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRuleQuerytaglistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmRuleQuerytaglistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPlanRuleQueryOpenReq is ParamPlanRuleQueryOpenReq Setter
// 请求参数
func (r *AlibabaAlscCrmRuleQuerytaglistAPIRequest) SetParamPlanRuleQueryOpenReq(_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
	r._paramPlanRuleQueryOpenReq = _paramPlanRuleQueryOpenReq
	r.Set("param_plan_rule_query_open_req", _paramPlanRuleQueryOpenReq)
	return nil
}

// GetParamPlanRuleQueryOpenReq ParamPlanRuleQueryOpenReq Getter
func (r AlibabaAlscCrmRuleQuerytaglistAPIRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
	return r._paramPlanRuleQueryOpenReq
}

var poolAlibabaAlscCrmRuleQuerytaglistAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmRuleQuerytaglistRequest()
	},
}

// GetAlibabaAlscCrmRuleQuerytaglistRequest 从 sync.Pool 获取 AlibabaAlscCrmRuleQuerytaglistAPIRequest
func GetAlibabaAlscCrmRuleQuerytaglistAPIRequest() *AlibabaAlscCrmRuleQuerytaglistAPIRequest {
	return poolAlibabaAlscCrmRuleQuerytaglistAPIRequest.Get().(*AlibabaAlscCrmRuleQuerytaglistAPIRequest)
}

// ReleaseAlibabaAlscCrmRuleQuerytaglistAPIRequest 将 AlibabaAlscCrmRuleQuerytaglistAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmRuleQuerytaglistAPIRequest(v *AlibabaAlscCrmRuleQuerytaglistAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmRuleQuerytaglistAPIRequest.Put(v)
}
