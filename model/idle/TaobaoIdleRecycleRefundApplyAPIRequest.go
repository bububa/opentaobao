package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoidlerecyclerefundapplyAPIRequest 闲鱼回收交易退款申请V2 API请求
// taobao.idle.recycle.refund.apply
//
// 回收商买家申请退款
type TaobaoidlerecyclerefundapplyAPIRequest struct {
	model.Params
	// 退款申请
	_refundApply *RecycleRefundTopRequest
}

// NewTaobaoidlerecyclerefundapplyRequest 初始化TaobaoidlerecyclerefundapplyAPIRequest对象
func NewTaobaoidlerecyclerefundapplyRequest() *TaobaoidlerecyclerefundapplyAPIRequest {
	return &TaobaoidlerecyclerefundapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoidlerecyclerefundapplyAPIRequest) GetApiMethodName() string {
	return "taobao.idle.recycle.refund.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoidlerecyclerefundapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoidlerecyclerefundapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundApply is RefundApply Setter
// 退款申请
func (r *TaobaoidlerecyclerefundapplyAPIRequest) SetRefundApply(_refundApply *RecycleRefundTopRequest) error {
	r._refundApply = _refundApply
	r.Set("refund_apply", _refundApply)
	return nil
}

// GetRefundApply RefundApply Getter
func (r TaobaoidlerecyclerefundapplyAPIRequest) GetRefundApply() *RecycleRefundTopRequest {
	return r._refundApply
}
