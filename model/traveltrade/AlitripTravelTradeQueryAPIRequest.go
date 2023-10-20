package traveltrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelTradeQueryAPIRequest 飞猪度假-订单详情查询接口 API请求
// alitrip.travel.trade.query
//
// 飞猪度假订单详情查询接口
type AlitripTravelTradeQueryAPIRequest struct {
	model.Params
	// 主订单id
	_orderId int64
}

// NewAlitripTravelTradeQueryRequest 初始化AlitripTravelTradeQueryAPIRequest对象
func NewAlitripTravelTradeQueryRequest() *AlitripTravelTradeQueryAPIRequest {
	return &AlitripTravelTradeQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripTravelTradeQueryAPIRequest) Reset() {
	r._orderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTravelTradeQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.trade.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTravelTradeQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTravelTradeQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 主订单id
func (r *AlitripTravelTradeQueryAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitripTravelTradeQueryAPIRequest) GetOrderId() int64 {
	return r._orderId
}

var poolAlitripTravelTradeQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripTravelTradeQueryRequest()
	},
}

// GetAlitripTravelTradeQueryRequest 从 sync.Pool 获取 AlitripTravelTradeQueryAPIRequest
func GetAlitripTravelTradeQueryAPIRequest() *AlitripTravelTradeQueryAPIRequest {
	return poolAlitripTravelTradeQueryAPIRequest.Get().(*AlitripTravelTradeQueryAPIRequest)
}

// ReleaseAlitripTravelTradeQueryAPIRequest 将 AlitripTravelTradeQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripTravelTradeQueryAPIRequest(v *AlitripTravelTradeQueryAPIRequest) {
	v.Reset()
	poolAlitripTravelTradeQueryAPIRequest.Put(v)
}
