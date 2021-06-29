package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
整车租车回传预约信息 API请求
tmall.car.lease.reserve

租赁公司回传预约到店信息
*/
type TmallCarLeaseReserveRequest struct {
    model.Params
    // 买家id
    _buyerId   int64
    // 订单id
    _orderId   int64
    // 文案
    _text   string
    // 车架号
    _vin   string
    // 1 代表 车辆到店，已预约用户到店提车   ; 2 车辆到店，未能联系到用户
    _flag   int64
    // 买家昵称
    _buyerNick   string
}

// 初始化TmallCarLeaseReserveRequest对象
func NewTmallCarLeaseReserveRequest() *TmallCarLeaseReserveRequest{
    return &TmallCarLeaseReserveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarLeaseReserveRequest) GetApiMethodName() string {
    return "tmall.car.lease.reserve"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarLeaseReserveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BuyerId Setter
// 买家id
func (r *TmallCarLeaseReserveRequest) SetBuyerId(_buyerId int64) error {
    r._buyerId = _buyerId
    r.Set("buyer_id", _buyerId)
    return nil
}

// BuyerId Getter
func (r TmallCarLeaseReserveRequest) GetBuyerId() int64 {
    return r._buyerId
}
// OrderId Setter
// 订单id
func (r *TmallCarLeaseReserveRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TmallCarLeaseReserveRequest) GetOrderId() int64 {
    return r._orderId
}
// Text Setter
// 文案
func (r *TmallCarLeaseReserveRequest) SetText(_text string) error {
    r._text = _text
    r.Set("text", _text)
    return nil
}

// Text Getter
func (r TmallCarLeaseReserveRequest) GetText() string {
    return r._text
}
// Vin Setter
// 车架号
func (r *TmallCarLeaseReserveRequest) SetVin(_vin string) error {
    r._vin = _vin
    r.Set("vin", _vin)
    return nil
}

// Vin Getter
func (r TmallCarLeaseReserveRequest) GetVin() string {
    return r._vin
}
// Flag Setter
// 1 代表 车辆到店，已预约用户到店提车   ; 2 车辆到店，未能联系到用户
func (r *TmallCarLeaseReserveRequest) SetFlag(_flag int64) error {
    r._flag = _flag
    r.Set("flag", _flag)
    return nil
}

// Flag Getter
func (r TmallCarLeaseReserveRequest) GetFlag() int64 {
    return r._flag
}
// BuyerNick Setter
// 买家昵称
func (r *TmallCarLeaseReserveRequest) SetBuyerNick(_buyerNick string) error {
    r._buyerNick = _buyerNick
    r.Set("buyer_nick", _buyerNick)
    return nil
}

// BuyerNick Getter
func (r TmallCarLeaseReserveRequest) GetBuyerNick() string {
    return r._buyerNick
}
