package ottpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建订单 APIRequest
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

func NewYoukuOttPayOrderCreateorderRequest() *YoukuOttPayOrderCreateorderRequest{
    return &YoukuOttPayOrderCreateorderRequest{
        Params: model.NewParams(),
    }
}

func (r YoukuOttPayOrderCreateorderRequest) GetApiMethodName() string {
    return "youku.ott.pay.order.createorder"
}

func (r YoukuOttPayOrderCreateorderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YoukuOttPayOrderCreateorderRequest) SetBuyer(buyer string) error {
    r.buyer = buyer
    r.Set("buyer", buyer)
    return nil
}

func (r YoukuOttPayOrderCreateorderRequest) GetBuyer() string {
    return r.buyer
}

func (r *YoukuOttPayOrderCreateorderRequest) SetProductId(productId string) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r YoukuOttPayOrderCreateorderRequest) GetProductId() string {
    return r.productId
}

func (r *YoukuOttPayOrderCreateorderRequest) SetProductName(productName string) error {
    r.productName = productName
    r.Set("product_name", productName)
    return nil
}

func (r YoukuOttPayOrderCreateorderRequest) GetProductName() string {
    return r.productName
}

func (r *YoukuOttPayOrderCreateorderRequest) SetOrderNo(orderNo string) error {
    r.orderNo = orderNo
    r.Set("order_no", orderNo)
    return nil
}

func (r YoukuOttPayOrderCreateorderRequest) GetOrderNo() string {
    return r.orderNo
}

func (r *YoukuOttPayOrderCreateorderRequest) SetPrice(price string) error {
    r.price = price
    r.Set("price", price)
    return nil
}

func (r YoukuOttPayOrderCreateorderRequest) GetPrice() string {
    return r.price
}

func (r *YoukuOttPayOrderCreateorderRequest) SetCallbackUrl(callbackUrl string) error {
    r.callbackUrl = callbackUrl
    r.Set("callback_url", callbackUrl)
    return nil
}

func (r YoukuOttPayOrderCreateorderRequest) GetCallbackUrl() string {
    return r.callbackUrl
}

func (r *YoukuOttPayOrderCreateorderRequest) SetExtra(extra string) error {
    r.extra = extra
    r.Set("extra", extra)
    return nil
}

func (r YoukuOttPayOrderCreateorderRequest) GetExtra() string {
    return r.extra
}

func (r *YoukuOttPayOrderCreateorderRequest) SetOrderType(orderType int64) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

func (r YoukuOttPayOrderCreateorderRequest) GetOrderType() int64 {
    return r.orderType
}

func (r *YoukuOttPayOrderCreateorderRequest) SetRealPrice(realPrice string) error {
    r.realPrice = realPrice
    r.Set("real_price", realPrice)
    return nil
}

func (r YoukuOttPayOrderCreateorderRequest) GetRealPrice() string {
    return r.realPrice
}

