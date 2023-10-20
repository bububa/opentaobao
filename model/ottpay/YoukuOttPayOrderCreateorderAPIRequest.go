package ottpay

import (
	"net/url"
	"sync"

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
	// 连续包月实际参数
	_realPrice string
	// 订单类型，1为连续包月类型
	_orderType int64
}

// NewYoukuOttPayOrderCreateorderRequest 初始化YoukuOttPayOrderCreateorderAPIRequest对象
func NewYoukuOttPayOrderCreateorderRequest() *YoukuOttPayOrderCreateorderAPIRequest {
	return &YoukuOttPayOrderCreateorderAPIRequest{
		Params: model.NewParams(9),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YoukuOttPayOrderCreateorderAPIRequest) Reset() {
	r._buyer = ""
	r._productId = ""
	r._productName = ""
	r._orderNo = ""
	r._price = ""
	r._callbackUrl = ""
	r._extra = ""
	r._realPrice = ""
	r._orderType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuOttPayOrderCreateorderAPIRequest) GetApiMethodName() string {
	return "youku.ott.pay.order.createorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuOttPayOrderCreateorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuOttPayOrderCreateorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBuyer is Buyer Setter
// 下单账号， cp账号
func (r *YoukuOttPayOrderCreateorderAPIRequest) SetBuyer(_buyer string) error {
	r._buyer = _buyer
	r.Set("buyer", _buyer)
	return nil
}

// GetBuyer Buyer Getter
func (r YoukuOttPayOrderCreateorderAPIRequest) GetBuyer() string {
	return r._buyer
}

// SetProductId is ProductId Setter
// 商品id
func (r *YoukuOttPayOrderCreateorderAPIRequest) SetProductId(_productId string) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r YoukuOttPayOrderCreateorderAPIRequest) GetProductId() string {
	return r._productId
}

// SetProductName is ProductName Setter
// 商品名称
func (r *YoukuOttPayOrderCreateorderAPIRequest) SetProductName(_productName string) error {
	r._productName = _productName
	r.Set("product_name", _productName)
	return nil
}

// GetProductName ProductName Getter
func (r YoukuOttPayOrderCreateorderAPIRequest) GetProductName() string {
	return r._productName
}

// SetOrderNo is OrderNo Setter
// cp订单号
func (r *YoukuOttPayOrderCreateorderAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r YoukuOttPayOrderCreateorderAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// SetPrice is Price Setter
// 价格， 单位：分
func (r *YoukuOttPayOrderCreateorderAPIRequest) SetPrice(_price string) error {
	r._price = _price
	r.Set("price", _price)
	return nil
}

// GetPrice Price Getter
func (r YoukuOttPayOrderCreateorderAPIRequest) GetPrice() string {
	return r._price
}

// SetCallbackUrl is CallbackUrl Setter
// 回调接口
func (r *YoukuOttPayOrderCreateorderAPIRequest) SetCallbackUrl(_callbackUrl string) error {
	r._callbackUrl = _callbackUrl
	r.Set("callback_url", _callbackUrl)
	return nil
}

// GetCallbackUrl CallbackUrl Getter
func (r YoukuOttPayOrderCreateorderAPIRequest) GetCallbackUrl() string {
	return r._callbackUrl
}

// SetExtra is Extra Setter
// 订单无关的其他参数,如埋点统计的utdid, mac地址等
func (r *YoukuOttPayOrderCreateorderAPIRequest) SetExtra(_extra string) error {
	r._extra = _extra
	r.Set("extra", _extra)
	return nil
}

// GetExtra Extra Getter
func (r YoukuOttPayOrderCreateorderAPIRequest) GetExtra() string {
	return r._extra
}

// SetRealPrice is RealPrice Setter
// 连续包月实际参数
func (r *YoukuOttPayOrderCreateorderAPIRequest) SetRealPrice(_realPrice string) error {
	r._realPrice = _realPrice
	r.Set("real_price", _realPrice)
	return nil
}

// GetRealPrice RealPrice Getter
func (r YoukuOttPayOrderCreateorderAPIRequest) GetRealPrice() string {
	return r._realPrice
}

// SetOrderType is OrderType Setter
// 订单类型，1为连续包月类型
func (r *YoukuOttPayOrderCreateorderAPIRequest) SetOrderType(_orderType int64) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// GetOrderType OrderType Getter
func (r YoukuOttPayOrderCreateorderAPIRequest) GetOrderType() int64 {
	return r._orderType
}

var poolYoukuOttPayOrderCreateorderAPIRequest = sync.Pool{
	New: func() any {
		return NewYoukuOttPayOrderCreateorderRequest()
	},
}

// GetYoukuOttPayOrderCreateorderRequest 从 sync.Pool 获取 YoukuOttPayOrderCreateorderAPIRequest
func GetYoukuOttPayOrderCreateorderAPIRequest() *YoukuOttPayOrderCreateorderAPIRequest {
	return poolYoukuOttPayOrderCreateorderAPIRequest.Get().(*YoukuOttPayOrderCreateorderAPIRequest)
}

// ReleaseYoukuOttPayOrderCreateorderAPIRequest 将 YoukuOttPayOrderCreateorderAPIRequest 放入 sync.Pool
func ReleaseYoukuOttPayOrderCreateorderAPIRequest(v *YoukuOttPayOrderCreateorderAPIRequest) {
	v.Reset()
	poolYoukuOttPayOrderCreateorderAPIRequest.Put(v)
}
