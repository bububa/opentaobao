package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券领取鉴权接口 APIResponse
alibaba.interact.coupon.apply

鉴权接口，为coupon.apply接口鉴权
*/
type AlibabaInteractCouponApplyAPIResponse struct {
    model.CommonResponse
    AlibabaInteractCouponApplyResponse
}

type AlibabaInteractCouponApplyResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_coupon_apply_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 无用参数，top限制一定要有出参，增加此参数
    
    Stub   string `json:"stub,omitempty" xml:"stub,omitempty"`

    
}
