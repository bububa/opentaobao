package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
废券 APIResponse
alibaba.wdk.coupon.abandon

优惠券废弃
*/
type AlibabaWdkCouponAbandonAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_coupon_abandon_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *AlibabaWdkCouponAbandonApiResult `json:"result,omitempty" xml:"