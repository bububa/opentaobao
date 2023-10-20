package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuottpayorderauthpayAPIRequest 委托代扣服务 API请求
// youku.ott.pay.order.authpay
//
// 应用中心sdk连续包月委托代扣服务
type YoukuottpayorderauthpayAPIRequest struct {
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

// NewYoukuottpayorderauthpayRequest 初始化YoukuottpayorderauthpayAPIRequest对象
func NewYoukuottpayorderauthpayRequest() *YoukuottpayorderauthpayAPIRequest {
	return &YoukuottpayorderauthpayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuottpayorderauthpayAPIRequest) GetApiMethodName() string {
	return "youku.ott.pay.order.authpay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuottpayorderauthpayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuottpayorderauthpayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBuyer is Buyer Setter
// cp用户名
func (r *YoukuottpayorderauthpayAPIRequest) SetBuyer(_buyer string) error {
	r._buyer = _buyer
	r.Set("buyer", _buyer)
	return nil
}

// GetBuyer Buyer Getter
func (r YoukuottpayorderauthpayAPIRequest) GetBuyer() string {
	return r._buyer
}

// SetOriginalOrderNo is OriginalOrderNo Setter
// 连续包月原始cp订单号
func (r *YoukuottpayorderauthpayAPIRequest) SetOriginalOrderNo(_originalOrderNo string) error {
	r._originalOrderNo = _originalOrderNo
	r.Set("original_order_no", _originalOrderNo)
	return nil
}

// GetOriginalOrderNo OriginalOrderNo Getter
func (r YoukuottpayorderauthpayAPIRequest) GetOriginalOrderNo() string {
	return r._originalOrderNo
}

// SetOrderNo is OrderNo Setter
// 委托扣款cp订单号
func (r *YoukuottpayorderauthpayAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r YoukuottpayorderauthpayAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// SetProductId is ProductId Setter
// 已配置开通连续包月的产品id
func (r *YoukuottpayorderauthpayAPIRequest) SetProductId(_productId string) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r YoukuottpayorderauthpayAPIRequest) GetProductId() string {
	return r._productId
}

// SetCallbackUrl is CallbackUrl Setter
// 回调
func (r *YoukuottpayorderauthpayAPIRequest) SetCallbackUrl(_callbackUrl string) error {
	r._callbackUrl = _callbackUrl
	r.Set("callback_url", _callbackUrl)
	return nil
}

// GetCallbackUrl CallbackUrl Getter
func (r YoukuottpayorderauthpayAPIRequest) GetCallbackUrl() string {
	return r._callbackUrl
}
