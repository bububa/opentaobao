package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取电子凭证预约门店信息 API请求
taobao.vmarket.eticket.store.get

用于给外部商家查询电子凭证预约门店信息
*/
type TaobaoVmarketEticketStoreGetAPIRequest struct {
    model.Params
    // 订单ID
    _orderId   int64
}

// 初始化TaobaoVmarketEticketStoreGetAPIRequest对象
func NewTaobaoVmarketEticketStoreGetRequest() *TaobaoVmarketEticketStoreGetAPIRequest{
    return &TaobaoVmarketEticketStoreGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketStoreGetAPIRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.store.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketStoreGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单ID
func (r *TaobaoVmarketEticketStoreGetAPIRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoVmarketEticketStoreGetAPIRequest) GetOrderId() int64 {
    return r._orderId
}
