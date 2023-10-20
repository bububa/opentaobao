package alitripcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripCarOrderRefundAPIRequest 用户投诉达成一致后给用户退款 API请求
// taobao.alitrip.car.order.refund
//
// 用户投诉后，供应商客服与客户沟通达成一致后通知飞猪给客户退款。退款金额以接口回调金额为准。
type TaobaoAlitripCarOrderRefundAPIRequest struct {
	model.Params
	// 退款对象
	_paramOrderRefund *OrderRefund
}

// NewTaobaoAlitripCarOrderRefundRequest 初始化TaobaoAlitripCarOrderRefundAPIRequest对象
func NewTaobaoAlitripCarOrderRefundRequest() *TaobaoAlitripCarOrderRefundAPIRequest {
	return &TaobaoAlitripCarOrderRefundAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripCarOrderRefundAPIRequest) Reset() {
	r._paramOrderRefund = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripCarOrderRefundAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.car.order.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripCarOrderRefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripCarOrderRefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOrderRefund is ParamOrderRefund Setter
// 退款对象
func (r *TaobaoAlitripCarOrderRefundAPIRequest) SetParamOrderRefund(_paramOrderRefund *OrderRefund) error {
	r._paramOrderRefund = _paramOrderRefund
	r.Set("param_order_refund", _paramOrderRefund)
	return nil
}

// GetParamOrderRefund ParamOrderRefund Getter
func (r TaobaoAlitripCarOrderRefundAPIRequest) GetParamOrderRefund() *OrderRefund {
	return r._paramOrderRefund
}

var poolTaobaoAlitripCarOrderRefundAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripCarOrderRefundRequest()
	},
}

// GetTaobaoAlitripCarOrderRefundRequest 从 sync.Pool 获取 TaobaoAlitripCarOrderRefundAPIRequest
func GetTaobaoAlitripCarOrderRefundAPIRequest() *TaobaoAlitripCarOrderRefundAPIRequest {
	return poolTaobaoAlitripCarOrderRefundAPIRequest.Get().(*TaobaoAlitripCarOrderRefundAPIRequest)
}

// ReleaseTaobaoAlitripCarOrderRefundAPIRequest 将 TaobaoAlitripCarOrderRefundAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripCarOrderRefundAPIRequest(v *TaobaoAlitripCarOrderRefundAPIRequest) {
	v.Reset()
	poolTaobaoAlitripCarOrderRefundAPIRequest.Put(v)
}
