package omniorder

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniDealerOdersRefundAddressAPIRequest) Reset() {
	r._orderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniDealerOdersRefundAddressAPIRequest) GetApiMethodName() string {
	return "taobao.omni.dealer.oders.refund.address"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniDealerOdersRefundAddressAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniDealerOdersRefundAddressAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoOmniDealerOdersRefundAddressAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniDealerOdersRefundAddressRequest()
	},
}

// GetTaobaoOmniDealerOdersRefundAddressRequest 从 sync.Pool 获取 TaobaoOmniDealerOdersRefundAddressAPIRequest
func GetTaobaoOmniDealerOdersRefundAddressAPIRequest() *TaobaoOmniDealerOdersRefundAddressAPIRequest {
	return poolTaobaoOmniDealerOdersRefundAddressAPIRequest.Get().(*TaobaoOmniDealerOdersRefundAddressAPIRequest)
}

// ReleaseTaobaoOmniDealerOdersRefundAddressAPIRequest 将 TaobaoOmniDealerOdersRefundAddressAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniDealerOdersRefundAddressAPIRequest(v *TaobaoOmniDealerOdersRefundAddressAPIRequest) {
	v.Reset()
	poolTaobaoOmniDealerOdersRefundAddressAPIRequest.Put(v)
}
