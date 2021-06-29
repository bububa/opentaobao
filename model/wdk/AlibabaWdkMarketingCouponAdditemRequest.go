package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
在优惠券活动下添加商品。【注意，此接口暂不支持并发！】 API请求
alibaba.wdk.marketing.coupon.additem

在优惠券活动下添加商品。【注意，此接口暂不支持并发！】
如果是商品券，则添加的商品为券适用的商品；
如果是品类券，则添加的商品为券排除的商品；
*/
type AlibabaWdkMarketingCouponAdditemRequest struct {
    model.Params
    // 商品对象
    param0   *ItemCouponSku
    // 活动基本信息
    param1   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingCouponAdditemRequest对象
func NewAlibabaWdkMarketingCouponAdditemRequest() *AlibabaWdkMarketingCouponAdditemRequest{
    return &AlibabaWdkMarketingCouponAdditemRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingCouponAdditemRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.coupon.additem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingCouponAdditemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 商品对象
func (r *AlibabaWdkMarketingCouponAdditemRequest) SetParam0(param0 *ItemCouponSku) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingCouponAdditemRequest) GetParam0() *ItemCouponSku {
    return r.param0
}
// Param1 Setter
// 活动基本信息
func (r *AlibabaWdkMarketingCouponAdditemRequest) SetParam1(param1 *CommonActivityParam) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r AlibabaWdkMarketingCouponAdditemRequest) GetParam1() *CommonActivityParam {
    return r.param1
}
