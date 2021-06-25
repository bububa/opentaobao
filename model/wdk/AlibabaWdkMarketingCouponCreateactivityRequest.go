package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券活动创建 APIRequest
alibaba.wdk.marketing.coupon.createactivity

添加优惠券活动
*/
type AlibabaWdkMarketingCouponCreateactivityRequest struct {
    model.Params

    // 创建优惠券活动请求入参
    param   *CouponActivity 

}

func NewAlibabaWdkMarketingCouponCreateactivityRequest() *AlibabaWdkMarketingCouponCreateactivityRequest{
    return &AlibabaWdkMarketingCouponCreateactivityRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMarketingCouponCreateactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.coupon.createactivity"
}

func (r AlibabaWdkMarketingCouponCreateactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMarketingCouponCreateactivityRequest) SetParam(param *CouponActivity) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkMarketingCouponCreateactivityRequest) GetParam() *CouponActivity {
    return r.param
}

