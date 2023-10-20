package happytrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiOrderGetAPIRequest 订单详情 API请求
// alibaba.happytrip.taxi.order.get
//
// 获取订单状态及详情
type AlibabaHappytripTaxiOrderGetAPIRequest struct {
	model.Params
	// 订单id
	_orderId string
}

// NewAlibabaHappytripTaxiOrderGetRequest 初始化AlibabaHappytripTaxiOrderGetAPIRequest对象
func NewAlibabaHappytripTaxiOrderGetRequest() *AlibabaHappytripTaxiOrderGetAPIRequest {
	return &AlibabaHappytripTaxiOrderGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHappytripTaxiOrderGetAPIRequest) Reset() {
	r._orderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiOrderGetAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHappytripTaxiOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabaHappytripTaxiOrderGetAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaHappytripTaxiOrderGetAPIRequest) GetOrderId() string {
	return r._orderId
}

var poolAlibabaHappytripTaxiOrderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHappytripTaxiOrderGetRequest()
	},
}

// GetAlibabaHappytripTaxiOrderGetRequest 从 sync.Pool 获取 AlibabaHappytripTaxiOrderGetAPIRequest
func GetAlibabaHappytripTaxiOrderGetAPIRequest() *AlibabaHappytripTaxiOrderGetAPIRequest {
	return poolAlibabaHappytripTaxiOrderGetAPIRequest.Get().(*AlibabaHappytripTaxiOrderGetAPIRequest)
}

// ReleaseAlibabaHappytripTaxiOrderGetAPIRequest 将 AlibabaHappytripTaxiOrderGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaHappytripTaxiOrderGetAPIRequest(v *AlibabaHappytripTaxiOrderGetAPIRequest) {
	v.Reset()
	poolAlibabaHappytripTaxiOrderGetAPIRequest.Put(v)
}
