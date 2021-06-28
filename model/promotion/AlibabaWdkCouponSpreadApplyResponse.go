package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
普通发券 APIResponse
alibaba.wdk.coupon.spread.apply

优惠券发放
*/
type AlibabaWdkCouponSpreadApplyAPIResponse struct {
    model.CommonResponse
    AlibabaWdkCouponSpreadApplyResponse
}

type AlibabaWdkCouponSpreadApplyResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_coupon_spread_apply_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *AlibabaWdkCouponSpreadApplyApiResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
