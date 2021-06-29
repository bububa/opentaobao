package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
主动发起通知接口 API请求
taobao.vmarket.eticket.manage.notify

外部合作商家主动发起通知接口
*/
type TaobaoVmarketEticketManageNotifyRequest struct {
    model.Params
    // 订单编号
    _orderId   int64
    // 码商ID，如果是码商，必须传，如果是信任卖家，不需要传
    _codemerchantId   int64
    // 需要调用的通知方法，目前仅支持是send（发码）或resend（重新发码）
    _notifyMethod   string
}

// 初始化TaobaoVmarketEticketManageNotifyRequest对象
func NewTaobaoVmarketEticketManageNotifyRequest() *TaobaoVmarketEticketManageNotifyRequest{
    return &TaobaoVmarketEticketManageNotifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketManageNotifyRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.manage.notify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketManageNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单编号
func (r *TaobaoVmarketEticketManageNotifyRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoVmarketEticketManageNotifyRequest) GetOrderId() int64 {
    return r._orderId
}
// CodemerchantId Setter
// 码商ID，如果是码商，必须传，如果是信任卖家，不需要传
func (r *TaobaoVmarketEticketManageNotifyRequest) SetCodemerchantId(_codemerchantId int64) error {
    r._codemerchantId = _codemerchantId
    r.Set("codemerchant_id", _codemerchantId)
    return nil
}

// CodemerchantId Getter
func (r TaobaoVmarketEticketManageNotifyRequest) GetCodemerchantId() int64 {
    return r._codemerchantId
}
// NotifyMethod Setter
// 需要调用的通知方法，目前仅支持是send（发码）或resend（重新发码）
func (r *TaobaoVmarketEticketManageNotifyRequest) SetNotifyMethod(_notifyMethod string) error {
    r._notifyMethod = _notifyMethod
    r.Set("notify_method", _notifyMethod)
    return nil
}

// NotifyMethod Getter
func (r TaobaoVmarketEticketManageNotifyRequest) GetNotifyMethod() string {
    return r._notifyMethod
}
