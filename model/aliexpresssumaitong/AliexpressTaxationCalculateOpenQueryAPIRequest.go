package aliexpresssumaitong

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresstaxationcalculateopenqueryAPIRequest 关务所需的申报清关字段 API请求
// aliexpress.taxation.calculate.open.query
//
// 关务所需的申报清关字段
type AliexpresstaxationcalculateopenqueryAPIRequest struct {
	model.Params
	// 主订单id
	_orderId string
}

// NewAliexpresstaxationcalculateopenqueryRequest 初始化AliexpresstaxationcalculateopenqueryAPIRequest对象
func NewAliexpresstaxationcalculateopenqueryRequest() *AliexpresstaxationcalculateopenqueryAPIRequest {
	return &AliexpresstaxationcalculateopenqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresstaxationcalculateopenqueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.taxation.calculate.open.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresstaxationcalculateopenqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresstaxationcalculateopenqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 主订单id
func (r *AliexpresstaxationcalculateopenqueryAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AliexpresstaxationcalculateopenqueryAPIRequest) GetOrderId() string {
	return r._orderId
}
