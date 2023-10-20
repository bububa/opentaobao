package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmopenrulegetAPIRequest 查询规则 API请求
// alibaba.alsc.crm.open.rule.get
//
// 查询会员规则
type AlibabaalsccrmopenrulegetAPIRequest struct {
	model.Params
	// 入参
	_paramRuleOpenReq *RuleOpenReq
}

// NewAlibabaalsccrmopenrulegetRequest 初始化AlibabaalsccrmopenrulegetAPIRequest对象
func NewAlibabaalsccrmopenrulegetRequest() *AlibabaalsccrmopenrulegetAPIRequest {
	return &AlibabaalsccrmopenrulegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmopenrulegetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.open.rule.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmopenrulegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmopenrulegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamRuleOpenReq is ParamRuleOpenReq Setter
// 入参
func (r *AlibabaalsccrmopenrulegetAPIRequest) SetParamRuleOpenReq(_paramRuleOpenReq *RuleOpenReq) error {
	r._paramRuleOpenReq = _paramRuleOpenReq
	r.Set("param_rule_open_req", _paramRuleOpenReq)
	return nil
}

// GetParamRuleOpenReq ParamRuleOpenReq Getter
func (r AlibabaalsccrmopenrulegetAPIRequest) GetParamRuleOpenReq() *RuleOpenReq {
	return r._paramRuleOpenReq
}
