package icbuseller

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSellerVendorTradePurchaseAPIRequest 查看购买人的订单记录以及授权时间 API请求
// alibaba.seller.vendor.trade.purchase
//
// 查看购买人的订单记录以及授权时间
type AlibabaSellerVendorTradePurchaseAPIRequest struct {
	model.Params
	// 买家登录账号
	_buyerLoginId string
	// 服务code
	_serviceCode string
}

// NewAlibabaSellerVendorTradePurchaseRequest 初始化AlibabaSellerVendorTradePurchaseAPIRequest对象
func NewAlibabaSellerVendorTradePurchaseRequest() *AlibabaSellerVendorTradePurchaseAPIRequest {
	return &AlibabaSellerVendorTradePurchaseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSellerVendorTradePurchaseAPIRequest) GetApiMethodName() string {
	return "alibaba.seller.vendor.trade.purchase"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSellerVendorTradePurchaseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSellerVendorTradePurchaseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBuyerLoginId is BuyerLoginId Setter
// 买家登录账号
func (r *AlibabaSellerVendorTradePurchaseAPIRequest) SetBuyerLoginId(_buyerLoginId string) error {
	r._buyerLoginId = _buyerLoginId
	r.Set("buyer_login_id", _buyerLoginId)
	return nil
}

// GetBuyerLoginId BuyerLoginId Getter
func (r AlibabaSellerVendorTradePurchaseAPIRequest) GetBuyerLoginId() string {
	return r._buyerLoginId
}

// SetServiceCode is ServiceCode Setter
// 服务code
func (r *AlibabaSellerVendorTradePurchaseAPIRequest) SetServiceCode(_serviceCode string) error {
	r._serviceCode = _serviceCode
	r.Set("service_code", _serviceCode)
	return nil
}

// GetServiceCode ServiceCode Getter
func (r AlibabaSellerVendorTradePurchaseAPIRequest) GetServiceCode() string {
	return r._serviceCode
}
