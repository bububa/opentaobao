package aliexpresssumaitong

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressTaxationCalculateOpenQueryAPIRequest 关务所需的申报清关字段 API请求
// aliexpress.taxation.calculate.open.query
//
// 关务所需的申报清关字段
type AliexpressTaxationCalculateOpenQueryAPIRequest struct {
	model.Params
	// 主订单id
	_orderId string
}

// NewAliexpressTaxationCalculateOpenQueryRequest 初始化AliexpressTaxationCalculateOpenQueryAPIRequest对象
func NewAliexpressTaxationCalculateOpenQueryRequest() *AliexpressTaxationCalculateOpenQueryAPIRequest {
	return &AliexpressTaxationCalculateOpenQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressTaxationCalculateOpenQueryAPIRequest) Reset() {
	r._orderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressTaxationCalculateOpenQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.taxation.calculate.open.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressTaxationCalculateOpenQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressTaxationCalculateOpenQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 主订单id
func (r *AliexpressTaxationCalculateOpenQueryAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AliexpressTaxationCalculateOpenQueryAPIRequest) GetOrderId() string {
	return r._orderId
}

var poolAliexpressTaxationCalculateOpenQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressTaxationCalculateOpenQueryRequest()
	},
}

// GetAliexpressTaxationCalculateOpenQueryRequest 从 sync.Pool 获取 AliexpressTaxationCalculateOpenQueryAPIRequest
func GetAliexpressTaxationCalculateOpenQueryAPIRequest() *AliexpressTaxationCalculateOpenQueryAPIRequest {
	return poolAliexpressTaxationCalculateOpenQueryAPIRequest.Get().(*AliexpressTaxationCalculateOpenQueryAPIRequest)
}

// ReleaseAliexpressTaxationCalculateOpenQueryAPIRequest 将 AliexpressTaxationCalculateOpenQueryAPIRequest 放入 sync.Pool
func ReleaseAliexpressTaxationCalculateOpenQueryAPIRequest(v *AliexpressTaxationCalculateOpenQueryAPIRequest) {
	v.Reset()
	poolAliexpressTaxationCalculateOpenQueryAPIRequest.Put(v)
}
