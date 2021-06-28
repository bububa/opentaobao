package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券活动创建 APIResponse
alibaba.wdk.marketing.coupon.createactivity

添加优惠券活动
*/
type AlibabaWdkMarketingCouponCreateactivityAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingCouponCreateactivityResponse
}

type AlibabaWdkMarketingCouponCreateactivityResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_coupon_createactivity_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 创建优惠券活动返回结果
    
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
