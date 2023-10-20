package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuottpayordercreateorderAPIRequest 创建订单 API请求
// youku.ott.pay.order.createorder
//
// ottpay创建订单
type YoukuottpayordercreateorderAPIRequest struct {
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
	// 连续包月实际参数
	_realPrice string
	// 订单类型，1为连续包月类型
	_orderType int64
}

// NewYoukuottpayordercreateorderRequest 初始化YoukuottpayordercreateorderAPIRequest对象
func NewYoukuottpayordercreateorderRequest() *YoukuottpayordercreateorderAPIRequest {
	return &YoukuottpayordercreateorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuottpayordercreateorderAPIRequest) GetApiMethodName() string {
	return "youku.ott.pay.order.createorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuottpayordercreateorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuottpayordercreateorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBuyer is Buyer Setter
// 下单账号， cp账号
func (r *YoukuottpayordercreateorderAPIRequest) SetBuyer(_buyer string) error {
	r._buyer = _buyer
	r.Set("buyer", _buyer)
	return nil
}

// GetBuyer Buyer Getter
func (r YoukuottpayordercreateorderAPIRequest) GetBuyer() string {
	return r._buyer
}

// SetProductId is ProductId Setter
// 商品id
func (r *YoukuottpayordercreateorderAPIRequest) SetProductId(_productId string) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r YoukuottpayordercreateorderAPIRequest) GetProductId() string {
	return r._productId
}

// SetProductName is ProductName Setter
// 商品名称
func (r *YoukuottpayordercreateorderAPIRequest) SetProductName(_productName string) error {
	r._productName = _productName
	r.Set("product_name", _productName)
	return nil
}

// GetProductName ProductName Getter
func (r YoukuottpayordercreateorderAPIRequest) GetProductName() string {
	return r._productName
}

// SetOrderNo is OrderNo Setter
// cp订单号
func (r *YoukuottpayordercreateorderAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r YoukuottpayordercreateorderAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// SetPrice is Price Setter
// 价格， 单位：分
func (r *YoukuottpayordercreateorderAPIRequest) SetPrice(_price string) error {
	r._price = _price
	r.Set("price", _price)
	return nil
}

// GetPrice Price Getter
func (r YoukuottpayordercreateorderAPIRequest) GetPrice() string {
	return r._price
}

// SetCallbackUrl is CallbackUrl Setter
// 回调接口
func (r *YoukuottpayordercreateorderAPIRequest) SetCallbackUrl(_callbackUrl string) error {
	r._callbackUrl = _callbackUrl
	r.Set("callback_url", _callbackUrl)
	return nil
}

// GetCallbackUrl CallbackUrl Getter
func (r YoukuottpayordercreateorderAPIRequest) GetCallbackUrl() string {
	return r._callbackUrl
}

// SetExtra is Extra Setter
// 订单无关的其他参数,如埋点统计的utdid, mac地址等
func (r *YoukuottpayordercreateorderAPIRequest) SetExtra(_extra string) error {
	r._extra = _extra
	r.Set("extra", _extra)
	return nil
}

// GetExtra Extra Getter
func (r YoukuottpayordercreateorderAPIRequest) GetExtra() string {
	return r._extra
}

// SetRealPrice is RealPrice Setter
// 连续包月实际参数
func (r *YoukuottpayordercreateorderAPIRequest) SetRealPrice(_realPrice string) error {
	r._realPrice = _realPrice
	r.Set("real_price", _realPrice)
	return nil
}

// GetRealPrice RealPrice Getter
func (r YoukuottpayordercreateorderAPIRequest) GetRealPrice() string {
	return r._realPrice
}

// SetOrderType is OrderType Setter
// 订单类型，1为连续包月类型
func (r *YoukuottpayordercreateorderAPIRequest) SetOrderType(_orderType int64) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// GetOrderType OrderType Getter
func (r YoukuottpayordercreateorderAPIRequest) GetOrderType() int64 {
	return r._orderType
}
