package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoIdleRecycleRefundCancleapplyAPIRequest 闲鱼回收取消退款申请V2 API请求
// taobao.idle.recycle.refund.cancleapply
//
// 回收商的回收订单取消退款申请
type TaobaoIdleRecycleRefundCancleapplyAPIRequest struct {
	model.Params
	// 订单号
	_bizOrderId int64
}

// NewTaobaoIdleRecycleRefundCancleapplyRequest 初始化TaobaoIdleRecycleRefundCancleapplyAPIRequest对象
func NewTaobaoIdleRecycleRefundCancleapplyRequest() *TaobaoIdleRecycleRefundCancleapplyAPIRequest {
	return &TaobaoIdleRecycleRefundCancleapplyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoIdleRecycleRefundCancleapplyAPIRequest) Reset() {
	r._bizOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoIdleRecycleRefundCancleapplyAPIRequest) GetApiMethodName() string {
	return "taobao.idle.recycle.refund.cancleapply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoIdleRecycleRefundCancleapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoIdleRecycleRefundCancleapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// 订单号
func (r *TaobaoIdleRecycleRefundCancleapplyAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TaobaoIdleRecycleRefundCancleapplyAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

var poolTaobaoIdleRecycleRefundCancleapplyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoIdleRecycleRefundCancleapplyRequest()
	},
}

// GetTaobaoIdleRecycleRefundCancleapplyRequest 从 sync.Pool 获取 TaobaoIdleRecycleRefundCancleapplyAPIRequest
func GetTaobaoIdleRecycleRefundCancleapplyAPIRequest() *TaobaoIdleRecycleRefundCancleapplyAPIRequest {
	return poolTaobaoIdleRecycleRefundCancleapplyAPIRequest.Get().(*TaobaoIdleRecycleRefundCancleapplyAPIRequest)
}

// ReleaseTaobaoIdleRecycleRefundCancleapplyAPIRequest 将 TaobaoIdleRecycleRefundCancleapplyAPIRequest 放入 sync.Pool
func ReleaseTaobaoIdleRecycleRefundCancleapplyAPIRequest(v *TaobaoIdleRecycleRefundCancleapplyAPIRequest) {
	v.Reset()
	poolTaobaoIdleRecycleRefundCancleapplyAPIRequest.Put(v)
}
