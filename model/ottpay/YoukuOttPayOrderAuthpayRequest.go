package ottpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
委托代扣服务 APIRequest
youku.ott.pay.order.authpay

应用中心sdk连续包月委托代扣服务
*/
type YoukuOttPayOrderAuthpayRequest struct {
    model.Params

    // cp用户名
    buyer   string 

    // 连续包月原始cp订单号
    originalOrderNo   string 

    // 委托扣款cp订单号
    orderNo   string 

    // 已配置开通连续包月的产品id
    productId   string 

    // 回调
    callbackUrl   string 

}

func NewYoukuOttPayOrderAuthpayRequest() *YoukuOttPayOrderAuthpayRequest{
    return &YoukuOttPayOrderAuthpayRequest{
        Params: model.NewParams(),
    }
}

func (r YoukuOttPayOrderAuthpayRequest) GetApiMethodName() string {
    return "youku.ott.pay.order.authpay"
}

func (r YoukuOttPayOrderAuthpayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YoukuOttPayOrderAuthpayRequest) SetBuyer(buyer string) error {
    r.buyer = buyer
    r.Set("buyer", buyer)
    return nil
}

func (r YoukuOttPayOrderAuthpayRequest) GetBuyer() string {
    return r.buyer
}

func (r *YoukuOttPayOrderAuthpayRequest) SetOriginalOrderNo(originalOrderNo string) error {
    r.originalOrderNo = originalOrderNo
    r.Set("original_order_no", originalOrderNo)
    return nil
}

func (r YoukuOttPayOrderAuthpayRequest) GetOriginalOrderNo() string {
    return r.originalOrderNo
}

func (r *YoukuOttPayOrderAuthpayRequest) SetOrderNo(orderNo string) error {
    r.orderNo = orderNo
    r.Set("order_no", orderNo)
    return nil
}

func (r YoukuOttPayOrderAuthpayRequest) GetOrderNo() string {
    return r.orderNo
}

func (r *YoukuOttPayOrderAuthpayRequest) SetProductId(productId string) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r YoukuOttPayOrderAuthpayRequest) GetProductId() string {
    return r.productId
}

func (r *YoukuOttPayOrderAuthpayRequest) SetCallbackUrl(callbackUrl string) error {
    r.callbackUrl = callbackUrl
    r.Set("callback_url", callbackUrl)
    return nil
}

func (r YoukuOttPayOrderAuthpayRequest) GetCallbackUrl() string {
    return r.callbackUrl
}

