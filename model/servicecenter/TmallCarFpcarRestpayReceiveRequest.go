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
    sellerId   int64
    // 订单id
    orderId   int64
    // 商品宝贝id
    itemId   int64
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
func (r *TmallCarFpcarRestpayReceiveRequest) SetSellerId(sellerId int64) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

// SellerId Getter
func (r TmallCarFpcarRestpayReceiveRequest) GetSellerId() int64 {
    return r.sellerId
}
// OrderId Setter
// 订单id
func (r *TmallCarFpcarRestpayReceiveRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TmallCarFpcarRestpayReceiveRequest) GetOrderId() int64 {
    return r.orderId
}
// ItemId Setter
// 商品宝贝id
func (r *TmallCarFpcarRestpayReceiveRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TmallCarFpcarRestpayReceiveRequest) GetItemId() int64 {
    return r.itemId
}
