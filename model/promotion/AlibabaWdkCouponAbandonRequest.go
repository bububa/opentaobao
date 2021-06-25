package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
废券 APIRequest
alibaba.wdk.coupon.abandon

优惠券废弃
*/
type AlibabaWdkCouponAbandonRequest struct {
    model.Params

    // 废券参数
    paramWdkCouponAbandonParam   *WdkCouponAbandonParam 

}

func NewAlibabaWdkCouponAbandonRequest() *AlibabaWdkCouponAbandonRequest{
    return &AlibabaWdkCouponAbandonRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkCouponAbandonRequest) GetApiMethodName() string {
    return "alibaba.wdk.coupon.abandon"
}

func (r AlibabaWdkCouponAbandonRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkCouponAbandonRequest) SetParamWdkCouponAbandonParam(paramWdkCouponAbandonParam *WdkCouponAbandonParam) error {
    r.paramWdkCouponAbandonParam = paramWdkCouponAbandonParam
    r.Set("param_wdk_coupon_abandon_param", paramWdkCouponAbandonParam)
    return nil
}

func (r AlibabaWdkCouponAbandonRequest) GetParamWdkCouponAbandonParam() *WdkCouponAbandonParam {
    return r.paramWdkCouponAbandonParam
}

