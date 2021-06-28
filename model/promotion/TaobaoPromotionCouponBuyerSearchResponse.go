package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询买家在相关app领取的优惠券信息 APIResponse
taobao.promotion.coupon.buyer.search

查询买家在相关app领取的优惠券信息
*/
type TaobaoPromotionCouponBuyerSearchAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPromotionCouponBuyerSearchResponse `json:"promotion_coupon_buyer_search_response,omitempty"` 
    TaobaoPromotionCouponBuyerSearchResponse
}

/* model for simplify = false
type TaobaoPromotionCouponBuyerSearchResponse struct {

    // 结果码
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // 调用是否成功
    
    InvokeResult   bool `json:"invoke_result,omitempty"`
    

    // 结果集
    
    BuyerCouponInfos  struct {
        BuyerCouponInfo  []BuyerCouponInfo `json:"buyer_coupon_info,omitempty"`
    } `json:"buyer_coupon_infos,omitempty"`
    

    // 符合条件的总数，用于分页判断
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

}
*/

type TaobaoPromotionCouponBuyerSearchResponse struct {

    // 结果码
    ResultCode   string `json:"result_code,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 调用是否成功
    InvokeResult   bool `json:"invoke_result,omitempty"`

    // 结果集
    BuyerCouponInfos   []BuyerCouponInfo `json:"buyer_coupon_infos,omitempty"`

    // 符合条件的总数，用于分页判断
    TotalCount   int64 `json:"total_count,omitempty"`

}
