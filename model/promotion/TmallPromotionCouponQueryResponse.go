package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询可用优惠券列表 APIResponse
tmall.promotion.coupon.query

查询用户的可用优惠券列表，仅包含优惠券基本信息和用户nick
*/
type TmallPromotionCouponQueryAPIResponse struct {
    model.CommonResponse
    TmallPromotionCouponQueryResponse
}

type TmallPromotionCouponQueryResponse struct {
    XMLName xml.Name `xml:"tmall_promotion_coupon_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TmallPromotionCouponQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
