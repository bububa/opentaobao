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
    _buyer   string
    // 商品id
    _productId   string
    // 商品名称
    _productName   string
    // cp订单号
    _orderNo   string
    // 价格， 单位：分
    _price   string
    // 回调接口
    _callbackUrl   string
    // 订单无关的其他参数,如埋点统计的utdid, mac地址等
    _extra   string
    // 订单类型，1为连续包月类型
    _orderType   int64
    // 连续包月实际参数
    _realPrice   string
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
func (r *YoukuOttPayOrderCreateorderRequest) SetBuyer(_buyer string) error {
    r._buyer = _buyer
    r.Set("buyer", _buyer)
    return nil
}

// Buyer Getter
func (r YoukuOttPayOrderCreateorderRequest) GetBuyer() string {
    return r._buyer
}
// ProductId Setter
// 商品id
func (r *YoukuOttPayOrderCreateorderRequest) SetProductId(_productId string) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r YoukuOttPayOrderCreateorderRequest) GetProductId() string {
    return r._productId
}
// ProductName Setter
// 商品名称
func (r *YoukuOttPayOrderCreateorderRequest) SetProductName(_productName string) error {
    r._productName = _productName
    r.Set("product_name", _productName)
    return nil
}

// ProductName Getter
func (r YoukuOttPayOrderCreateorderRequest) GetProductName() string {
    return r._productName
}
// OrderNo Setter
// cp订单号
func (r *YoukuOttPayOrderCreateorderRequest) SetOrderNo(_orderNo string) error {
    r._orderNo = _orderNo
    r.Set("order_no", _orderNo)
    return nil
}

// OrderNo Getter
func (r YoukuOttPayOrderCreateorderRequest) GetOrderNo() string {
    return r._orderNo
}
// Price Setter
// 价格， 单位：分
func (r *YoukuOttPayOrderCreateorderRequest) SetPrice(_price string) error {
    r._price = _price
    r.Set("price", _price)
    return nil
}

// Price Getter
func (r YoukuOttPayOrderCreateorderRequest) GetPrice() string {
    return r._price
}
// CallbackUrl Setter
// 回调接口
func (r *YoukuOttPayOrderCreateorderRequest) SetCallbackUrl(_callbackUrl string) error {
    r._callbackUrl = _callbackUrl
    r.Set("callback_url", _callbackUrl)
    return nil
}

// CallbackUrl Getter
func (r YoukuOttPayOrderCreateorderRequest) GetCallbackUrl() string {
    return r._callbackUrl
}
// Extra Setter
// 订单无关的其他参数,如埋点统计的utdid, mac地址等
func (r *YoukuOttPayOrderCreateorderRequest) SetExtra(_extra string) error {
    r._extra = _extra
    r.Set("extra", _extra)
    return nil
}

// Extra Getter
func (r YoukuOttPayOrderCreateorderRequest) GetExtra() string {
    return r._extra
}
// OrderType Setter
// 订单类型，1为连续包月类型
func (r *YoukuOttPayOrderCreateorderRequest) SetOrderType(_orderType int64) error {
    r._orderType = _orderType
    r.Set("order_type", _orderType)
    return nil
}

// OrderType Getter
func (r YoukuOttPayOrderCreateorderRequest) GetOrderType() int64 {
    return r._orderType
}
// RealPrice Setter
// 连续包月实际参数
func (r *YoukuOttPayOrderCreateorderRequest) SetRealPrice(_realPrice string) error {
    r._realPrice = _realPrice
    r.Set("real_price", _realPrice)
    return nil
}

// RealPrice Getter
func (r YoukuOttPayOrderCreateorderRequest) GetRealPrice() string {
    return r._realPrice
}
