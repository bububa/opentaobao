package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询可用优惠券列表 API返回值 
tmall.promotion.coupon.query

查询用户的可用优惠券列表，仅包含优惠券基本信息和用户nick
*/
type TmallPromotionCouponQueryAPIResponse struct {
    model.CommonResponse
    TmallPromotionCouponQueryAPIResponseModel
}

// 查询可用优惠券列表 成功返回结果
type TmallPromotionCouponQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_promotion_coupon_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TmallPromotionCouponQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
