package alsc

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmOpenRuleGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.open.rule.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmOpenRuleGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
