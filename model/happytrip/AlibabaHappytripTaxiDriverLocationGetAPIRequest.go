package happytrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiDriverLocationGetAPIRequest 司机位置 API请求
// alibaba.happytrip.taxi.driver.location.get
//
// 获取司机实时位置
type AlibabaHappytripTaxiDriverLocationGetAPIRequest struct {
	model.Params
	// 供应商订单号
	_orderId string
}

// NewAlibabaHappytripTaxiDriverLocationGetRequest 初始化AlibabaHappytripTaxiDriverLocationGetAPIRequest对象
func NewAlibabaHappytripTaxiDriverLocationGetRequest() *AlibabaHappytripTaxiDriverLocationGetAPIRequest {
	return &AlibabaHappytripTaxiDriverLocationGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHappytripTaxiDriverLocationGetAPIRequest) Reset() {
	r._orderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiDriverLocationGetAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.driver.location.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiDriverLocationGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHappytripTaxiDriverLocationGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 供应商订单号
func (r *AlibabaHappytripTaxiDriverLocationGetAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaHappytripTaxiDriverLocationGetAPIRequest) GetOrderId() string {
	return r._orderId
}

var poolAlibabaHappytripTaxiDriverLocationGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHappytripTaxiDriverLocationGetRequest()
	},
}

// GetAlibabaHappytripTaxiDriverLocationGetRequest 从 sync.Pool 获取 AlibabaHappytripTaxiDriverLocationGetAPIRequest
func GetAlibabaHappytripTaxiDriverLocationGetAPIRequest() *AlibabaHappytripTaxiDriverLocationGetAPIRequest {
	return poolAlibabaHappytripTaxiDriverLocationGetAPIRequest.Get().(*AlibabaHappytripTaxiDriverLocationGetAPIRequest)
}

// ReleaseAlibabaHappytripTaxiDriverLocationGetAPIRequest 将 AlibabaHappytripTaxiDriverLocationGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaHappytripTaxiDriverLocationGetAPIRequest(v *AlibabaHappytripTaxiDriverLocationGetAPIRequest) {
	v.Reset()
	poolAlibabaHappytripTaxiDriverLocationGetAPIRequest.Put(v)
}
