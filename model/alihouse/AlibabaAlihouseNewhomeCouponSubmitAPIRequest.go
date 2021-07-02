package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeCouponSubmitAPIRequest 提交专车优惠券活动 API请求
// alibaba.alihouse.newhome.coupon.submit
//
// 提交专车优惠券活动
type AlibabaAlihouseNewhomeCouponSubmitAPIRequest struct {
	model.Params
	// 打车券活动对象
	_couponDto *MarketingCouponDto
}

// NewAlibabaAlihouseNewhomeCouponSubmitRequest 初始化AlibabaAlihouseNewhomeCouponSubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeCouponSubmitRequest() *AlibabaAlihouseNewhomeCouponSubmitAPIRequest {
	return &AlibabaAlihouseNewhomeCouponSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeCouponSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.coupon.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeCouponSubmitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCouponDto is CouponDto Setter
// 打车券活动对象
func (r *AlibabaAlihouseNewhomeCouponSubmitAPIRequest) SetCouponDto(_couponDto *MarketingCouponDto) error {
	r._couponDto = _couponDto
	r.Set("coupon_dto", _couponDto)
	return nil
}

// GetCouponDto CouponDto Getter
func (r AlibabaAlihouseNewhomeCouponSubmitAPIRequest) GetCouponDto() *MarketingCouponDto {
	return r._couponDto
}
