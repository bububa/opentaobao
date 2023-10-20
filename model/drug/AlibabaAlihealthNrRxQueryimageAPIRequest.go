package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthnrrxqueryimageAPIRequest o2o查看处方图片 API请求
// alibaba.alihealth.nr.rx.queryimage
//
// o2o商家查看处方图片，包括电子图片与纸质图片
type AlibabaalihealthnrrxqueryimageAPIRequest struct {
	model.Params
	// 订单编号
	_orderId int64
}

// NewAlibabaalihealthnrrxqueryimageRequest 初始化AlibabaalihealthnrrxqueryimageAPIRequest对象
func NewAlibabaalihealthnrrxqueryimageRequest() *AlibabaalihealthnrrxqueryimageAPIRequest {
	return &AlibabaalihealthnrrxqueryimageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthnrrxqueryimageAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.nr.rx.queryimage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthnrrxqueryimageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthnrrxqueryimageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单编号
func (r *AlibabaalihealthnrrxqueryimageAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaalihealthnrrxqueryimageAPIRequest) GetOrderId() int64 {
	return r._orderId
}
