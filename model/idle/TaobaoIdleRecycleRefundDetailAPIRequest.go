package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoIdleRecycleRefundDetailAPIRequest 闲鱼回收退款详情V2 API请求
// taobao.idle.recycle.refund.detail
//
// 回收订单退款详情，主要包括退款状态，超时时间，和同意退款的卖家退货地址信息
type TaobaoIdleRecycleRefundDetailAPIRequest struct {
	model.Params
	// 订单号
	_bizOrderId int64
}

// NewTaobaoIdleRecycleRefundDetailRequest 初始化TaobaoIdleRecycleRefundDetailAPIRequest对象
func NewTaobaoIdleRecycleRefundDetailRequest() *TaobaoIdleRecycleRefundDetailAPIRequest {
	return &TaobaoIdleRecycleRefundDetailAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoIdleRecycleRefundDetailAPIRequest) Reset() {
	r._bizOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoIdleRecycleRefundDetailAPIRequest) GetApiMethodName() string {
	return "taobao.idle.recycle.refund.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoIdleRecycleRefundDetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoIdleRecycleRefundDetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// 订单号
func (r *TaobaoIdleRecycleRefundDetailAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TaobaoIdleRecycleRefundDetailAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

var poolTaobaoIdleRecycleRefundDetailAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoIdleRecycleRefundDetailRequest()
	},
}

// GetTaobaoIdleRecycleRefundDetailRequest 从 sync.Pool 获取 TaobaoIdleRecycleRefundDetailAPIRequest
func GetTaobaoIdleRecycleRefundDetailAPIRequest() *TaobaoIdleRecycleRefundDetailAPIRequest {
	return poolTaobaoIdleRecycleRefundDetailAPIRequest.Get().(*TaobaoIdleRecycleRefundDetailAPIRequest)
}

// ReleaseTaobaoIdleRecycleRefundDetailAPIRequest 将 TaobaoIdleRecycleRefundDetailAPIRequest 放入 sync.Pool
func ReleaseTaobaoIdleRecycleRefundDetailAPIRequest(v *TaobaoIdleRecycleRefundDetailAPIRequest) {
	v.Reset()
	poolTaobaoIdleRecycleRefundDetailAPIRequest.Put(v)
}
