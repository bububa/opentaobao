package eticket

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoVmarketEticketStoreGetAPIRequest) Reset() {
	r._orderId = 0
	r.Params.ToZero()
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

var poolTaobaoVmarketEticketStoreGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoVmarketEticketStoreGetRequest()
	},
}

// GetTaobaoVmarketEticketStoreGetRequest 从 sync.Pool 获取 TaobaoVmarketEticketStoreGetAPIRequest
func GetTaobaoVmarketEticketStoreGetAPIRequest() *TaobaoVmarketEticketStoreGetAPIRequest {
	return poolTaobaoVmarketEticketStoreGetAPIRequest.Get().(*TaobaoVmarketEticketStoreGetAPIRequest)
}

// ReleaseTaobaoVmarketEticketStoreGetAPIRequest 将 TaobaoVmarketEticketStoreGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoVmarketEticketStoreGetAPIRequest(v *TaobaoVmarketEticketStoreGetAPIRequest) {
	v.Reset()
	poolTaobaoVmarketEticketStoreGetAPIRequest.Put(v)
}
