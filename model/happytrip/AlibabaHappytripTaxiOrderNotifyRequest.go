package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
状态通知 APIRequest
alibaba.happytrip.taxi.order.notify

当订单发生变化是供应商通过状态通知API通知欢行，欢行获取最新的订单详情和状态进行业务处理。
*/
type AlibabaHappytripTaxiOrderNotifyRequest struct {
    model.Params

    // 返回自 1970 年 1 月 1 日 00:00:00 GMT 以来此 Date 对象表示的毫秒数
    time   int64 

    // 通知类型: 1-订单中间状态流转 2-订单终态通知 3-支付确认通知 4-订单退款通知 5-订单改价通知 6-客服关单通知。参考：https://open-hatrip.alibaba.com/doc/car/order_status_callback.html
    notifyType   int64 

    // 通知说明
    notifyDesc   string 

    // 订单id
    orderId   string 

}

func NewAlibabaHappytripTaxiOrderNotifyRequest() *AlibabaHappytripTaxiOrderNotifyRequest{
    return &AlibabaHappytripTaxiOrderNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHappytripTaxiOrderNotifyRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.order.notify"
}

func (r AlibabaHappytripTaxiOrderNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHappytripTaxiOrderNotifyRequest) SetTime(time int64) error {
    r.time = time
    r.Set("time", time)
    return nil
}

func (r AlibabaHappytripTaxiOrderNotifyRequest) GetTime() int64 {
    return r.time
}

func (r *AlibabaHappytripTaxiOrderNotifyRequest) SetNotifyType(notifyType int64) error {
    r.notifyType = notifyType
    r.Set("notify_type", notifyType)
    return nil
}

func (r AlibabaHappytripTaxiOrderNotifyRequest) GetNotifyType() int64 {
    return r.notifyType
}

func (r *AlibabaHappytripTaxiOrderNotifyRequest) SetNotifyDesc(notifyDesc string) error {
    r.notifyDesc = notifyDesc
    r.Set("notify_desc", notifyDesc)
    return nil
}

func (r AlibabaHappytripTaxiOrderNotifyRequest) GetNotifyDesc() string {
    return r.notifyDesc
}

func (r *AlibabaHappytripTaxiOrderNotifyRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaHappytripTaxiOrderNotifyRequest) GetOrderId() string {
    return r.orderId
}

