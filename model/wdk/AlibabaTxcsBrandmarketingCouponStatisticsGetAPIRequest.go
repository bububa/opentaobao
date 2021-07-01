package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌营销导购员券推广统计数据回流 API请求
alibaba.txcs.brandmarketing.coupon.statistics.get

请求券统计数据回流
*/
type AlibabaTxcsBrandmarketingCouponStatisticsGetAPIRequest struct {
    model.Params
    // 请求信息
    _couponStatisticsParamDo   *CouponStatisticsParamDo
}

// 初始化AlibabaTxcsBrandmarketingCouponStatisticsGetAPIRequest对象
func NewAlibabaTxcsBrandmarketingCouponStatisticsGetRequest() *AlibabaTxcsBrandmarketingCouponStatisticsGetAPIRequest{
    return &AlibabaTxcsBrandmarketingCouponStatisticsGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTxcsBrandmarketingCouponStatisticsGetAPIRequest) GetApiMethodName() string {
    return "alibaba.txcs.brandmarketing.coupon.statistics.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTxcsBrandmarketingCouponStatisticsGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CouponStatisticsParamDo Setter
// 请求信息
func (r *AlibabaTxcsBrandmarketingCouponStatisticsGetAPIRequest) SetCouponStatisticsParamDo(_couponStatisticsParamDo *CouponStatisticsParamDo) error {
    r._couponStatisticsParamDo = _couponStatisticsParamDo
    r.Set("coupon_statistics_param_do", _couponStatisticsParamDo)
    return nil
}

// CouponStatisticsParamDo Getter
func (r AlibabaTxcsBrandmarketingCouponStatisticsGetAPIRequest) GetCouponStatisticsParamDo() *CouponStatisticsParamDo {
    return r._couponStatisticsParamDo
}
