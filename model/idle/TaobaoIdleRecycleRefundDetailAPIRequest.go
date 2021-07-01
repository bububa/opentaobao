package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoIdleRecycleRefundDetailAPIRequest
闲鱼回收退款详情V2 API请求
taobao.idle.recycle.refund.detail

回收订单退款详情，主要包括退款状态，超时时间，和同意退款的卖家退货地址信息 */
type TaobaoIdleRecycleRefundDetailAPIRequest struct {
	model.Params
	// 订单号
	_bizOrderId int64
}

// NewTaobaoIdleRecycleRefundDetailRequest 初始化TaobaoIdleRecycleRefundDetailAPIRequest对象
func NewTaobaoIdleRecycleRefundDetailRequest() *TaobaoIdleRecycleRefundDetailAPIRequest {
	return &TaobaoIdleRecycleRefundDetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoIdleRecycleRefundDetailAPIRequest) GetApiMethodName() string {
	return "taobao.idle.recycle.refund.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoIdleRecycleRefundDetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BizOrderId Setter
// 订单号
func (r *TaobaoIdleRecycleRefundDetailAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// Get BizOrderId Getter
func (r TaobaoIdleRecycleRefundDetailAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}
