package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
内购服务订单付款页获取接口 API请求
taobao.fuwu.purchase.order.pay

通过接口获取某一订单的付款页面链接
*/
type TaobaoFuwuPurchaseOrderPayRequest struct {
    model.Params
    // APPKEY，必填
    appkey   string
    // 订单号，与外部订单号二选一
    orderId   int64
    // 外部订单号，使用该参数完成查询订单等操作，与外部订单号二选一
    outOrderId   string
    // 设备类型，目前只支持PC，可选
    deviceType   string
}

// 初始化TaobaoFuwuPurchaseOrderPayRequest对象
func NewTaobaoFuwuPurchaseOrderPayRequest() *TaobaoFuwuPurchaseOrderPayRequest{
    return &TaobaoFuwuPurchaseOrderPayRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFuwuPurchaseOrderPayRequest) GetApiMethodName() string {
    return "taobao.fuwu.purchase.order.pay"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFuwuPurchaseOrderPayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Appkey Setter
// APPKEY，必填
func (r *TaobaoFuwuPurchaseOrderPayRequest) SetAppkey(appkey string) error {
    r.appkey = appkey
    r.Set("appkey", appkey)
    return nil
}

// Appkey Getter
func (r TaobaoFuwuPurchaseOrderPayRequest) GetAppkey() string {
    return r.appkey
}
// OrderId Setter
// 订单号，与外部订单号二选一
func (r *TaobaoFuwuPurchaseOrderPayRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TaobaoFuwuPurchaseOrderPayRequest) GetOrderId() int64 {
    return r.orderId
}
// OutOrderId Setter
// 外部订单号，使用该参数完成查询订单等操作，与外部订单号二选一
func (r *TaobaoFuwuPurchaseOrderPayRequest) SetOutOrderId(outOrderId string) error {
    r.outOrderId = outOrderId
    r.Set("out_order_id", outOrderId)
    return nil
}

// OutOrderId Getter
func (r TaobaoFuwuPurchaseOrderPayRequest) GetOutOrderId() string {
    return r.outOrderId
}
// DeviceType Setter
// 设备类型，目前只支持PC，可选
func (r *TaobaoFuwuPurchaseOrderPayRequest) SetDeviceType(deviceType string) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

// DeviceType Getter
func (r TaobaoFuwuPurchaseOrderPayRequest) GetDeviceType() string {
    return r.deviceType
}
