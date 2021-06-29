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
    _buyer   string
    // 商品id
    _productId   string
    // 商品名称
    _productName   string
    // cp订单号
    _orderNo   string
    // 回调地址
    _callbackUrl   string
    // 订单无关的其他参数,如埋点统计的utdid, mac地址等
    _extra   string
    // 订单类型，1为连续包月类型,2为取消连续包月
    _orderType   int64
    // 连续包月原始订单
    _originalOrderNo   string
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
func (r *YoukuOttPayOrderDeleteorderRequest) SetBuyer(_buyer string) error {
    r._buyer = _buyer
    r.Set("buyer", _buyer)
    return nil
}

// Buyer Getter
func (r YoukuOttPayOrderDeleteorderRequest) GetBuyer() string {
    return r._buyer
}
// ProductId Setter
// 商品id
func (r *YoukuOttPayOrderDeleteorderRequest) SetProductId(_productId string) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r YoukuOttPayOrderDeleteorderRequest) GetProductId() string {
    return r._productId
}
// ProductName Setter
// 商品名称
func (r *YoukuOttPayOrderDeleteorderRequest) SetProductName(_productName string) error {
    r._productName = _productName
    r.Set("product_name", _productName)
    return nil
}

// ProductName Getter
func (r YoukuOttPayOrderDeleteorderRequest) GetProductName() string {
    return r._productName
}
// OrderNo Setter
// cp订单号
func (r *YoukuOttPayOrderDeleteorderRequest) SetOrderNo(_orderNo string) error {
    r._orderNo = _orderNo
    r.Set("order_no", _orderNo)
    return nil
}

// OrderNo Getter
func (r YoukuOttPayOrderDeleteorderRequest) GetOrderNo() string {
    return r._orderNo
}
// CallbackUrl Setter
// 回调地址
func (r *YoukuOttPayOrderDeleteorderRequest) SetCallbackUrl(_callbackUrl string) error {
    r._callbackUrl = _callbackUrl
    r.Set("callback_url", _callbackUrl)
    return nil
}

// CallbackUrl Getter
func (r YoukuOttPayOrderDeleteorderRequest) GetCallbackUrl() string {
    return r._callbackUrl
}
// Extra Setter
// 订单无关的其他参数,如埋点统计的utdid, mac地址等
func (r *YoukuOttPayOrderDeleteorderRequest) SetExtra(_extra string) error {
    r._extra = _extra
    r.Set("extra", _extra)
    return nil
}

// Extra Getter
func (r YoukuOttPayOrderDeleteorderRequest) GetExtra() string {
    return r._extra
}
// OrderType Setter
// 订单类型，1为连续包月类型,2为取消连续包月
func (r *YoukuOttPayOrderDeleteorderRequest) SetOrderType(_orderType int64) error {
    r._orderType = _orderType
    r.Set("order_type", _orderType)
    return nil
}

// OrderType Getter
func (r YoukuOttPayOrderDeleteorderRequest) GetOrderType() int64 {
    return r._orderType
}
// OriginalOrderNo Setter
// 连续包月原始订单
func (r *YoukuOttPayOrderDeleteorderRequest) SetOriginalOrderNo(_originalOrderNo string) error {
    r._originalOrderNo = _originalOrderNo
    r.Set("original_order_no", _originalOrderNo)
    return nil
}

// OriginalOrderNo Getter
func (r YoukuOttPayOrderDeleteorderRequest) GetOriginalOrderNo() string {
    return r._originalOrderNo
}
