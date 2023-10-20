package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoidlerecyclerefunddetailAPIRequest 闲鱼回收退款详情V2 API请求
// taobao.idle.recycle.refund.detail
//
// 回收订单退款详情，主要包括退款状态，超时时间，和同意退款的卖家退货地址信息
type TaobaoidlerecyclerefunddetailAPIRequest struct {
	model.Params
	// 订单号
	_bizOrderId int64
}

// NewTaobaoidlerecyclerefunddetailRequest 初始化TaobaoidlerecyclerefunddetailAPIRequest对象
func NewTaobaoidlerecyclerefunddetailRequest() *TaobaoidlerecyclerefunddetailAPIRequest {
	return &TaobaoidlerecyclerefunddetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoidlerecyclerefunddetailAPIRequest) GetApiMethodName() string {
	return "taobao.idle.recycle.refund.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoidlerecyclerefunddetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoidlerecyclerefunddetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// 订单号
func (r *TaobaoidlerecyclerefunddetailAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TaobaoidlerecyclerefunddetailAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}
