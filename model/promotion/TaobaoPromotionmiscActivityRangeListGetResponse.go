package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询活动参与的商品 APIResponse
taobao.promotionmisc.activity.range.list.get

查询活动参与的商品
*/
type TaobaoPromotionmiscActivityRangeListGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPromotionmiscActivityRangeListGetResponse `json:"promotionmisc_activity_range_list_get_response,omitempty"` 
    TaobaoPromotionmiscActivityRangeListGetResponse
}

/* model for simplify = false
type TaobaoPromotionmiscActivityRangeListGetResponse struct {

    // 活动参与的商品列表
    
    PromotionRangeList  struct {
        PromotionRange  []PromotionRange `json:"promotion_range,omitempty"`
    } `json:"promotion_range_list,omitempty"`
    

}
*/

type TaobaoPromotionmiscActivityRangeListGetResponse struct {

    // 活动参与的商品列表
    PromotionRangeList   []PromotionRange `json:"promotion_range_list,omitempty"`

}
