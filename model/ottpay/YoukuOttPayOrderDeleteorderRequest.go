package ottpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退订应用中心支付订单 APIRequest
youku.ott.pay.order.deleteorder

应用中心sdk连续包月退订接口
*/
type YoukuOttPayOrderDeleteorderRequest struct {
    model.Params

    // 下单账号， cp账号
    buyer   string 

    // 商品id
    productId   string 

    // 商品名称
    productName   string 

    // cp订单号
    orderNo   string 

    // 回调地址
    callbackUrl   string 

    // 订单无关的其他参数,如埋点统计的utdid, mac地址等
    extra   string 

    // 订单类型，1为连续包月类型,2为取消连续包月
    orderType   int64 

    // 连续包月原始订单
    originalOrderNo   string 

}

func NewYoukuOttPayOrderDeleteorderRequest() *YoukuOttPayOrderDeleteorderRequest{
    return &YoukuOttPayOrderDeleteorderRequest{
        Params: model.NewParams(),
    }
}

func (r YoukuOttPayOrderDeleteorderRequest) GetApiMethodName() string {
    return "youku.ott.pay.order.deleteorder"
}

func (r YoukuOttPayOrderDeleteorderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YoukuOttPayOrderDeleteorderRequest) SetBuyer(buyer string) error {
    r.buyer = buyer
    r.Set("buyer", buyer)
    return nil
}

func (r YoukuOttPayOrderDeleteorderRequest) GetBuyer() string {
    return r.buyer
}

func (r *YoukuOttPayOrderDeleteorderRequest) SetProductId(productId string) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r YoukuOttPayOrderDeleteorderRequest) GetProductId() string {
    return r.productId
}

func (r *YoukuOttPayOrderDeleteorderRequest) SetProductName(productName string) error {
    r.productName = productName
    r.Set("product_name", productName)
    return nil
}

func (r YoukuOttPayOrderDeleteorderRequest) GetProductName() string {
    return r.productName
}

func (r *YoukuOttPayOrderDeleteorderRequest) SetOrderNo(orderNo string) error {
    r.orderNo = orderNo
    r.Set("order_no", orderNo)
    return nil
}

func (r YoukuOttPayOrderDeleteorderRequest) GetOrderNo() string {
    return r.orderNo
}

func (r *YoukuOttPayOrderDeleteorderRequest) SetCallbackUrl(callbackUrl string) error {
    r.callbackUrl = callbackUrl
    r.Set("callback_url", callbackUrl)
    return nil
}

func (r YoukuOttPayOrderDeleteorderRequest) GetCallbackUrl() string {
    return r.callbackUrl
}

func (r *YoukuOttPayOrderDeleteorderRequest) SetExtra(extra string) error {
    r.extra = extra
    r.Set("extra", extra)
    return nil
}

func (r YoukuOttPayOrderDeleteorderRequest) GetExtra() string {
    return r.extra
}

func (r *YoukuOttPayOrderDeleteorderRequest) SetOrderType(orderType int64) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

func (r YoukuOttPayOrderDeleteorderRequest) GetOrderType() int64 {
    return r.orderType
}

func (r *YoukuOttPayOrderDeleteorderRequest) SetOriginalOrderNo(originalOrderNo string) error {
    r.originalOrderNo = originalOrderNo
    r.Set("original_order_no", originalOrderNo)
    return nil
}

func (r YoukuOttPayOrderDeleteorderRequest) GetOriginalOrderNo() string {
    return r.originalOrderNo
}

