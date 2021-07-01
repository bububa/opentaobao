package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户信息查询接口 API返回值 
tmall.promotion.coupon.user

开发给外部合作商（例如：苏宁），通过会员付款码获得会员nick
*/
type TmallPromotionCouponUserAPIResponse struct {
    model.CommonResponse
    TmallPromotionCouponUserAPIResponseModel
}

// 用户信息查询接口 成功返回结果
type TmallPromotionCouponUserAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_promotion_coupon_user_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TmallPromotionCouponUserResult `json:"result,omitempty" xml:"result,omitempty"`
}
