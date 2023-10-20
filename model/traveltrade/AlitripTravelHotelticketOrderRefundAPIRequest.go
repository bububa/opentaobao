package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptravelhotelticketorderrefundAPIRequest 退款结结果通知 API请求
// alitrip.travel.hotelticket.order.refund
//
// 退款结果通知
type AlitriptravelhotelticketorderrefundAPIRequest struct {
	model.Params
	// 系统商订单号
	_orderId string
	// 退款失败原因
	_failMsg string
	// 飞猪订单号
	_fliggyOrderId string
	// 退款结果状态 1:退款成功  2:退款失败
	_status int64
}

// NewAlitriptravelhotelticketorderrefundRequest 初始化AlitriptravelhotelticketorderrefundAPIRequest对象
func NewAlitriptravelhotelticketorderrefundRequest() *AlitriptravelhotelticketorderrefundAPIRequest {
	return &AlitriptravelhotelticketorderrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptravelhotelticketorderrefundAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.hotelticket.order.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptravelhotelticketorderrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptravelhotelticketorderrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 系统商订单号
func (r *AlitriptravelhotelticketorderrefundAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitriptravelhotelticketorderrefundAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetFailMsg is FailMsg Setter
// 退款失败原因
func (r *AlitriptravelhotelticketorderrefundAPIRequest) SetFailMsg(_failMsg string) error {
	r._failMsg = _failMsg
	r.Set("fail_msg", _failMsg)
	return nil
}

// GetFailMsg FailMsg Getter
func (r AlitriptravelhotelticketorderrefundAPIRequest) GetFailMsg() string {
	return r._failMsg
}

// SetFliggyOrderId is FliggyOrderId Setter
// 飞猪订单号
func (r *AlitriptravelhotelticketorderrefundAPIRequest) SetFliggyOrderId(_fliggyOrderId string) error {
	r._fliggyOrderId = _fliggyOrderId
	r.Set("fliggy_order_id", _fliggyOrderId)
	return nil
}

// GetFliggyOrderId FliggyOrderId Getter
func (r AlitriptravelhotelticketorderrefundAPIRequest) GetFliggyOrderId() string {
	return r._fliggyOrderId
}

// SetStatus is Status Setter
// 退款结果状态 1:退款成功  2:退款失败
func (r *AlitriptravelhotelticketorderrefundAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlitriptravelhotelticketorderrefundAPIRequest) GetStatus() int64 {
	return r._status
}
