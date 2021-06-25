package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌营销导购员券推广统计数据回流 APIRequest
alibaba.txcs.brandmarketing.coupon.statistics.get

请求券统计数据回流
*/
type AlibabaTxcsBrandmarketingCouponStatisticsGetRequest struct {
    model.Params

    // 请求信息
    couponStatisticsParamDo   *CouponStatisticsParamDO 

}

func NewAlibabaTxcsBrandmarketingCouponStatisticsGetRequest() *AlibabaTxcsBrandmarketingCouponStatisticsGetRequest{
    return &AlibabaTxcsBrandmarketingCouponStatisticsGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTxcsBrandmarketingCouponStatisticsGetRequest) GetApiMethodName() string {
    return "alibaba.txcs.brandmarketing.coupon.statistics.get"
}

func (r AlibabaTxcsBrandmarketingCouponStatisticsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTxcsBrandmarketingCouponStatisticsGetRequest) SetCouponStatisticsParamDo(couponStatisticsParamDo *CouponStatisticsParamDO) error {
    r.couponStatisticsParamDo = couponStatisticsParamDo
    r.Set("coupon_statistics_param_do", couponStatisticsParamDo)
    return nil
}

func (r AlibabaTxcsBrandmarketingCouponStatisticsGetRequest) GetCouponStatisticsParamDo() *CouponStatisticsParamDO {
    return r.couponStatisticsParamDo
}

