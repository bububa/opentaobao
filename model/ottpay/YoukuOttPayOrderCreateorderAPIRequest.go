package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttPayOrderCreateorderAPIRequest 创建订单 API请求
// youku.ott.pay.order.createorder
//
// ottpay创建订单
type YoukuOttPayOrderCreateorderAPIRequest struct {
	model.Params
	// 下单账号， cp账号
	_buyer string
	// 商品id
	_productId string
	// 商品名称
	_productName string
	// cp订单号
	_orderNo string
	// 价格， 单位：分
	_price string
	// 回调接口
	_callbackUrl string
	// 订单无关的其他参数,如埋点统计的utdid, mac地址等
	_extra string
	// 订单类型，1为连续包月类型
	_orderType int64
	// 连续包月实际参数
	_realPrice string
}

// NewYoukuOttPayOrderCreateorderRequest 初始化YoukuOttPayOrderCreateorderAPIRequest对象
func NewYoukuOttPayOrderCreateorderRequest() *YoukuOttPayOrderCreateorderAPIRequest {
	return &YoukuOttPayOrderCreateorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuOttPayOrderCreateorderAPIRequest) GetApiMethodName() string {
	return "youku.ott.pay.order.createorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuOttPayOrderCreateorderAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Buyer Setter
// 下单账号， cp账号
func (r *YoukuOttPayOrderCreateorderAPIRequest) SetBuyer(_buyer string) error {
	r._buyer = _buyer
	r.Set("buyer", _buyer)
	return nil
}

// Get Buyer Getter
func (r YoukuOttPayOrderCreateorderAPIRequest) GetBuyer() string {
	return r._buyer
}

// Set is ProductId Setter
// 商品id
func (r *YoukuOttPayOrderCreateorderAPIRequest) SetProductId(_productId string) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// Get ProductId Getter
func (r YoukuOttPayOrderCreateorderAPIRequest) GetProductId() string {
	return r._productId
}

// Set is ProductName Setter
// 商品名称
func (r *YoukuOttPayOrderCreateorderAPIRequest) SetProductName(_productName string) error {
	r._productName = _productName
	r.Set("product_name", _productName)
	return nil
}

// Get ProductName Getter
func (r YoukuOttPayOrderCreateorderAPIRequest) GetProductName() string {
	return r._productName
}

// Set is OrderNo Setter
// cp订单号
func (r *YoukuOttPayOrderCreateorderAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// Get OrderNo Getter
func (r YoukuOttPayOrderCreateorderAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// Set is Price Setter
// 价格， 单位：分
func (r *YoukuOttPayOrderCreateorderAPIRequest) SetPrice(_price string) error {
	r._price = _price
	r.Set("price", _price)
	return nil
}

// Get Price Getter
func (r YoukuOttPayOrderCreateorderAPIRequest) GetPrice() string {
	return r._price
}

// Set is CallbackUrl Setter
// 回调接口
func (r *YoukuOttPayOrderCreateorderAPIRequest) SetCallbackUrl(_callbackUrl string) error {
	r._callbackUrl = _callbackUrl
	r.Set("callback_url", _callbackUrl)
	return nil
}

// Get CallbackUrl Getter
func (r YoukuOttPayOrderCreateorderAPIRequest) GetCallbackUrl() string {
	return r._callbackUrl
}

// Set is Extra Setter
// 订单无关的其他参数,如埋点统计的utdid, mac地址等
func (r *YoukuOttPayOrderCreateorderAPIRequest) SetExtra(_extra string) error {
	r._extra = _extra
	r.Set("extra", _extra)
	return nil
}

// Get Extra Getter
func (r YoukuOttPayOrderCreateorderAPIRequest) GetExtra() string {
	return r._extra
}

// Set is OrderType Setter
// 订单类型，1为连续包月类型
func (r *YoukuOttPayOrderCreateorderAPIRequest) SetOrderType(_orderType int64) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// Get OrderType Getter
func (r YoukuOttPayOrderCreateorderAPIRequest) GetOrderType() int64 {
	return r._orderType
}

// Set is RealPrice Setter
// 连续包月实际参数
func (r *YoukuOttPayOrderCreateorderAPIRequest) SetRealPrice(_realPrice string) error {
	r._realPrice = _realPrice
	r.Set("real_price", _realPrice)
	return nil
}

// Get RealPrice Getter
func (r YoukuOttPayOrderCreateorderAPIRequest) GetRealPrice() string {
	return r._realPrice
}
