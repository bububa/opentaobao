package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketStoreGetAPIRequest 获取电子凭证预约门店信息 API请求
// taobao.vmarket.eticket.store.get
//
// 用于给外部商家查询电子凭证预约门店信息
type TaobaoVmarketEticketStoreGetAPIRequest struct {
	model.Params
	// 订单ID
	_orderId int64
}

// NewTaobaoVmarketEticketStoreGetRequest 初始化TaobaoVmarketEticketStoreGetAPIRequest对象
func NewTaobaoVmarketEticketStoreGetRequest() *TaobaoVmarketEticketStoreGetAPIRequest {
	return &TaobaoVmarketEticketStoreGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketStoreGetAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.store.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketStoreGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoVmarketEticketStoreGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单ID
func (r *TaobaoVmarketEticketStoreGetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoVmarketEticketStoreGetAPIRequest) GetOrderId() int64 {
	return r._orderId
}
