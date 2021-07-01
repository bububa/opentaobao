package icbuseller

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSellerCouponAuthVerifyAPIRequest
优惠券校验 API请求
alibaba.seller.coupon.auth.verify

优惠券校验 */
type AlibabaSellerCouponAuthVerifyAPIRequest struct {
	model.Params
	// 服务代码
	_serviceCode string
	// 卡券验证码
	_couponSeqNumber string
}

// NewAlibabaSellerCouponAuthVerifyRequest 初始化AlibabaSellerCouponAuthVerifyAPIRequest对象
func NewAlibabaSellerCouponAuthVerifyRequest() *AlibabaSellerCouponAuthVerifyAPIRequest {
	return &AlibabaSellerCouponAuthVerifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSellerCouponAuthVerifyAPIRequest) GetApiMethodName() string {
	return "alibaba.seller.coupon.auth.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSellerCouponAuthVerifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ServiceCode Setter
// 服务代码
func (r *AlibabaSellerCouponAuthVerifyAPIRequest) SetServiceCode(_serviceCode string) error {
	r._serviceCode = _serviceCode
	r.Set("service_code", _serviceCode)
	return nil
}

// Get ServiceCode Getter
func (r AlibabaSellerCouponAuthVerifyAPIRequest) GetServiceCode() string {
	return r._serviceCode
}

// Set is CouponSeqNumber Setter
// 卡券验证码
func (r *AlibabaSellerCouponAuthVerifyAPIRequest) SetCouponSeqNumber(_couponSeqNumber string) error {
	r._couponSeqNumber = _couponSeqNumber
	r.Set("coupon_seq_number", _couponSeqNumber)
	return nil
}

// Get CouponSeqNumber Getter
func (r AlibabaSellerCouponAuthVerifyAPIRequest) GetCouponSeqNumber() string {
	return r._couponSeqNumber
}
