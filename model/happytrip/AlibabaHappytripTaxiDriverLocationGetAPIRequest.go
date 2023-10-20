package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahappytriptaxidriverlocationgetAPIRequest 司机位置 API请求
// alibaba.happytrip.taxi.driver.location.get
//
// 获取司机实时位置
type AlibabahappytriptaxidriverlocationgetAPIRequest struct {
	model.Params
	// 供应商订单号
	_orderId string
}

// NewAlibabahappytriptaxidriverlocationgetRequest 初始化AlibabahappytriptaxidriverlocationgetAPIRequest对象
func NewAlibabahappytriptaxidriverlocationgetRequest() *AlibabahappytriptaxidriverlocationgetAPIRequest {
	return &AlibabahappytriptaxidriverlocationgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahappytriptaxidriverlocationgetAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.driver.location.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahappytriptaxidriverlocationgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahappytriptaxidriverlocationgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 供应商订单号
func (r *AlibabahappytriptaxidriverlocationgetAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabahappytriptaxidriverlocationgetAPIRequest) GetOrderId() string {
	return r._orderId
}
