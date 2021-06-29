package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发放匿名码 API请求
alibaba.wdk.marketing.coupon.sendma

根据优惠券活动id打印单个匿名码
*/
type AlibabaWdkMarketingCouponSendmaRequest struct {
    model.Params
    // 发放匿名码入参
    _param0   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingCouponSendmaRequest对象
func NewAlibabaWdkMarketingCouponSendmaRequest() *AlibabaWdkMarketingCouponSendmaRequest{
    return &AlibabaWdkMarketingCouponSendmaRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingCouponSendmaRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.coupon.sendma"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingCouponSendmaRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 发放匿名码入参
func (r *AlibabaWdkMarketingCouponSendmaRequest) SetParam0(_param0 *CommonActivityParam) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingCouponSendmaRequest) GetParam0() *CommonActivityParam {
    return r._param0
}
