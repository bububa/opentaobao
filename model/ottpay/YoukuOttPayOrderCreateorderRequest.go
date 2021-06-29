package ottpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建订单 API请求
youku.ott.pay.order.createorder

ottpay创建订单
*/
type YoukuOttPayOrderCreateorderRequest struct {
    model.Params
    // 下单账号， cp账号
    buyer   string
    // 商品id
    productId   string
    // 商品名称
    productName   string
    // cp订单号
    orderNo   string
    // 价格， 单位：分
    price   string
    // 回调接口
    callbackUrl   string
    // 订单无关的其他参数,如埋点统计的utdid, mac地址等
    extra   string
    // 订单类型，1为连续包月类型
    orderType   int64
    // 连续包月实际参数
    realPrice   string
}

// 初始化YoukuOttPayOrderCreateorderRequest对象
func NewYoukuOttPayOrderCreateorderRequest() *YoukuOttPayOrderCreateorderRequest{
    return &YoukuOttPayOrderCreateorderRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuOttPayOrderCreateorderRequest) GetApiMethodName() string {
    return "youku.ott.pay.order.createorder"
}

// IRequest interface 方法, 获取API参数
func (r YoukuOttPayOrderCreateorderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Buyer Setter
// 下单账号， cp账号
func (r *YoukuOttPayOrderCreateorderRequest) SetBuyer(buyer string) error {
    r.buyer = buyer
    r.Set("buyer", buyer)
    return nil
}

// Buyer Getter
func (r YoukuOttPayOrderCreateorderRequest) GetBuyer() string {
    return r.buyer
}
// ProductId Setter
// 商品id
func (r *YoukuOttPayOrderCreateorderRequest) SetProductId(productId string) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r YoukuOttPayOrderCreateorderRequest) GetProductId() string {
    return r.productId
}
// ProductName Setter
// 商品名称
func (r *YoukuOttPayOrderCreateorderRequest) SetProductName(productName string) error {
    r.productName = productName
    r.Set("product_name", productName)
    return nil
}

// ProductName Getter
func (r YoukuOttPayOrderCreateorderRequest) GetProductName() string {
    return r.productName
}
// OrderNo Setter
// cp订单号
func (r *YoukuOttPayOrderCreateorderRequest) SetOrderNo(orderNo string) error {
    r.orderNo = orderNo
    r.Set("order_no", orderNo)
    return nil
}

// OrderNo Getter
func (r YoukuOttPayOrderCreateorderRequest) GetOrderNo() string {
    return r.orderNo
}
// Price Setter
// 价格， 单位：分
func (r *YoukuOttPayOrderCreateorderRequest) SetPrice(price string) error {
    r.price = price
    r.Set("price", price)
    return nil
}

// Price Getter
func (r YoukuOttPayOrderCreateorderRequest) GetPrice() string {
    return r.price
}
// CallbackUrl Setter
// 回调接口
func (r *YoukuOttPayOrderCreateorderRequest) SetCallbackUrl(callbackUrl string) error {
    r.callbackUrl = callbackUrl
    r.Set("callback_url", callbackUrl)
    return nil
}

// CallbackUrl Getter
func (r YoukuOttPayOrderCreateorderRequest) GetCallbackUrl() string {
    return r.callbackUrl
}
// Extra Setter
// 订单无关的其他参数,如埋点统计的utdid, mac地址等
func (r *YoukuOttPayOrderCreateorderRequest) SetExtra(extra string) error {
    r.extra = extra
    r.Set("extra", extra)
    return nil
}

// Extra Getter
func (r YoukuOttPayOrderCreateorderRequest) GetExtra() string {
    return r.extra
}
// OrderType Setter
// 订单类型，1为连续包月类型
func (r *YoukuOttPayOrderCreateorderRequest) SetOrderType(orderType int64) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

// OrderType Getter
func (r YoukuOttPayOrderCreateorderRequest) GetOrderType() int64 {
    return r.orderType
}
// RealPrice Setter
// 连续包月实际参数
func (r *YoukuOttPayOrderCreateorderRequest) SetRealPrice(realPrice string) error {
    r.realPrice = realPrice
    r.Set("real_price", realPrice)
    return nil
}

// RealPrice Getter
func (r YoukuOttPayOrderCreateorderRequest) GetRealPrice() string {
    return r.realPrice
}
