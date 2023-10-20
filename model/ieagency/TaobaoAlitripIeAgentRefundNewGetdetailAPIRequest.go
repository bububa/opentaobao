package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentRefundNewGetdetailAPIRequest 查询申请单详情(新版) API请求
// taobao.alitrip.ie.agent.refund.new.getdetail
//
// 查询申请单详情
type TaobaoAlitripIeAgentRefundNewGetdetailAPIRequest struct {
	model.Params
	// 请求
	_paramRefundOrderQueryDetailRq *RefundOrderQueryDetailRq
}

// NewTaobaoAlitripIeAgentRefundNewGetdetailRequest 初始化TaobaoAlitripIeAgentRefundNewGetdetailAPIRequest对象
func NewTaobaoAlitripIeAgentRefundNewGetdetailRequest() *TaobaoAlitripIeAgentRefundNewGetdetailAPIRequest {
	return &TaobaoAlitripIeAgentRefundNewGetdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentRefundNewGetdetailAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.refund.new.getdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentRefundNewGetdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripIeAgentRefundNewGetdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamRefundOrderQueryDetailRq is ParamRefundOrderQueryDetailRq Setter
// 请求
func (r *TaobaoAlitripIeAgentRefundNewGetdetailAPIRequest) SetParamRefundOrderQueryDetailRq(_paramRefundOrderQueryDetailRq *RefundOrderQueryDetailRq) error {
	r._paramRefundOrderQueryDetailRq = _paramRefundOrderQueryDetailRq
	r.Set("param_refund_order_query_detail_rq", _paramRefundOrderQueryDetailRq)
	return nil
}

// GetParamRefundOrderQueryDetailRq ParamRefundOrderQueryDetailRq Getter
func (r TaobaoAlitripIeAgentRefundNewGetdetailAPIRequest) GetParamRefundOrderQueryDetailRq() *RefundOrderQueryDetailRq {
	return r._paramRefundOrderQueryDetailRq
}
