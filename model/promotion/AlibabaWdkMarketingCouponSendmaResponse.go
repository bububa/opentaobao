package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发放匿名码 APIResponse
alibaba.wdk.marketing.coupon.sendma

根据优惠券活动id打印单个匿名码
*/
type AlibabaWdkMarketingCouponSendmaAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingCouponSendmaResponse `json:"alibaba_wdk_marketing_coupon_sendma_response,omitempty"`
}

type AlibabaWdkMarketingCouponSendmaResponse struct {

    // 发放匿名码返回结果
    Result   *MarketResult `json:"result,omitempty"`

}
