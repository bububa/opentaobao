package aliexpresssumaitong

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressTaxationCalculateOpenQueryAPIRequest
关务所需的申报清关字段 API请求
aliexpress.taxation.calculate.open.query

关务所需的申报清关字段 */
type AliexpressTaxationCalculateOpenQueryAPIRequest struct {
	model.Params
	// 主订单id
	_orderId string
}

// NewAliexpressTaxationCalculateOpenQueryRequest 初始化AliexpressTaxationCalculateOpenQueryAPIRequest对象
func NewAliexpressTaxationCalculateOpenQueryRequest() *AliexpressTaxationCalculateOpenQueryAPIRequest {
	return &AliexpressTaxationCalculateOpenQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressTaxationCalculateOpenQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.taxation.calculate.open.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressTaxationCalculateOpenQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderId Setter
// 主订单id
func (r *AliexpressTaxationCalculateOpenQueryAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r AliexpressTaxationCalculateOpenQueryAPIRequest) GetOrderId() string {
	return r._orderId
}
