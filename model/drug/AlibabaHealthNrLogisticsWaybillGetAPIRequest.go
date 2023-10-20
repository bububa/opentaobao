package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahealthnrlogisticswaybillgetAPIRequest 电子面单查询接口 API请求
// alibaba.health.nr.logistics.waybill.get
//
// 商家登录后根据订单号查询物流单号及电子面单信息
type AlibabahealthnrlogisticswaybillgetAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
}

// NewAlibabahealthnrlogisticswaybillgetRequest 初始化AlibabahealthnrlogisticswaybillgetAPIRequest对象
func NewAlibabahealthnrlogisticswaybillgetRequest() *AlibabahealthnrlogisticswaybillgetAPIRequest {
	return &AlibabahealthnrlogisticswaybillgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahealthnrlogisticswaybillgetAPIRequest) GetApiMethodName() string {
	return "alibaba.health.nr.logistics.waybill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahealthnrlogisticswaybillgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahealthnrlogisticswaybillgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabahealthnrlogisticswaybillgetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabahealthnrlogisticswaybillgetAPIRequest) GetOrderId() int64 {
	return r._orderId
}
