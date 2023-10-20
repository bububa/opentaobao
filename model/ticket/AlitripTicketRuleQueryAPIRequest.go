package ticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripticketrulequeryAPIRequest 【门票API2.0】门票规则信息查询接口 API请求
// alitrip.ticket.rule.query
//
// 门票规则信息查询接口：返回商家上传的门票规则信息
type AlitripticketrulequeryAPIRequest struct {
	model.Params
	// 卖家景点规则编码
	_outRuleId string
}

// NewAlitripticketrulequeryRequest 初始化AlitripticketrulequeryAPIRequest对象
func NewAlitripticketrulequeryRequest() *AlitripticketrulequeryAPIRequest {
	return &AlitripticketrulequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripticketrulequeryAPIRequest) GetApiMethodName() string {
	return "alitrip.ticket.rule.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripticketrulequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripticketrulequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutRuleId is OutRuleId Setter
// 卖家景点规则编码
func (r *AlitripticketrulequeryAPIRequest) SetOutRuleId(_outRuleId string) error {
	r._outRuleId = _outRuleId
	r.Set("out_rule_id", _outRuleId)
	return nil
}

// GetOutRuleId OutRuleId Getter
func (r AlitripticketrulequeryAPIRequest) GetOutRuleId() string {
	return r._outRuleId
}
