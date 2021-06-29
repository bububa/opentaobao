package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单延时接口 API请求
taobao.vmarket.eticket.time.expand

提供码商操作订单延期接口
*/
type TaobaoVmarketEticketTimeExpandRequest struct {
    model.Params
    // 订单ID
    orderId   int64
    // 延长天数，延长时间=当前过期时间+延长天数
    expandDays   int64
}

// 初始化TaobaoVmarketEticketTimeExpandRequest对象
func NewTaobaoVmarketEticketTimeExpandRequest() *TaobaoVmarketEticketTimeExpandRequest{
    return &TaobaoVmarketEticketTimeExpandRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketTimeExpandRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.time.expand"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketTimeExpandRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单ID
func (r *TaobaoVmarketEticketTimeExpandRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TaobaoVmarketEticketTimeExpandRequest) GetOrderId() int64 {
    return r.orderId
}
// ExpandDays Setter
// 延长天数，延长时间=当前过期时间+延长天数
func (r *TaobaoVmarketEticketTimeExpandRequest) SetExpandDays(expandDays int64) error {
    r.expandDays = expandDays
    r.Set("expand_days", expandDays)
    return nil
}

// ExpandDays Getter
func (r TaobaoVmarketEticketTimeExpandRequest) GetExpandDays() int64 {
    return r.expandDays
}
