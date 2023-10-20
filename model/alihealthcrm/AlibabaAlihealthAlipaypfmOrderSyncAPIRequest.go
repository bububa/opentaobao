package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthalipaypfmordersyncAPIRequest 订单数据回传接口 API请求
// alibaba.alihealth.alipaypfm.order.sync
//
// 订单数据回传接口，各个isv通过我们渠道产生订单需要回传进行统计
type AlibabaalihealthalipaypfmordersyncAPIRequest struct {
	model.Params
	// 订单id
	_orderId string
	// 订单价格
	_orderPrice string
	// 订单状态
	_orderStatus string
	// 扩展参数
	_extParam string
	// user_id
	_userId int64
}

// NewAlibabaalihealthalipaypfmordersyncRequest 初始化AlibabaalihealthalipaypfmordersyncAPIRequest对象
func NewAlibabaalihealthalipaypfmordersyncRequest() *AlibabaalihealthalipaypfmordersyncAPIRequest {
	return &AlibabaalihealthalipaypfmordersyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthalipaypfmordersyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.alipaypfm.order.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthalipaypfmordersyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthalipaypfmordersyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabaalihealthalipaypfmordersyncAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaalihealthalipaypfmordersyncAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetOrderPrice is OrderPrice Setter
// 订单价格
func (r *AlibabaalihealthalipaypfmordersyncAPIRequest) SetOrderPrice(_orderPrice string) error {
	r._orderPrice = _orderPrice
	r.Set("order_price", _orderPrice)
	return nil
}

// GetOrderPrice OrderPrice Getter
func (r AlibabaalihealthalipaypfmordersyncAPIRequest) GetOrderPrice() string {
	return r._orderPrice
}

// SetOrderStatus is OrderStatus Setter
// 订单状态
func (r *AlibabaalihealthalipaypfmordersyncAPIRequest) SetOrderStatus(_orderStatus string) error {
	r._orderStatus = _orderStatus
	r.Set("order_status", _orderStatus)
	return nil
}

// GetOrderStatus OrderStatus Getter
func (r AlibabaalihealthalipaypfmordersyncAPIRequest) GetOrderStatus() string {
	return r._orderStatus
}

// SetExtParam is ExtParam Setter
// 扩展参数
func (r *AlibabaalihealthalipaypfmordersyncAPIRequest) SetExtParam(_extParam string) error {
	r._extParam = _extParam
	r.Set("ext_param", _extParam)
	return nil
}

// GetExtParam ExtParam Getter
func (r AlibabaalihealthalipaypfmordersyncAPIRequest) GetExtParam() string {
	return r._extParam
}

// SetUserId is UserId Setter
// user_id
func (r *AlibabaalihealthalipaypfmordersyncAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaalihealthalipaypfmordersyncAPIRequest) GetUserId() int64 {
	return r._userId
}
