package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptransferorderdetailAPIRequest 接送订单详情接口 API请求
// alitrip.transfer.order.detail
//
// 接送订单详情接口
type AlitriptransferorderdetailAPIRequest struct {
	model.Params
	// 飞猪平台订单id
	_orderId string
	// 分配的服务商Id
	_providerId string
}

// NewAlitriptransferorderdetailRequest 初始化AlitriptransferorderdetailAPIRequest对象
func NewAlitriptransferorderdetailRequest() *AlitriptransferorderdetailAPIRequest {
	return &AlitriptransferorderdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptransferorderdetailAPIRequest) GetApiMethodName() string {
	return "alitrip.transfer.order.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptransferorderdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptransferorderdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 飞猪平台订单id
func (r *AlitriptransferorderdetailAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitriptransferorderdetailAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetProviderId is ProviderId Setter
// 分配的服务商Id
func (r *AlitriptransferorderdetailAPIRequest) SetProviderId(_providerId string) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r AlitriptransferorderdetailAPIRequest) GetProviderId() string {
	return r._providerId
}
