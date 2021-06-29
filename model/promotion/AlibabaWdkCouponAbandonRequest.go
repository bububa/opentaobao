package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
废券 API请求
alibaba.wdk.coupon.abandon

优惠券废弃
*/
type AlibabaWdkCouponAbandonRequest struct {
    model.Params
    // 废券参数
    paramWdkCouponAbandonParam   *WdkCouponAbandonParam
}

// 初始化AlibabaWdkCouponAbandonRequest对象
func NewAlibabaWdkCouponAbandonRequest() *AlibabaWdkCouponAbandonRequest{
    return &AlibabaWdkCouponAbandonRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponAbandonRequest) GetApiMethodName() string {
    return "alibaba.wdk.coupon.abandon"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponAbandonRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamWdkCouponAbandonParam Setter
// 废券参数
func (r *AlibabaWdkCouponAbandonRequest) SetParamWdkCouponAbandonParam(paramWdkCouponAbandonParam *WdkCouponAbandonParam) error {
    r.paramWdkCouponAbandonParam = paramWdkCouponAbandonParam
    r.Set("param_wdk_coupon_abandon_param", paramWdkCouponAbandonParam)
    return nil
}

// ParamWdkCouponAbandonParam Getter
func (r AlibabaWdkCouponAbandonRequest) GetParamWdkCouponAbandonParam() *WdkCouponAbandonParam {
    return r.paramWdkCouponAbandonParam
}
