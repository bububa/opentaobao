package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券活动创建 API返回值 
alibaba.wdk.marketing.coupon.createactivity

添加优惠券活动
*/
type AlibabaWdkMarketingCouponCreateactivityAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingCouponCreateactivityResponse
}

// 优惠券活动创建 成功返回结果
type AlibabaWdkMarketingCouponCreateactivityResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_coupon_createactivity_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 创建优惠券活动返回结果
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
