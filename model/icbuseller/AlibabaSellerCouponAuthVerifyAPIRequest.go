package icbuseller

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSellerCouponAuthVerifyAPIRequest 优惠券校验 API请求
// alibaba.seller.coupon.auth.verify
//
// 优惠券校验
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSellerCouponAuthVerifyAPIRequest) Reset() {
	r._serviceCode = ""
	r._couponSeqNumber = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSellerCouponAuthVerifyAPIRequest) GetApiMethodName() string {
	return "alibaba.seller.coupon.auth.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSellerCouponAuthVerifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSellerCouponAuthVerifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceCode is ServiceCode Setter
// 服务代码
func (r *AlibabaSellerCouponAuthVerifyAPIRequest) SetServiceCode(_serviceCode string) error {
	r._serviceCode = _serviceCode
	r.Set("service_code", _serviceCode)
	return nil
}

// GetServiceCode ServiceCode Getter
func (r AlibabaSellerCouponAuthVerifyAPIRequest) GetServiceCode() string {
	return r._serviceCode
}

// SetCouponSeqNumber is CouponSeqNumber Setter
// 卡券验证码
func (r *AlibabaSellerCouponAuthVerifyAPIRequest) SetCouponSeqNumber(_couponSeqNumber string) error {
	r._couponSeqNumber = _couponSeqNumber
	r.Set("coupon_seq_number", _couponSeqNumber)
	return nil
}

// GetCouponSeqNumber CouponSeqNumber Getter
func (r AlibabaSellerCouponAuthVerifyAPIRequest) GetCouponSeqNumber() string {
	return r._couponSeqNumber
}

var poolAlibabaSellerCouponAuthVerifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSellerCouponAuthVerifyRequest()
	},
}

// GetAlibabaSellerCouponAuthVerifyRequest 从 sync.Pool 获取 AlibabaSellerCouponAuthVerifyAPIRequest
func GetAlibabaSellerCouponAuthVerifyAPIRequest() *AlibabaSellerCouponAuthVerifyAPIRequest {
	return poolAlibabaSellerCouponAuthVerifyAPIRequest.Get().(*AlibabaSellerCouponAuthVerifyAPIRequest)
}

// ReleaseAlibabaSellerCouponAuthVerifyAPIRequest 将 AlibabaSellerCouponAuthVerifyAPIRequest 放入 sync.Pool
func ReleaseAlibabaSellerCouponAuthVerifyAPIRequest(v *AlibabaSellerCouponAuthVerifyAPIRequest) {
	v.Reset()
	poolAlibabaSellerCouponAuthVerifyAPIRequest.Put(v)
}
