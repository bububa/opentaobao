package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店线下已收尾款 API请求
tmall.car.fpcar.restpay.receive

提供给外部(大搜或其它合作方)的接口-门店线下已收尾款(不执行分佣)
*/
type TmallCarFpcarRestpayReceiveRequest struct {
    model.Params
    // 卖家id
    _sellerId   int64
    // 订单id
    _orderId   int64
    // 商品宝贝id
    _itemId   int64
}

// 初始化TmallCarFpcarRestpayReceiveRequest对象
func NewTmallCarFpcarRestpayReceiveRequest() *TmallCarFpcarRestpayReceiveRequest{
    return &TmallCarFpcarRestpayReceiveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarFpcarRestpayReceiveRequest) GetApiMethodName() string {
    return "tmall.car.fpcar.restpay.receive"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarFpcarRestpayReceiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SellerId Setter
// 卖家id
func (r *TmallCarFpcarRestpayReceiveRequest) SetSellerId(_sellerId int64) error {
    r._sellerId = _sellerId
    r.Set("seller_id", _sellerId)
    return nil
}

// SellerId Getter
func (r TmallCarFpcarRestpayReceiveRequest) GetSellerId() int64 {
    return r._sellerId
}
// OrderId Setter
// 订单id
func (r *TmallCarFpcarRestpayReceiveRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TmallCarFpcarRestpayReceiveRequest) GetOrderId() int64 {
    return r._orderId
}
// ItemId Setter
// 商品宝贝id
func (r *TmallCarFpcarRestpayReceiveRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TmallCarFpcarRestpayReceiveRequest) GetItemId() int64 {
    return r._itemId
}
