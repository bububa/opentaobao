package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取电子凭证预约门店信息 APIRequest
taobao.vmarket.eticket.store.get

用于给外部商家查询电子凭证预约门店信息
*/
type TaobaoVmarketEticketStoreGetRequest struct {
    model.Params

    // 订单ID
    orderId   int64 

}

func NewTaobaoVmarketEticketStoreGetRequest() *TaobaoVmarketEticketStoreGetRequest{
    return &TaobaoVmarketEticketStoreGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoVmarketEticketStoreGetRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.store.get"
}

func (r TaobaoVmarketEticketStoreGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoVmarketEticketStoreGetRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TaobaoVmarketEticketStoreGetRequest) GetOrderId() int64 {
    return r.orderId
}

