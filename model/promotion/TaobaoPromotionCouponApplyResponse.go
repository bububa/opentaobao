package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券领取 APIResponse
taobao.promotion.coupon.apply

优惠券领取
*/
type TaobaoPromotionCouponApplyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"promotion_coupon_apply_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 失败详细描述
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"