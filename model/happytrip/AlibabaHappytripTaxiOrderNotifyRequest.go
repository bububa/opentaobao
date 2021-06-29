package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
状态通知 API请求
alibaba.happytrip.taxi.order.notify

当订单发生变化是供应商通过状态通知API通知欢行，欢行获取最新的订单详情和状态进行业务处理。
*/
type AlibabaHappytripTaxiOrderNotifyRequest struct {
    model.Params
    // 返回自 1970 年 1 月 1 日 00:00:00 GMT 以来此 Date 对象表示的毫秒数
    _time   int64
    // 通知类型: 1-订单中间状态流转 2-订单终态通知 3-支付确认通知 4-订单退款通知 5-订单改价通知 6-客服关单通知。参考：https://open-hatrip.alibaba.com/doc/car/order_status_callback.html
    _notifyType   int64
    // 通知说明
    _notifyDesc   string
    // 订单id
    _orderId   string
}

// 初始化AlibabaHappytripTaxiOrderNotifyRequest对象
func NewAlibabaHappytripTaxiOrderNotifyRequest() *AlibabaHappytripTaxiOrderNotifyRequest{
    return &AlibabaHappytripTaxiOrderNotifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiOrderNotifyRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.order.notify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiOrderNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Time Setter
// 返回自 1970 年 1 月 1 日 00:00:00 GMT 以来此 Date 对象表示的毫秒数
func (r *AlibabaHappytripTaxiOrderNotifyRequest) SetTime(_time int64) error {
    r._time = _time
    r.Set("time", _time)
    return nil
}

// Time Getter
func (r AlibabaHappytripTaxiOrderNotifyRequest) GetTime() int64 {
    return r._time
}
// NotifyType Setter
// 通知类型: 1-订单中间状态流转 2-订单终态通知 3-支付确认通知 4-订单退款通知 5-订单改价通知 6-客服关单通知。参考：https://open-hatrip.alibaba.com/doc/car/order_status_callback.html
func (r *AlibabaHappytripTaxiOrderNotifyRequest) SetNotifyType(_notifyType int64) error {
    r._notifyType = _notifyType
    r.Set("notify_type", _notifyType)
    return nil
}

// NotifyType Getter
func (r AlibabaHappytripTaxiOrderNotifyRequest) GetNotifyType() int64 {
    return r._notifyType
}
// NotifyDesc Setter
// 通知说明
func (r *AlibabaHappytripTaxiOrderNotifyRequest) SetNotifyDesc(_notifyDesc string) error {
    r._notifyDesc = _notifyDesc
    r.Set("notify_desc", _notifyDesc)
    return nil
}

// NotifyDesc Getter
func (r AlibabaHappytripTaxiOrderNotifyRequest) GetNotifyDesc() string {
    return r._notifyDesc
}
// OrderId Setter
// 订单id
func (r *AlibabaHappytripTaxiOrderNotifyRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHappytripTaxiOrderNotifyRequest) GetOrderId() string {
    return r._orderId
}
