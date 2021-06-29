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
    _appkey   string
    // 订单号，与外部订单号二选一
    _orderId   int64
    // 外部订单号，使用该参数完成查询订单等操作，与外部订单号二选一
    _outOrderId   string
    // 设备类型，目前只支持PC，可选
    _deviceType   string
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
func (r *TaobaoFuwuPurchaseOrderPayRequest) SetAppkey(_appkey string) error {
    r._appkey = _appkey
    r.Set("appkey", _appkey)
    return nil
}

// Appkey Getter
func (r TaobaoFuwuPurchaseOrderPayRequest) GetAppkey() string {
    return r._appkey
}
// OrderId Setter
// 订单号，与外部订单号二选一
func (r *TaobaoFuwuPurchaseOrderPayRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoFuwuPurchaseOrderPayRequest) GetOrderId() int64 {
    return r._orderId
}
// OutOrderId Setter
// 外部订单号，使用该参数完成查询订单等操作，与外部订单号二选一
func (r *TaobaoFuwuPurchaseOrderPayRequest) SetOutOrderId(_outOrderId string) error {
    r._outOrderId = _outOrderId
    r.Set("out_order_id", _outOrderId)
    return nil
}

// OutOrderId Getter
func (r TaobaoFuwuPurchaseOrderPayRequest) GetOutOrderId() string {
    return r._outOrderId
}
// DeviceType Setter
// 设备类型，目前只支持PC，可选
func (r *TaobaoFuwuPurchaseOrderPayRequest) SetDeviceType(_deviceType string) error {
    r._deviceType = _deviceType
    r.Set("device_type", _deviceType)
    return nil
}

// DeviceType Getter
func (r TaobaoFuwuPurchaseOrderPayRequest) GetDeviceType() string {
    return r._deviceType
}
