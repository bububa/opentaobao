package idle

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoIdleRecycleRefundApplyAPIRequest) GetApiMethodName() string {
	return "taobao.idle.recycle.refund.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoIdleRecycleRefundApplyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
