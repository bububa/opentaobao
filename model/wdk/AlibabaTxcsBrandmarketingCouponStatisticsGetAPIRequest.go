package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTxcsBrandmarketingCouponStatisticsGetAPIRequest 品牌营销导购员券推广统计数据回流 API请求
// alibaba.txcs.brandmarketing.coupon.statistics.get
//
// 请求券统计数据回流
type AlibabaTxcsBrandmarketingCouponStatisticsGetAPIRequest struct {
	model.Params
	// 请求信息
	_couponStatisticsParamDo *CouponStatisticsParamDo
}

// NewAlibabaTxcsBrandmarketingCouponStatisticsGetRequest 初始化AlibabaTxcsBrandmarketingCouponStatisticsGetAPIRequest对象
func NewAlibabaTxcsBrandmarketingCouponStatisticsGetRequest() *AlibabaTxcsBrandmarketingCouponStatisticsGetAPIRequest {
	return &AlibabaTxcsBrandmarketingCouponStatisticsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTxcsBrandmarketingCouponStatisticsGetAPIRequest) GetApiMethodName() string {
	return "alibaba.txcs.brandmarketing.coupon.statistics.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTxcsBrandmarketingCouponStatisticsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTxcsBrandmarketingCouponStatisticsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCouponStatisticsParamDo is CouponStatisticsParamDo Setter
// 请求信息
func (r *AlibabaTxcsBrandmarketingCouponStatisticsGetAPIRequest) SetCouponStatisticsParamDo(_couponStatisticsParamDo *CouponStatisticsParamDo) error {
	r._couponStatisticsParamDo = _couponStatisticsParamDo
	r.Set("coupon_statistics_param_do", _couponStatisticsParamDo)
	return nil
}

// GetCouponStatisticsParamDo CouponStatisticsParamDo Getter
func (r AlibabaTxcsBrandmarketingCouponStatisticsGetAPIRequest) GetCouponStatisticsParamDo() *CouponStatisticsParamDo {
	return r._couponStatisticsParamDo
}
