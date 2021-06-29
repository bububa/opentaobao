package ottpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退订应用中心支付订单 API请求
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

// 初始化YoukuOttPayOrderDeleteorderRequest对象
func NewYoukuOttPayOrderDeleteorderRequest() *YoukuOttPayOrderDeleteorderRequest{
    return &YoukuOttPayOrderDeleteorderRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuOttPayOrderDeleteorderRequest) GetApiMethodName() string {
    return "youku.ott.pay.order.deleteorder"
}

// IRequest interface 方法, 获取API参数
func (r YoukuOttPayOrderDeleteorderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Buyer Setter
// 下单账号， cp账号
func (r *YoukuOttPayOrderDeleteorderRequest) SetBuyer(buyer string) error {
    r.buyer = buyer
    r.Set("buyer", buyer)
    return nil
}

// Buyer Getter
func (r YoukuOttPayOrderDeleteorderRequest) GetBuyer() string {
    return r.buyer
}
// ProductId Setter
// 商品id
func (r *YoukuOttPayOrderDeleteorderRequest) SetProductId(productId string) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r YoukuOttPayOrderDeleteorderRequest) GetProductId() string {
    return r.productId
}
// ProductName Setter
// 商品名称
func (r *YoukuOttPayOrderDeleteorderRequest) SetProductName(productName string) error {
    r.productName = productName
    r.Set("product_name", productName)
    return nil
}

// ProductName Getter
func (r YoukuOttPayOrderDeleteorderRequest) GetProductName() string {
    return r.productName
}
// OrderNo Setter
// cp订单号
func (r *YoukuOttPayOrderDeleteorderRequest) SetOrderNo(orderNo string) error {
    r.orderNo = orderNo
    r.Set("order_no", orderNo)
    return nil
}

// OrderNo Getter
func (r YoukuOttPayOrderDeleteorderRequest) GetOrderNo() string {
    return r.orderNo
}
// CallbackUrl Setter
// 回调地址
func (r *YoukuOttPayOrderDeleteorderRequest) SetCallbackUrl(callbackUrl string) error {
    r.callbackUrl = callbackUrl
    r.Set("callback_url", callbackUrl)
    return nil
}

// CallbackUrl Getter
func (r YoukuOttPayOrderDeleteorderRequest) GetCallbackUrl() string {
    return r.callbackUrl
}
// Extra Setter
// 订单无关的其他参数,如埋点统计的utdid, mac地址等
func (r *YoukuOttPayOrderDeleteorderRequest) SetExtra(extra string) error {
    r.extra = extra
    r.Set("extra", extra)
    return nil
}

// Extra Getter
func (r YoukuOttPayOrderDeleteorderRequest) GetExtra() string {
    return r.extra
}
// OrderType Setter
// 订单类型，1为连续包月类型,2为取消连续包月
func (r *YoukuOttPayOrderDeleteorderRequest) SetOrderType(orderType int64) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

// OrderType Getter
func (r YoukuOttPayOrderDeleteorderRequest) GetOrderType() int64 {
    return r.orderType
}
// OriginalOrderNo Setter
// 连续包月原始订单
func (r *YoukuOttPayOrderDeleteorderRequest) SetOriginalOrderNo(originalOrderNo string) error {
    r.originalOrderNo = originalOrderNo
    r.Set("original_order_no", originalOrderNo)
    return nil
}

// OriginalOrderNo Getter
func (r YoukuOttPayOrderDeleteorderRequest) GetOriginalOrderNo() string {
    return r.originalOrderNo
}
