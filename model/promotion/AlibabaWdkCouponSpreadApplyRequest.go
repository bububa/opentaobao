package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
普通发券 API请求
alibaba.wdk.coupon.spread.apply

优惠券发放
*/
type AlibabaWdkCouponSpreadApplyRequest struct {
    model.Params
    // 参数对象
    _paramWdkCouponApplyParam   *WdkCouponApplyParam
}

// 初始化AlibabaWdkCouponSpreadApplyRequest对象
func NewAlibabaWdkCouponSpreadApplyRequest() *AlibabaWdkCouponSpreadApplyRequest{
    return &AlibabaWdkCouponSpreadApplyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkCouponSpreadApplyRequest) GetApiMethodName() string {
    return "alibaba.wdk.coupon.spread.apply"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkCouponSpreadApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamWdkCouponApplyParam Setter
// 参数对象
func (r *AlibabaWdkCouponSpreadApplyRequest) SetParamWdkCouponApplyParam(_paramWdkCouponApplyParam *WdkCouponApplyParam) error {
    r._paramWdkCouponApplyParam = _paramWdkCouponApplyParam
    r.Set("param_wdk_coupon_apply_param", _paramWdkCouponApplyParam)
    return nil
}

// ParamWdkCouponApplyParam Getter
func (r AlibabaWdkCouponSpreadApplyRequest) GetParamWdkCouponApplyParam() *WdkCouponApplyParam {
    return r._paramWdkCouponApplyParam
}
