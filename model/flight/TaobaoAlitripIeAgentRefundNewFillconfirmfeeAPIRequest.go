package flight

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest) Reset() {
	r._paramRefundOrderFillConfirmFeeRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.refund.new.fillconfirmfee"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripIeAgentRefundNewFillconfirmfeeRequest()
	},
}

// GetTaobaoAlitripIeAgentRefundNewFillconfirmfeeRequest 从 sync.Pool 获取 TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest
func GetTaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest() *TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest {
	return poolTaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest.Get().(*TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest)
}

// ReleaseTaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest 将 TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest(v *TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest) {
	v.Reset()
	poolTaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIRequest.Put(v)
}
