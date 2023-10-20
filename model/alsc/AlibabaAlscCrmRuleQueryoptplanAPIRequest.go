package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRuleQueryoptplanAPIRequest 查询运营计划 API请求
// alibaba.alsc.crm.rule.queryoptplan
//
// 查询运营计划
type AlibabaAlscCrmRuleQueryoptplanAPIRequest struct {
	model.Params
	// 请求参数
	_planRuleQueryOpenRequest *PlanRuleQueryOpenReq
}

// NewAlibabaAlscCrmRuleQueryoptplanRequest 初始化AlibabaAlscCrmRuleQueryoptplanAPIRequest对象
func NewAlibabaAlscCrmRuleQueryoptplanRequest() *AlibabaAlscCrmRuleQueryoptplanAPIRequest {
	return &AlibabaAlscCrmRuleQueryoptplanAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmRuleQueryoptplanAPIRequest) Reset() {
	r._planRuleQueryOpenRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRuleQueryoptplanAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.rule.queryoptplan"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRuleQueryoptplanAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmRuleQueryoptplanAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPlanRuleQueryOpenRequest is PlanRuleQueryOpenRequest Setter
// 请求参数
func (r *AlibabaAlscCrmRuleQueryoptplanAPIRequest) SetPlanRuleQueryOpenRequest(_planRuleQueryOpenRequest *PlanRuleQueryOpenReq) error {
	r._planRuleQueryOpenRequest = _planRuleQueryOpenRequest
	r.Set("plan_rule_query_open_request", _planRuleQueryOpenRequest)
	return nil
}

// GetPlanRuleQueryOpenRequest PlanRuleQueryOpenRequest Getter
func (r AlibabaAlscCrmRuleQueryoptplanAPIRequest) GetPlanRuleQueryOpenRequest() *PlanRuleQueryOpenReq {
	return r._planRuleQueryOpenRequest
}

var poolAlibabaAlscCrmRuleQueryoptplanAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmRuleQueryoptplanRequest()
	},
}

// GetAlibabaAlscCrmRuleQueryoptplanRequest 从 sync.Pool 获取 AlibabaAlscCrmRuleQueryoptplanAPIRequest
func GetAlibabaAlscCrmRuleQueryoptplanAPIRequest() *AlibabaAlscCrmRuleQueryoptplanAPIRequest {
	return poolAlibabaAlscCrmRuleQueryoptplanAPIRequest.Get().(*AlibabaAlscCrmRuleQueryoptplanAPIRequest)
}

// ReleaseAlibabaAlscCrmRuleQueryoptplanAPIRequest 将 AlibabaAlscCrmRuleQueryoptplanAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmRuleQueryoptplanAPIRequest(v *AlibabaAlscCrmRuleQueryoptplanAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmRuleQueryoptplanAPIRequest.Put(v)
}
