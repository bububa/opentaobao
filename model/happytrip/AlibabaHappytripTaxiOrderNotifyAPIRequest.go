package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahappytriptaxiordernotifyAPIRequest 状态通知 API请求
// alibaba.happytrip.taxi.order.notify
//
// 当订单发生变化是供应商通过状态通知API通知欢行，欢行获取最新的订单详情和状态进行业务处理。
type AlibabahappytriptaxiordernotifyAPIRequest struct {
	model.Params
	// 通知说明
	_notifyDesc string
	// 订单id
	_orderId string
	// 工单id（通知投诉处理完成时传递）
	_caseId string
	// 返回自 1970 年 1 月 1 日 00:00:00 GMT 以来此 Date 对象表示的毫秒数
	_time int64
	// 通知类型: 1-订单中间状态流转 2-订单终态通知 3-支付确认通知 4-订单退款通知 5-订单改价通知 6-客服关单通知 7-投诉处理通知。
	_notifyType int64
}

// NewAlibabahappytriptaxiordernotifyRequest 初始化AlibabahappytriptaxiordernotifyAPIRequest对象
func NewAlibabahappytriptaxiordernotifyRequest() *AlibabahappytriptaxiordernotifyAPIRequest {
	return &AlibabahappytriptaxiordernotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahappytriptaxiordernotifyAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.order.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahappytriptaxiordernotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahappytriptaxiordernotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNotifyDesc is NotifyDesc Setter
// 通知说明
func (r *AlibabahappytriptaxiordernotifyAPIRequest) SetNotifyDesc(_notifyDesc string) error {
	r._notifyDesc = _notifyDesc
	r.Set("notify_desc", _notifyDesc)
	return nil
}

// GetNotifyDesc NotifyDesc Getter
func (r AlibabahappytriptaxiordernotifyAPIRequest) GetNotifyDesc() string {
	return r._notifyDesc
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabahappytriptaxiordernotifyAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabahappytriptaxiordernotifyAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetCaseId is CaseId Setter
// 工单id（通知投诉处理完成时传递）
func (r *AlibabahappytriptaxiordernotifyAPIRequest) SetCaseId(_caseId string) error {
	r._caseId = _caseId
	r.Set("case_id", _caseId)
	return nil
}

// GetCaseId CaseId Getter
func (r AlibabahappytriptaxiordernotifyAPIRequest) GetCaseId() string {
	return r._caseId
}

// SetTime is Time Setter
// 返回自 1970 年 1 月 1 日 00:00:00 GMT 以来此 Date 对象表示的毫秒数
func (r *AlibabahappytriptaxiordernotifyAPIRequest) SetTime(_time int64) error {
	r._time = _time
	r.Set("time", _time)
	return nil
}

// GetTime Time Getter
func (r AlibabahappytriptaxiordernotifyAPIRequest) GetTime() int64 {
	return r._time
}

// SetNotifyType is NotifyType Setter
// 通知类型: 1-订单中间状态流转 2-订单终态通知 3-支付确认通知 4-订单退款通知 5-订单改价通知 6-客服关单通知 7-投诉处理通知。
func (r *AlibabahappytriptaxiordernotifyAPIRequest) SetNotifyType(_notifyType int64) error {
	r._notifyType = _notifyType
	r.Set("notify_type", _notifyType)
	return nil
}

// GetNotifyType NotifyType Getter
func (r AlibabahappytriptaxiordernotifyAPIRequest) GetNotifyType() int64 {
	return r._notifyType
}
