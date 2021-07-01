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
type TmallCarFpcarGetcarNotifyAPIRequest struct {
    model.Params
    // 商品宝贝id
    _itemId   int64
    // 订单id
    _orderId   int64
    // 卖家id
    _sellerId   int64
}

// 初始化TmallCarFpcarGetcarNotifyAPIRequest对象
func NewTmallCarFpcarGetcarNotifyRequest() *TmallCarFpcarGetcarNotifyAPIRequest{
    return &TmallCarFpcarGetcarNotifyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarFpcarGetcarNotifyAPIRequest) GetApiMethodName() string {
    return "tmall.car.fpcar.getcar.notify"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarFpcarGetcarNotifyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品宝贝id
func (r *TmallCarFpcarGetcarNotifyAPIRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TmallCarFpcarGetcarNotifyAPIRequest) GetItemId() int64 {
    return r._itemId
}
// OrderId Setter
// 订单id
func (r *TmallCarFpcarGetcarNotifyAPIRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TmallCarFpcarGetcarNotifyAPIRequest) GetOrderId() int64 {
    return r._orderId
}
// SellerId Setter
// 卖家id
func (r *TmallCarFpcarGetcarNotifyAPIRequest) SetSellerId(_sellerId int64) error {
    r._sellerId = _sellerId
    r.Set("seller_id", _sellerId)
    return nil
}

// SellerId Getter
func (r TmallCarFpcarGetcarNotifyAPIRequest) GetSellerId() int64 {
    return r._sellerId
}
