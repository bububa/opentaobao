package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest 新模型-回填申请单费用 API请求
// taobao.alitrip.ie.agent.refund.new.fillconfirmfee
//
// 1. 回填退票费用
type TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest struct {
	model.Params
	// 请求
	_paramRefundOrderFillConfirmFeeRq *RefundOrderFillConfirmFeeRq
}

// NewTaobaoAlitripIeAgentRefundNewFillconfirmfeeRequest 初始化TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest对象
func NewTaobaoAlitripIeAgentRefundNewFillconfirmfeeRequest() *TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest {
	return &TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.refund.new.fillconfirmfee"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamRefundOrderFillConfirmFeeRq is ParamRefundOrderFillConfirmFeeRq Setter
// 请求
func (r *TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest) SetParamRefundOrderFillConfirmFeeRq(_paramRefundOrderFillConfirmFeeRq *RefundOrderFillConfirmFeeRq) error {
	r._paramRefundOrderFillConfirmFeeRq = _paramRefundOrderFillConfirmFeeRq
	r.Set("param_refund_order_fill_confirm_fee_rq", _paramRefundOrderFillConfirmFeeRq)
	return nil
}

// GetParamRefundOrderFillConfirmFeeRq ParamRefundOrderFillConfirmFeeRq Getter
func (r TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest) GetParamRefundOrderFillConfirmFeeRq() *RefundOrderFillConfirmFeeRq {
	return r._paramRefundOrderFillConfirmFeeRq
}
