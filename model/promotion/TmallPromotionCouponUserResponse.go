package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户信息查询接口 APIResponse
tmall.promotion.coupon.user

开发给外部合作商（例如：苏宁），通过会员付款码获得会员nick
*/
type TmallPromotionCouponUserAPIResponse struct {
    model.CommonResponse
    TmallPromotionCouponUserResponse
}

type TmallPromotionCouponUserResponse struct {
    XMLName xml.Name `xml:"tmall_promotion_coupon_user_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TmallPromotionCouponUserResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
