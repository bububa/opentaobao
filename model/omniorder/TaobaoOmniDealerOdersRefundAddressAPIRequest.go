package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniDealerOdersRefundAddressAPIRequest 经销商查询逆向退货地址 API请求
// taobao.omni.dealer.oders.refund.address
//
// 经销商查询逆向退货地址
type TaobaoOmniDealerOdersRefundAddressAPIRequest struct {
	model.Params
	// 经销商子订单id
	_orderId int64
}

// NewTaobaoOmniDealerOdersRefundAddressRequest 初始化TaobaoOmniDealerOdersRefundAddressAPIRequest对象
func NewTaobaoOmniDealerOdersRefundAddressRequest() *TaobaoOmniDealerOdersRefundAddressAPIRequest {
	return &TaobaoOmniDealerOdersRefundAddressAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniDealerOdersRefundAddressAPIRequest) GetApiMethodName() string {
	return "taobao.omni.dealer.oders.refund.address"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniDealerOdersRefundAddressAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderId is OrderId Setter
// 经销商子订单id
func (r *TaobaoOmniDealerOdersRefundAddressAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoOmniDealerOdersRefundAddressAPIRequest) GetOrderId() int64 {
	return r._orderId
}
