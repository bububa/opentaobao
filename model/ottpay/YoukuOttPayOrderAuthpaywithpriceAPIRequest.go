package ottpay

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttPayOrderAuthpaywithpriceAPIRequest 委托代扣可配定价服务 API请求
// youku.ott.pay.order.authpaywithprice
//
// 应用中心sdk连续包月委托代扣服务，次月可配置营销价
type YoukuOttPayOrderAuthpaywithpriceAPIRequest struct {
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
	// 用于次月扣款营销价，需小于签约时设定续费价
	_payPrice string
}

// NewYoukuOttPayOrderAuthpaywithpriceRequest 初始化YoukuOttPayOrderAuthpaywithpriceAPIRequest对象
func NewYoukuOttPayOrderAuthpaywithpriceRequest() *YoukuOttPayOrderAuthpaywithpriceAPIRequest {
	return &YoukuOttPayOrderAuthpaywithpriceAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YoukuOttPayOrderAuthpaywithpriceAPIRequest) Reset() {
	r._buyer = ""
	r._originalOrderNo = ""
	r._orderNo = ""
	r._productId = ""
	r._callbackUrl = ""
	r._payPrice = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuOttPayOrderAuthpaywithpriceAPIRequest) GetApiMethodName() string {
	return "youku.ott.pay.order.authpaywithprice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuOttPayOrderAuthpaywithpriceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuOttPayOrderAuthpaywithpriceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBuyer is Buyer Setter
// cp用户名
func (r *YoukuOttPayOrderAuthpaywithpriceAPIRequest) SetBuyer(_buyer string) error {
	r._buyer = _buyer
	r.Set("buyer", _buyer)
	return nil
}

// GetBuyer Buyer Getter
func (r YoukuOttPayOrderAuthpaywithpriceAPIRequest) GetBuyer() string {
	return r._buyer
}

// SetOriginalOrderNo is OriginalOrderNo Setter
// 连续包月原始cp订单号
func (r *YoukuOttPayOrderAuthpaywithpriceAPIRequest) SetOriginalOrderNo(_originalOrderNo string) error {
	r._originalOrderNo = _originalOrderNo
	r.Set("original_order_no", _originalOrderNo)
	return nil
}

// GetOriginalOrderNo OriginalOrderNo Getter
func (r YoukuOttPayOrderAuthpaywithpriceAPIRequest) GetOriginalOrderNo() string {
	return r._originalOrderNo
}

// SetOrderNo is OrderNo Setter
// 委托扣款cp订单号
func (r *YoukuOttPayOrderAuthpaywithpriceAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r YoukuOttPayOrderAuthpaywithpriceAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// SetProductId is ProductId Setter
// 已配置开通连续包月的产品id
func (r *YoukuOttPayOrderAuthpaywithpriceAPIRequest) SetProductId(_productId string) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r YoukuOttPayOrderAuthpaywithpriceAPIRequest) GetProductId() string {
	return r._productId
}

// SetCallbackUrl is CallbackUrl Setter
// 回调
func (r *YoukuOttPayOrderAuthpaywithpriceAPIRequest) SetCallbackUrl(_callbackUrl string) error {
	r._callbackUrl = _callbackUrl
	r.Set("callback_url", _callbackUrl)
	return nil
}

// GetCallbackUrl CallbackUrl Getter
func (r YoukuOttPayOrderAuthpaywithpriceAPIRequest) GetCallbackUrl() string {
	return r._callbackUrl
}

// SetPayPrice is PayPrice Setter
// 用于次月扣款营销价，需小于签约时设定续费价
func (r *YoukuOttPayOrderAuthpaywithpriceAPIRequest) SetPayPrice(_payPrice string) error {
	r._payPrice = _payPrice
	r.Set("pay_price", _payPrice)
	return nil
}

// GetPayPrice PayPrice Getter
func (r YoukuOttPayOrderAuthpaywithpriceAPIRequest) GetPayPrice() string {
	return r._payPrice
}

var poolYoukuOttPayOrderAuthpaywithpriceAPIRequest = sync.Pool{
	New: func() any {
		return NewYoukuOttPayOrderAuthpaywithpriceRequest()
	},
}

// GetYoukuOttPayOrderAuthpaywithpriceRequest 从 sync.Pool 获取 YoukuOttPayOrderAuthpaywithpriceAPIRequest
func GetYoukuOttPayOrderAuthpaywithpriceAPIRequest() *YoukuOttPayOrderAuthpaywithpriceAPIRequest {
	return poolYoukuOttPayOrderAuthpaywithpriceAPIRequest.Get().(*YoukuOttPayOrderAuthpaywithpriceAPIRequest)
}

// ReleaseYoukuOttPayOrderAuthpaywithpriceAPIRequest 将 YoukuOttPayOrderAuthpaywithpriceAPIRequest 放入 sync.Pool
func ReleaseYoukuOttPayOrderAuthpaywithpriceAPIRequest(v *YoukuOttPayOrderAuthpaywithpriceAPIRequest) {
	v.Reset()
	poolYoukuOttPayOrderAuthpaywithpriceAPIRequest.Put(v)
}
