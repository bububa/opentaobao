package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询优惠券活动 API请求
alibaba.wdk.marketing.coupon.queryactivity

查询优惠券活动
*/
type AlibabaWdkMarketingCouponQueryactivityRequest struct {
    model.Params
    // 查询优惠券活动入参
    param0   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingCouponQueryactivityRequest对象
func NewAlibabaWdkMarketingCouponQueryactivityRequest() *AlibabaWdkMarketingCouponQueryactivityRequest{
    return &AlibabaWdkMarketingCouponQueryactivityRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingCouponQueryactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.coupon.queryactivity"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingCouponQueryactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 查询优惠券活动入参
func (r *AlibabaWdkMarketingCouponQueryactivityRequest) SetParam0(param0 *CommonActivityParam) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingCouponQueryactivityRequest) GetParam0() *CommonActivityParam {
    return r.param0
}
