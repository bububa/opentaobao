package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuOttPayOrderAuthpayAPIRequest
委托代扣服务 API请求
youku.ott.pay.order.authpay

应用中心sdk连续包月委托代扣服务 */
type YoukuOttPayOrderAuthpayAPIRequest struct {
	model.Params
	// cp用户名
	_buyer string
	// 连续包月原始cp订单号
	_originalOrderNo string
	// 委托扣款cp订单号
	_orderNo string
	// 已配置开通连续包月的产品id
	_productId string
	// 回调
	_callbackUrl string
}

// NewYoukuOttPayOrderAuthpayRequest 初始化YoukuOttPayOrderAuthpayAPIRequest对象
func NewYoukuOttPayOrderAuthpayRequest() *YoukuOttPayOrderAuthpayAPIRequest {
	return &YoukuOttPayOrderAuthpayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuOttPayOrderAuthpayAPIRequest) GetApiMethodName() string {
	return "youku.ott.pay.order.authpay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuOttPayOrderAuthpayAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Buyer Setter
// cp用户名
func (r *YoukuOttPayOrderAuthpayAPIRequest) SetBuyer(_buyer string) error {
	r._buyer = _buyer
	r.Set("buyer", _buyer)
	return nil
}

// Get Buyer Getter
func (r YoukuOttPayOrderAuthpayAPIRequest) GetBuyer() string {
	return r._buyer
}

// Set is OriginalOrderNo Setter
// 连续包月原始cp订单号
func (r *YoukuOttPayOrderAuthpayAPIRequest) SetOriginalOrderNo(_originalOrderNo string) error {
	r._originalOrderNo = _originalOrderNo
	r.Set("original_order_no", _originalOrderNo)
	return nil
}

// Get OriginalOrderNo Getter
func (r YoukuOttPayOrderAuthpayAPIRequest) GetOriginalOrderNo() string {
	return r._originalOrderNo
}

// Set is OrderNo Setter
// 委托扣款cp订单号
func (r *YoukuOttPayOrderAuthpayAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// Get OrderNo Getter
func (r YoukuOttPayOrderAuthpayAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// Set is ProductId Setter
// 已配置开通连续包月的产品id
func (r *YoukuOttPayOrderAuthpayAPIRequest) SetProductId(_productId string) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// Get ProductId Getter
func (r YoukuOttPayOrderAuthpayAPIRequest) GetProductId() string {
	return r._productId
}

// Set is CallbackUrl Setter
// 回调
func (r *YoukuOttPayOrderAuthpayAPIRequest) SetCallbackUrl(_callbackUrl string) error {
	r._callbackUrl = _callbackUrl
	r.Set("callback_url", _callbackUrl)
	return nil
}

// Get CallbackUrl Getter
func (r YoukuOttPayOrderAuthpayAPIRequest) GetCallbackUrl() string {
	return r._callbackUrl
}
