package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketTimeExpandAPIRequest 订单延时接口 API请求
// taobao.vmarket.eticket.time.expand
//
// 提供码商操作订单延期接口
type TaobaoVmarketEticketTimeExpandAPIRequest struct {
	model.Params
	// 订单ID
	_orderId int64
	// 延长天数，延长时间=当前过期时间+延长天数
	_expandDays int64
}

// NewTaobaoVmarketEticketTimeExpandRequest 初始化TaobaoVmarketEticketTimeExpandAPIRequest对象
func NewTaobaoVmarketEticketTimeExpandRequest() *TaobaoVmarketEticketTimeExpandAPIRequest {
	return &TaobaoVmarketEticketTimeExpandAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketTimeExpandAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.time.expand"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketTimeExpandAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderId is OrderId Setter
// 订单ID
func (r *TaobaoVmarketEticketTimeExpandAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoVmarketEticketTimeExpandAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetExpandDays is ExpandDays Setter
// 延长天数，延长时间=当前过期时间+延长天数
func (r *TaobaoVmarketEticketTimeExpandAPIRequest) SetExpandDays(_expandDays int64) error {
	r._expandDays = _expandDays
	r.Set("expand_days", _expandDays)
	return nil
}

// GetExpandDays ExpandDays Getter
func (r TaobaoVmarketEticketTimeExpandAPIRequest) GetExpandDays() int64 {
	return r._expandDays
}
