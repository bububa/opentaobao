package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店通知用户提车 API请求
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

// 初始化TmallCarFpcarGetcarNotifyRequest对象
func NewTmallCarFpcarGetcarNotifyRequest() *TmallCarFpcarGetcarNotifyRequest{
    return &TmallCarFpcarGetcarNotifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarFpcarGetcarNotifyRequest) GetApiMethodName() string {
    return "tmall.car.fpcar.getcar.notify"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarFpcarGetcarNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品宝贝id
func (r *TmallCarFpcarGetcarNotifyRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TmallCarFpcarGetcarNotifyRequest) GetItemId() int64 {
    return r.itemId
}
// OrderId Setter
// 订单id
func (r *TmallCarFpcarGetcarNotifyRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TmallCarFpcarGetcarNotifyRequest) GetOrderId() int64 {
    return r.orderId
}
// SellerId Setter
// 卖家id
func (r *TmallCarFpcarGetcarNotifyRequest) SetSellerId(sellerId int64) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

// SellerId Getter
func (r TmallCarFpcarGetcarNotifyRequest) GetSellerId() int64 {
    return r.sellerId
}
