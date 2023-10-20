package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmOpenRuleGetAPIRequest 查询规则 API请求
// alibaba.alsc.crm.open.rule.get
//
// 查询会员规则
type AlibabaAlscCrmOpenRuleGetAPIRequest struct {
	model.Params
	// 入参
	_paramRuleOpenReq *RuleOpenReq
}

// NewAlibabaAlscCrmOpenRuleGetRequest 初始化AlibabaAlscCrmOpenRuleGetAPIRequest对象
func NewAlibabaAlscCrmOpenRuleGetRequest() *AlibabaAlscCrmOpenRuleGetAPIRequest {
	return &AlibabaAlscCrmOpenRuleGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmOpenRuleGetAPIRequest) Reset() {
	r._paramRuleOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmOpenRuleGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.open.rule.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmOpenRuleGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmOpenRuleGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamRuleOpenReq is ParamRuleOpenReq Setter
// 入参
func (r *AlibabaAlscCrmOpenRuleGetAPIRequest) SetParamRuleOpenReq(_paramRuleOpenReq *RuleOpenReq) error {
	r._paramRuleOpenReq = _paramRuleOpenReq
	r.Set("param_rule_open_req", _paramRuleOpenReq)
	return nil
}

// GetParamRuleOpenReq ParamRuleOpenReq Getter
func (r AlibabaAlscCrmOpenRuleGetAPIRequest) GetParamRuleOpenReq() *RuleOpenReq {
	return r._paramRuleOpenReq
}

var poolAlibabaAlscCrmOpenRuleGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmOpenRuleGetRequest()
	},
}

// GetAlibabaAlscCrmOpenRuleGetRequest 从 sync.Pool 获取 AlibabaAlscCrmOpenRuleGetAPIRequest
func GetAlibabaAlscCrmOpenRuleGetAPIRequest() *AlibabaAlscCrmOpenRuleGetAPIRequest {
	return poolAlibabaAlscCrmOpenRuleGetAPIRequest.Get().(*AlibabaAlscCrmOpenRuleGetAPIRequest)
}

// ReleaseAlibabaAlscCrmOpenRuleGetAPIRequest 将 AlibabaAlscCrmOpenRuleGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmOpenRuleGetAPIRequest(v *AlibabaAlscCrmOpenRuleGetAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmOpenRuleGetAPIRequest.Put(v)
}
