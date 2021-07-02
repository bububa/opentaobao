package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmPointRuleGetAPIRequest 查询积分规则 API请求
// alibaba.alsc.crm.point.rule.get
//
// 新增积分规则查询接口,传入includeLogicalDelete和maxUpdateTime时走同步下行逻辑不然则走普通积分规则查询接口
type AlibabaAlscCrmPointRuleGetAPIRequest struct {
	model.Params
	// 入参
	_paramQueryPointRuleOpenReq *QueryPointRuleOpenReq
}

// NewAlibabaAlscCrmPointRuleGetRequest 初始化AlibabaAlscCrmPointRuleGetAPIRequest对象
func NewAlibabaAlscCrmPointRuleGetRequest() *AlibabaAlscCrmPointRuleGetAPIRequest {
	return &AlibabaAlscCrmPointRuleGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmPointRuleGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.point.rule.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmPointRuleGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamQueryPointRuleOpenReq Setter
// 入参
func (r *AlibabaAlscCrmPointRuleGetAPIRequest) SetParamQueryPointRuleOpenReq(_paramQueryPointRuleOpenReq *QueryPointRuleOpenReq) error {
	r._paramQueryPointRuleOpenReq = _paramQueryPointRuleOpenReq
	r.Set("param_query_point_rule_open_req", _paramQueryPointRuleOpenReq)
	return nil
}

// Get ParamQueryPointRuleOpenReq Getter
func (r AlibabaAlscCrmPointRuleGetAPIRequest) GetParamQueryPointRuleOpenReq() *QueryPointRuleOpenReq {
	return r._paramQueryPointRuleOpenReq
}
