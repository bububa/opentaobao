package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
券核销接口 APIResponse
tmall.promotion.coupon.use

核销用户的一张优惠券，返回核销结果
*/
type TmallPromotionCouponUseAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_promotion_coupon_use_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // data
    
    Data   *UseResultDo `json:"data,omitempty" xml:"