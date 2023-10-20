package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmpointrulegetAPIRequest 查询积分规则 API请求
// alibaba.alsc.crm.point.rule.get
//
// 新增积分规则查询接口,传入includeLogicalDelete和maxUpdateTime时走同步下行逻辑不然则走普通积分规则查询接口
type AlibabaalsccrmpointrulegetAPIRequest struct {
	model.Params
	// 入参
	_paramQueryPointRuleOpenReq *QueryPointRuleOpenReq
}

// NewAlibabaalsccrmpointrulegetRequest 初始化AlibabaalsccrmpointrulegetAPIRequest对象
func NewAlibabaalsccrmpointrulegetRequest() *AlibabaalsccrmpointrulegetAPIRequest {
	return &AlibabaalsccrmpointrulegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmpointrulegetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.point.rule.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmpointrulegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmpointrulegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQueryPointRuleOpenReq is ParamQueryPointRuleOpenReq Setter
// 入参
func (r *AlibabaalsccrmpointrulegetAPIRequest) SetParamQueryPointRuleOpenReq(_paramQueryPointRuleOpenReq *QueryPointRuleOpenReq) error {
	r._paramQueryPointRuleOpenReq = _paramQueryPointRuleOpenReq
	r.Set("param_query_point_rule_open_req", _paramQueryPointRuleOpenReq)
	return nil
}

// GetParamQueryPointRuleOpenReq ParamQueryPointRuleOpenReq Getter
func (r AlibabaalsccrmpointrulegetAPIRequest) GetParamQueryPointRuleOpenReq() *QueryPointRuleOpenReq {
	return r._paramQueryPointRuleOpenReq
}
