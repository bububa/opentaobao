package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交专车优惠券活动 APIRequest
alibaba.alihouse.newhome.coupon.submit

提交专车优惠券活动
*/
type AlibabaAlihouseNewhomeCouponSubmitRequest struct {
    model.Params

    // 打车券活动对象
    couponDto   *MarketingCouponDto 

}

func NewAlibabaAlihouseNewhomeCouponSubmitRequest() *AlibabaAlihouseNewhomeCouponSubmitRequest{
    return &AlibabaAlihouseNewhomeCouponSubmitRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeCouponSubmitRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.coupon.submit"
}

func (r AlibabaAlihouseNewhomeCouponSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeCouponSubmitRequest) SetCouponDto(couponDto *MarketingCouponDto) error {
    r.couponDto = couponDto
    r.Set("coupon_dto", couponDto)
    return nil
}

func (r AlibabaAlihouseNewhomeCouponSubmitRequest) GetCouponDto() *MarketingCouponDto {
    return r.couponDto
}

