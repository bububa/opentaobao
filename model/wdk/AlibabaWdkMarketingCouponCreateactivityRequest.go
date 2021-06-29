package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券活动创建 API请求
alibaba.wdk.marketing.coupon.createactivity

添加优惠券活动
*/
type AlibabaWdkMarketingCouponCreateactivityRequest struct {
    model.Params
    // 创建优惠券活动请求入参
    param   *CouponActivity
}

// 初始化AlibabaWdkMarketingCouponCreateactivityRequest对象
func NewAlibabaWdkMarketingCouponCreateactivityRequest() *AlibabaWdkMarketingCouponCreateactivityRequest{
    return &AlibabaWdkMarketingCouponCreateactivityRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingCouponCreateactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.coupon.createactivity"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingCouponCreateactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 创建优惠券活动请求入参
func (r *AlibabaWdkMarketingCouponCreateactivityRequest) SetParam(param *CouponActivity) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaWdkMarketingCouponCreateactivityRequest) GetParam() *CouponActivity {
    return r.param
}
