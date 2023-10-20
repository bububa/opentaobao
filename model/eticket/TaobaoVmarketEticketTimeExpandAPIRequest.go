package eticket

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoVmarketEticketTimeExpandAPIRequest) Reset() {
	r._orderId = 0
	r._expandDays = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketTimeExpandAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.time.expand"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketTimeExpandAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoVmarketEticketTimeExpandAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoVmarketEticketTimeExpandAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoVmarketEticketTimeExpandRequest()
	},
}

// GetTaobaoVmarketEticketTimeExpandRequest 从 sync.Pool 获取 TaobaoVmarketEticketTimeExpandAPIRequest
func GetTaobaoVmarketEticketTimeExpandAPIRequest() *TaobaoVmarketEticketTimeExpandAPIRequest {
	return poolTaobaoVmarketEticketTimeExpandAPIRequest.Get().(*TaobaoVmarketEticketTimeExpandAPIRequest)
}

// ReleaseTaobaoVmarketEticketTimeExpandAPIRequest 将 TaobaoVmarketEticketTimeExpandAPIRequest 放入 sync.Pool
func ReleaseTaobaoVmarketEticketTimeExpandAPIRequest(v *TaobaoVmarketEticketTimeExpandAPIRequest) {
	v.Reset()
	poolTaobaoVmarketEticketTimeExpandAPIRequest.Put(v)
}
