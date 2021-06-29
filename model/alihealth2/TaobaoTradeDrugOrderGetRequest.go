package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查看订单详情 APIRequest
taobao.trade.drug.order.get

商家查看订单详情
*/
type TaobaoTradeDrugOrderGetRequest struct {
    model.Params

    // 订单id
    orderId   int64 

}

func NewTaobaoTradeDrugOrderGetRequest() *TaobaoTradeDrugOrderGetRequest{
    return &TaobaoTradeDrugOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTradeDrugOrderGetRequest) GetApiMethodName() string {
    return "taobao.trade.drug.order.get"
}

func (r TaobaoTradeDrugOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTradeDrugOrderGetRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TaobaoTradeDrugOrderGetRequest) GetOrderId() int64 {
    return r.orderId
}

