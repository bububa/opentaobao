package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发放匿名码 APIResponse
alibaba.wdk.marketing.coupon.sendma

根据优惠券活动id打印单个匿名码
*/
type AlibabaWdkMarketingCouponSendmaAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_marketing_coupon_sendma_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 发放匿名码返回结果
    
    Result   *MarketResult `json:"result,omitempty" xml:"