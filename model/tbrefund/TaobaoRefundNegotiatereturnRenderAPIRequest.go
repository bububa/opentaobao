package tbrefund

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRefundNegotiatereturnRenderAPIRequest 协商退货退款渲染 API请求
// taobao.refund.negotiatereturn.render
//
// 协商退货退款渲染
type TaobaoRefundNegotiatereturnRenderAPIRequest struct {
	model.Params
	// 退款编号
	_refundId int64
}

// NewTaobaoRefundNegotiatereturnRenderRequest 初始化TaobaoRefundNegotiatereturnRenderAPIRequest对象
func NewTaobaoRefundNegotiatereturnRenderRequest() *TaobaoRefundNegotiatereturnRenderAPIRequest {
	return &TaobaoRefundNegotiatereturnRenderAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRefundNegotiatereturnRenderAPIRequest) Reset() {
	r._refundId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRefundNegotiatereturnRenderAPIRequest) GetApiMethodName() string {
	return "taobao.refund.negotiatereturn.render"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRefundNegotiatereturnRenderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRefundNegotiatereturnRenderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundId is RefundId Setter
// 退款编号
func (r *TaobaoRefundNegotiatereturnRenderAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoRefundNegotiatereturnRenderAPIRequest) GetRefundId() int64 {
	return r._refundId
}

var poolTaobaoRefundNegotiatereturnRenderAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRefundNegotiatereturnRenderRequest()
	},
}

// GetTaobaoRefundNegotiatereturnRenderRequest 从 sync.Pool 获取 TaobaoRefundNegotiatereturnRenderAPIRequest
func GetTaobaoRefundNegotiatereturnRenderAPIRequest() *TaobaoRefundNegotiatereturnRenderAPIRequest {
	return poolTaobaoRefundNegotiatereturnRenderAPIRequest.Get().(*TaobaoRefundNegotiatereturnRenderAPIRequest)
}

// ReleaseTaobaoRefundNegotiatereturnRenderAPIRequest 将 TaobaoRefundNegotiatereturnRenderAPIRequest 放入 sync.Pool
func ReleaseTaobaoRefundNegotiatereturnRenderAPIRequest(v *TaobaoRefundNegotiatereturnRenderAPIRequest) {
	v.Reset()
	poolTaobaoRefundNegotiatereturnRenderAPIRequest.Put(v)
}
