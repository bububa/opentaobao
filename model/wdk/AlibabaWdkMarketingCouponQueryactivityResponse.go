package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询优惠券活动 APIResponse
alibaba.wdk.marketing.coupon.queryactivity

查询优惠券活动
*/
type AlibabaWdkMarketingCouponQueryactivityAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingCouponQueryactivityResponse
}

type AlibabaWdkMarketingCouponQueryactivityResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_coupon_queryactivity_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询返回结果
    
    Result   *MarketResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
