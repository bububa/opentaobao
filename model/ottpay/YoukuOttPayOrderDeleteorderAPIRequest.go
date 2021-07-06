package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttPayOrderDeleteorderAPIRequest 退订应用中心支付订单 API请求
// youku.ott.pay.order.deleteorder
//
// 应用中心sdk连续包月退订接口
type YoukuOttPayOrderDeleteorderAPIRequest struct {
	model.Params
	// 下单账号， cp账号
	_buyer string
	// 商品id
	_productId string
	// 商品名称
	_productName string
	// cp订单号
	_orderNo string
	// 回调地址
	_callbackUrl string
	// 订单无关的其他参数,如埋点统计的utdid, mac地址等
	_extra string
	// 连续包月原始订单
	_originalOrderNo string
	// 订单类型，1为连续包月类型,2为取消连续包月
	_orderType int64
}

// NewYoukuOttPayOrderDeleteorderRequest 初始化YoukuOttPayOrderDeleteorderAPIRequest对象
func NewYoukuOttPayOrderDeleteorderRequest() *YoukuOttPayOrderDeleteorderAPIRequest {
	return &YoukuOttPayOrderDeleteorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuOttPayOrderDeleteorderAPIRequest) GetApiMethodName() string {
	return "youku.ott.pay.order.deleteorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuOttPayOrderDeleteorderAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBuyer is Buyer Setter
// 下单账号， cp账号
func (r *YoukuOttPayOrderDeleteorderAPIRequest) SetBuyer(_buyer string) error {
	r._buyer = _buyer
	r.Set("buyer", _buyer)
	return nil
}

// GetBuyer Buyer Getter
func (r YoukuOttPayOrderDeleteorderAPIRequest) GetBuyer() string {
	return r._buyer
}

// SetProductId is ProductId Setter
// 商品id
func (r *YoukuOttPayOrderDeleteorderAPIRequest) SetProductId(_productId string) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r YoukuOttPayOrderDeleteorderAPIRequest) GetProductId() string {
	return r._productId
}

// SetProductName is ProductName Setter
// 商品名称
func (r *YoukuOttPayOrderDeleteorderAPIRequest) SetProductName(_productName string) error {
	r._productName = _productName
	r.Set("product_name", _productName)
	return nil
}

// GetProductName ProductName Getter
func (r YoukuOttPayOrderDeleteorderAPIRequest) GetProductName() string {
	return r._productName
}

// SetOrderNo is OrderNo Setter
// cp订单号
func (r *YoukuOttPayOrderDeleteorderAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r YoukuOttPayOrderDeleteorderAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// SetCallbackUrl is CallbackUrl Setter
// 回调地址
func (r *YoukuOttPayOrderDeleteorderAPIRequest) SetCallbackUrl(_callbackUrl string) error {
	r._callbackUrl = _callbackUrl
	r.Set("callback_url", _callbackUrl)
	return nil
}

// GetCallbackUrl CallbackUrl Getter
func (r YoukuOttPayOrderDeleteorderAPIRequest) GetCallbackUrl() string {
	return r._callbackUrl
}

// SetExtra is Extra Setter
// 订单无关的其他参数,如埋点统计的utdid, mac地址等
func (r *YoukuOttPayOrderDeleteorderAPIRequest) SetExtra(_extra string) error {
	r._extra = _extra
	r.Set("extra", _extra)
	return nil
}

// GetExtra Extra Getter
func (r YoukuOttPayOrderDeleteorderAPIRequest) GetExtra() string {
	return r._extra
}

// SetOriginalOrderNo is OriginalOrderNo Setter
// 连续包月原始订单
func (r *YoukuOttPayOrderDeleteorderAPIRequest) SetOriginalOrderNo(_originalOrderNo string) error {
	r._originalOrderNo = _originalOrderNo
	r.Set("original_order_no", _originalOrderNo)
	return nil
}

// GetOriginalOrderNo OriginalOrderNo Getter
func (r YoukuOttPayOrderDeleteorderAPIRequest) GetOriginalOrderNo() string {
	return r._originalOrderNo
}

// SetOrderType is OrderType Setter
// 订单类型，1为连续包月类型,2为取消连续包月
func (r *YoukuOttPayOrderDeleteorderAPIRequest) SetOrderType(_orderType int64) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// GetOrderType OrderType Getter
func (r YoukuOttPayOrderDeleteorderAPIRequest) GetOrderType() int64 {
	return r._orderType
}
