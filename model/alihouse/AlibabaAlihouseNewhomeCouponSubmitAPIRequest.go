package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomecouponsubmitAPIRequest 提交专车优惠券活动 API请求
// alibaba.alihouse.newhome.coupon.submit
//
// 提交专车优惠券活动
type AlibabaalihousenewhomecouponsubmitAPIRequest struct {
	model.Params
	// 打车券活动对象
	_couponDto *MarketingCouponDto
}

// NewAlibabaalihousenewhomecouponsubmitRequest 初始化AlibabaalihousenewhomecouponsubmitAPIRequest对象
func NewAlibabaalihousenewhomecouponsubmitRequest() *AlibabaalihousenewhomecouponsubmitAPIRequest {
	return &AlibabaalihousenewhomecouponsubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomecouponsubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.coupon.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomecouponsubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomecouponsubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCouponDto is CouponDto Setter
// 打车券活动对象
func (r *AlibabaalihousenewhomecouponsubmitAPIRequest) SetCouponDto(_couponDto *MarketingCouponDto) error {
	r._couponDto = _couponDto
	r.Set("coupon_dto", _couponDto)
	return nil
}

// GetCouponDto CouponDto Getter
func (r AlibabaalihousenewhomecouponsubmitAPIRequest) GetCouponDto() *MarketingCouponDto {
	return r._couponDto
}
