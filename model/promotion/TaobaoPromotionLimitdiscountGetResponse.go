package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
限时打折查询 APIResponse
taobao.promotion.limitdiscount.get

分页查询某个卖家的限时打折信息。每页20条数据，按照结束时间降序排列。也可指定某一个限时打折id查询唯一的限时打折信息。
*/
type TaobaoPromotionLimitdiscountGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPromotionLimitdiscountGetResponse `json:"promotion_limitdiscount_get_response,omitempty"` 
    TaobaoPromotionLimitdiscountGetResponse
}

/* model for simplify = false
type TaobaoPromotionLimitdiscountGetResponse struct {

    // 满足该查询条件的限时打折总数量。
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

    // 限时打折列表。
    
    LimitDiscountList  struct {
        LimitDiscount  []LimitDiscount `json:"limit_discount,omitempty"`
    } `json:"limit_discount_list,omitempty"`
    

}
*/

type TaobaoPromotionLimitdiscountGetResponse struct {

    // 满足该查询条件的限时打折总数量。
    TotalCount   int64 `json:"total_count,omitempty"`

    // 限时打折列表。
    LimitDiscountList   []LimitDiscount `json:"limit_discount_list,omitempty"`

}
