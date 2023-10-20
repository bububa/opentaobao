package happytrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiOrderAssignAPIRequest 订单指派 API请求
// alibaba.happytrip.taxi.order.assign
//
// 通知供应商订单指派成功
type AlibabaHappytripTaxiOrderAssignAPIRequest struct {
	model.Params
	// 供应商订单号
	_orderId string
}

// NewAlibabaHappytripTaxiOrderAssignRequest 初始化AlibabaHappytripTaxiOrderAssignAPIRequest对象
func NewAlibabaHappytripTaxiOrderAssignRequest() *AlibabaHappytripTaxiOrderAssignAPIRequest {
	return &AlibabaHappytripTaxiOrderAssignAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHappytripTaxiOrderAssignAPIRequest) Reset() {
	r._orderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiOrderAssignAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.order.assign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiOrderAssignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHappytripTaxiOrderAssignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 供应商订单号
func (r *AlibabaHappytripTaxiOrderAssignAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaHappytripTaxiOrderAssignAPIRequest) GetOrderId() string {
	return r._orderId
}

var poolAlibabaHappytripTaxiOrderAssignAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHappytripTaxiOrderAssignRequest()
	},
}

// GetAlibabaHappytripTaxiOrderAssignRequest 从 sync.Pool 获取 AlibabaHappytripTaxiOrderAssignAPIRequest
func GetAlibabaHappytripTaxiOrderAssignAPIRequest() *AlibabaHappytripTaxiOrderAssignAPIRequest {
	return poolAlibabaHappytripTaxiOrderAssignAPIRequest.Get().(*AlibabaHappytripTaxiOrderAssignAPIRequest)
}

// ReleaseAlibabaHappytripTaxiOrderAssignAPIRequest 将 AlibabaHappytripTaxiOrderAssignAPIRequest 放入 sync.Pool
func ReleaseAlibabaHappytripTaxiOrderAssignAPIRequest(v *AlibabaHappytripTaxiOrderAssignAPIRequest) {
	v.Reset()
	poolAlibabaHappytripTaxiOrderAssignAPIRequest.Put(v)
}
