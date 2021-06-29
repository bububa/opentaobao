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
type AlibabaTxcsBrandmarketingCouponStatisticsGetRequest struct {
    model.Params
    // 请求信息
    _couponStatisticsParamDo   *CouponStatisticsParamDO
}

// 初始化AlibabaTxcsBrandmarketingCouponStatisticsGetRequest对象
func NewAlibabaTxcsBrandmarketingCouponStatisticsGetRequest() *AlibabaTxcsBrandmarketingCouponStatisticsGetRequest{
    return &AlibabaTxcsBrandmarketingCouponStatisticsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTxcsBrandmarketingCouponStatisticsGetRequest) GetApiMethodName() string {
    return "alibaba.txcs.brandmarketing.coupon.statistics.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTxcsBrandmarketingCouponStatisticsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CouponStatisticsParamDo Setter
// 请求信息
func (r *AlibabaTxcsBrandmarketingCouponStatisticsGetRequest) SetCouponStatisticsParamDo(_couponStatisticsParamDo *CouponStatisticsParamDO) error {
    r._couponStatisticsParamDo = _couponStatisticsParamDo
    r.Set("coupon_statistics_param_do", _couponStatisticsParamDo)
    return nil
}

// CouponStatisticsParamDo Getter
func (r AlibabaTxcsBrandmarketingCouponStatisticsGetRequest) GetCouponStatisticsParamDo() *CouponStatisticsParamDO {
    return r._couponStatisticsParamDo
}
