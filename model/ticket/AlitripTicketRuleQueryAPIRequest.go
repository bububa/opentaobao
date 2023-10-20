package ticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripTicketRuleQueryAPIRequest 【门票API2.0】门票规则信息查询接口 API请求
// alitrip.ticket.rule.query
//
// 门票规则信息查询接口：返回商家上传的门票规则信息
type AlitripTicketRuleQueryAPIRequest struct {
	model.Params
	// 卖家景点规则编码
	_outRuleId string
}

// NewAlitripTicketRuleQueryRequest 初始化AlitripTicketRuleQueryAPIRequest对象
func NewAlitripTicketRuleQueryRequest() *AlitripTicketRuleQueryAPIRequest {
	return &AlitripTicketRuleQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTicketRuleQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.ticket.rule.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTicketRuleQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTicketRuleQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutRuleId is OutRuleId Setter
// 卖家景点规则编码
func (r *AlitripTicketRuleQueryAPIRequest) SetOutRuleId(_outRuleId string) error {
	r._outRuleId = _outRuleId
	r.Set("out_rule_id", _outRuleId)
	return nil
}

// GetOutRuleId OutRuleId Getter
func (r AlitripTicketRuleQueryAPIRequest) GetOutRuleId() string {
	return r._outRuleId
}
