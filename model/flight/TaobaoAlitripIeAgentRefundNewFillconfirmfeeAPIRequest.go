package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripieagentrefundnewfillconfirmfeeAPIRequest 新模型-回填申请单费用 API请求
// taobao.alitrip.ie.agent.refund.new.fillconfirmfee
//
// 1. 回填退票费用
type TaobaoalitripieagentrefundnewfillconfirmfeeAPIRequest struct {
	model.Params
	// 请求
	_paramRefundOrderFillConfirmFeeRq *RefundOrderFillConfirmFeeRq
}

// NewTaobaoalitripieagentrefundnewfillconfirmfeeRequest 初始化TaobaoalitripieagentrefundnewfillconfirmfeeAPIRequest对象
func NewTaobaoalitripieagentrefundnewfillconfirmfeeRequest() *TaobaoalitripieagentrefundnewfillconfirmfeeAPIRequest {
	return &TaobaoalitripieagentrefundnewfillconfirmfeeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripieagentrefundnewfillconfirmfeeAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.refund.new.fillconfirmfee"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripieagentrefundnewfillconfirmfeeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripieagentrefundnewfillconfirmfeeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamRefundOrderFillConfirmFeeRq is ParamRefundOrderFillConfirmFeeRq Setter
// 请求
func (r *TaobaoalitripieagentrefundnewfillconfirmfeeAPIRequest) SetParamRefundOrderFillConfirmFeeRq(_paramRefundOrderFillConfirmFeeRq *RefundOrderFillConfirmFeeRq) error {
	r._paramRefundOrderFillConfirmFeeRq = _paramRefundOrderFillConfirmFeeRq
	r.Set("param_refund_order_fill_confirm_fee_rq", _paramRefundOrderFillConfirmFeeRq)
	return nil
}

// GetParamRefundOrderFillConfirmFeeRq ParamRefundOrderFillConfirmFeeRq Getter
func (r TaobaoalitripieagentrefundnewfillconfirmfeeAPIRequest) GetParamRefundOrderFillConfirmFeeRq() *RefundOrderFillConfirmFeeRq {
	return r._paramRefundOrderFillConfirmFeeRq
}
