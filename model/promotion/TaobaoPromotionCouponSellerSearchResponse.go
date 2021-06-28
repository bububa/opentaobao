package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询绑定卖家优惠券相关信息 APIResponse
taobao.promotion.coupon.seller.search

查询绑定卖家相关优惠券信息  如isv  百川 等外部业务方
*/
type TaobaoPromotionCouponSellerSearchAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPromotionCouponSellerSearchResponse `json:"promotion_coupon_seller_search_response,omitempty"` 
    TaobaoPromotionCouponSellerSearchResponse
}

/* model for simplify = false
type TaobaoPromotionCouponSellerSearchResponse struct {

    // 调用错误码，只有调用失败的时候才会有
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // 失败详细描述
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // 接口调用结果，调用成功为true，否则为false
    
    InvokeResult   bool `json:"invoke_result,omitempty"`
    

    // 结果
    
    SellerCouponDetails  struct {
        SellerCouponDetail  []SellerCouponDetail `json:"seller_coupon_detail,omitempty"`
    } `json:"seller_coupon_details,omitempty"`
    

    // 符合条件总数量，用于分页等判断
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

}
*/

type TaobaoPromotionCouponSellerSearchResponse struct {

    // 调用错误码，只有调用失败的时候才会有
    ResultCode   string `json:"result_code,omitempty"`

    // 失败详细描述
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 接口调用结果，调用成功为true，否则为false
    InvokeResult   bool `json:"invoke_result,omitempty"`

    // 结果
    SellerCouponDetails   []SellerCouponDetail `json:"seller_coupon_details,omitempty"`

    // 符合条件总数量，用于分页等判断
    TotalCount   int64 `json:"total_count,omitempty"`

}
