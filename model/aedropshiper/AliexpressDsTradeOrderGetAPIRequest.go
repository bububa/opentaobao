package aedropshiper

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressDsTradeOrderGetAPIRequest 交易订单查询 API请求
// aliexpress.ds.trade.order.get
//
// 交易订单查询
type AliexpressDsTradeOrderGetAPIRequest struct {
	model.Params
	// AE order id
	_orderId int64
}

// NewAliexpressDsTradeOrderGetRequest 初始化AliexpressDsTradeOrderGetAPIRequest对象
func NewAliexpressDsTradeOrderGetRequest() *AliexpressDsTradeOrderGetAPIRequest {
	return &AliexpressDsTradeOrderGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressDsTradeOrderGetAPIRequest) Reset() {
	r._orderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressDsTradeOrderGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.ds.trade.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressDsTradeOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressDsTradeOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// AE order id
func (r *AliexpressDsTradeOrderGetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AliexpressDsTradeOrderGetAPIRequest) GetOrderId() int64 {
	return r._orderId
}

var poolAliexpressDsTradeOrderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressDsTradeOrderGetRequest()
	},
}

// GetAliexpressDsTradeOrderGetRequest 从 sync.Pool 获取 AliexpressDsTradeOrderGetAPIRequest
func GetAliexpressDsTradeOrderGetAPIRequest() *AliexpressDsTradeOrderGetAPIRequest {
	return poolAliexpressDsTradeOrderGetAPIRequest.Get().(*AliexpressDsTradeOrderGetAPIRequest)
}

// ReleaseAliexpressDsTradeOrderGetAPIRequest 将 AliexpressDsTradeOrderGetAPIRequest 放入 sync.Pool
func ReleaseAliexpressDsTradeOrderGetAPIRequest(v *AliexpressDsTradeOrderGetAPIRequest) {
	v.Reset()
	poolAliexpressDsTradeOrderGetAPIRequest.Put(v)
}
