package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
在优惠券活动下添加商品。【注意，此接口暂不支持并发！】 APIResponse
alibaba.wdk.marketing.coupon.additem

在优惠券活动下添加商品。【注意，此接口暂不支持并发！】
如果是商品券，则添加的商品为券适用的商品；
如果是品类券，则添加的商品为券排除的商品；
*/
type AlibabaWdkMarketingCouponAdditemAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_marketing_coupon_additem_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品报名活动的返回结果
    
    Result   *MarketResult `json:"result,omitempty" xml:"