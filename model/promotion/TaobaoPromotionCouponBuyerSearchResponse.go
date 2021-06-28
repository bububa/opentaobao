package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询买家在相关app领取的优惠券信息 APIResponse
taobao.promotion.coupon.buyer.search

查询买家在相关app领取的优惠券信息
*/
type TaobaoPromotionCouponBuyerSearchAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"promotion_coupon_buyer_search_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果码
    
    ResultCode   string `json:"result_code,omitempty" xml:"