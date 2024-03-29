package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoIdleRecycleRefundApplyAPIRequest 闲鱼回收交易退款申请V2 API请求
// taobao.idle.recycle.refund.apply
//
// 回收商买家申请退款
type TaobaoIdleRecycleRefundApplyAPIRequest struct {
	model.Params
	// 退款申请
	_refundApply *RecycleRefundTopRequest
}

// NewTaobaoIdleRecycleRefundApplyRequest 初始化TaobaoIdleRecycleRefundApplyAPIRequest对象
func NewTaobaoIdleRecycleRefundApplyRequest() *TaobaoIdleRecycleRefundApplyAPIRequest {
	return &TaobaoIdleRecycleRefundApplyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoIdleRecycleRefundApplyAPIRequest) Reset() {
	r._refundApply = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoIdleRecycleRefundApplyAPIRequest) GetApiMethodName() string {
	return "taobao.idle.recycle.refund.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoIdleRecycleRefundApplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoIdleRecycleRefundApplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundApply is RefundApply Setter
// 退款申请
func (r *TaobaoIdleRecycleRefundApplyAPIRequest) SetRefundApply(_refundApply *RecycleRefundTopRequest) error {
	r._refundApply = _refundApply
	r.Set("refund_apply", _refundApply)
	return nil
}

// GetRefundApply RefundApply Getter
func (r TaobaoIdleRecycleRefundApplyAPIRequest) GetRefundApply() *RecycleRefundTopRequest {
	return r._refundApply
}

var poolTaobaoIdleRecycleRefundApplyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoIdleRecycleRefundApplyRequest()
	},
}

// GetTaobaoIdleRecycleRefundApplyRequest 从 sync.Pool 获取 TaobaoIdleRecycleRefundApplyAPIRequest
func GetTaobaoIdleRecycleRefundApplyAPIRequest() *TaobaoIdleRecycleRefundApplyAPIRequest {
	return poolTaobaoIdleRecycleRefundApplyAPIRequest.Get().(*TaobaoIdleRecycleRefundApplyAPIRequest)
}

// ReleaseTaobaoIdleRecycleRefundApplyAPIRequest 将 TaobaoIdleRecycleRefundApplyAPIRequest 放入 sync.Pool
func ReleaseTaobaoIdleRecycleRefundApplyAPIRequest(v *TaobaoIdleRecycleRefundApplyAPIRequest) {
	v.Reset()
	poolTaobaoIdleRecycleRefundApplyAPIRequest.Put(v)
}
