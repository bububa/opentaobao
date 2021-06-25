package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
普通发券 APIRequest
alibaba.wdk.coupon.spread.apply

优惠券发放
*/
type AlibabaWdkCouponSpreadApplyRequest struct {
    model.Params

    // 参数对象
    paramWdkCouponApplyParam   *WdkCouponApplyParam 

}

func NewAlibabaWdkCouponSpreadApplyRequest() *AlibabaWdkCouponSpreadApplyRequest{
    return &AlibabaWdkCouponSpreadApplyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkCouponSpreadApplyRequest) GetApiMethodName() string {
    return "alibaba.wdk.coupon.spread.apply"
}

func (r AlibabaWdkCouponSpreadApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkCouponSpreadApplyRequest) SetParamWdkCouponApplyParam(paramWdkCouponApplyParam *WdkCouponApplyParam) error {
    r.paramWdkCouponApplyParam = paramWdkCouponApplyParam
    r.Set("param_wdk_coupon_apply_param", paramWdkCouponApplyParam)
    return nil
}

func (r AlibabaWdkCouponSpreadApplyRequest) GetParamWdkCouponApplyParam() *WdkCouponApplyParam {
    return r.paramWdkCouponApplyParam
}

