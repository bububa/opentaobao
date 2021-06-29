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
type TaobaoVmarketEticketStoreGetRequest struct {
    model.Params
    // 订单ID
    _orderId   int64
}

// 初始化TaobaoVmarketEticketStoreGetRequest对象
func NewTaobaoVmarketEticketStoreGetRequest() *TaobaoVmarketEticketStoreGetRequest{
    return &TaobaoVmarketEticketStoreGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketStoreGetRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.store.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketStoreGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单ID
func (r *TaobaoVmarketEticketStoreGetRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoVmarketEticketStoreGetRequest) GetOrderId() int64 {
    return r._orderId
}
