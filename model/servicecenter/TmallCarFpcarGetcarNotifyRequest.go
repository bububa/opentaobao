package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店通知用户提车 APIRequest
tmall.car.fpcar.getcar.notify

提供给外部(大搜或其它合作方)的接口-门店通知用户提车
*/
type TmallCarFpcarGetcarNotifyRequest struct {
    model.Params

    // 商品宝贝id
    itemId   int64 

    // 订单id
    orderId   int64 

    // 卖家id
    sellerId   int64 

}

func NewTmallCarFpcarGetcarNotifyRequest() *TmallCarFpcarGetcarNotifyRequest{
    return &TmallCarFpcarGetcarNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCarFpcarGetcarNotifyRequest) GetApiMethodName() string {
    return "tmall.car.fpcar.getcar.notify"
}

func (r TmallCarFpcarGetcarNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCarFpcarGetcarNotifyRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TmallCarFpcarGetcarNotifyRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TmallCarFpcarGetcarNotifyRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TmallCarFpcarGetcarNotifyRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *TmallCarFpcarGetcarNotifyRequest) SetSellerId(sellerId int64) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

func (r TmallCarFpcarGetcarNotifyRequest) GetSellerId() int64 {
    return r.sellerId
}

