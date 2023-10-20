package icbuseller

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasellercouponauthverifyAPIRequest 优惠券校验 API请求
// alibaba.seller.coupon.auth.verify
//
// 优惠券校验
type AlibabasellercouponauthverifyAPIRequest struct {
	model.Params
	// 服务代码
	_serviceCode string
	// 卡券验证码
	_couponSeqNumber string
}

// NewAlibabasellercouponauthverifyRequest 初始化AlibabasellercouponauthverifyAPIRequest对象
func NewAlibabasellercouponauthverifyRequest() *AlibabasellercouponauthverifyAPIRequest {
	return &AlibabasellercouponauthverifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasellercouponauthverifyAPIRequest) GetApiMethodName() string {
	return "alibaba.seller.coupon.auth.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasellercouponauthverifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasellercouponauthverifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceCode is ServiceCode Setter
// 服务代码
func (r *AlibabasellercouponauthverifyAPIRequest) SetServiceCode(_serviceCode string) error {
	r._serviceCode = _serviceCode
	r.Set("service_code", _serviceCode)
	return nil
}

// GetServiceCode ServiceCode Getter
func (r AlibabasellercouponauthverifyAPIRequest) GetServiceCode() string {
	return r._serviceCode
}

// SetCouponSeqNumber is CouponSeqNumber Setter
// 卡券验证码
func (r *AlibabasellercouponauthverifyAPIRequest) SetCouponSeqNumber(_couponSeqNumber string) error {
	r._couponSeqNumber = _couponSeqNumber
	r.Set("coupon_seq_number", _couponSeqNumber)
	return nil
}

// GetCouponSeqNumber CouponSeqNumber Getter
func (r AlibabasellercouponauthverifyAPIRequest) GetCouponSeqNumber() string {
	return r._couponSeqNumber
}
