package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询无条件单品优惠活动 APIResponse
taobao.promotionmisc.item.activity.get

查询无条件单品优惠活动
*/
type TaobaoPromotionmiscItemActivityGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPromotionmiscItemActivityGetResponse `json:"promotionmisc_item_activity_get_response,omitempty"` 
    TaobaoPromotionmiscItemActivityGetResponse
}

/* model for simplify = false
type TaobaoPromotionmiscItemActivityGetResponse struct {

    // 单品优惠活动信息。
    
    ItemPromotion  *struct {
        ItemPromotion  *ItemPromotion `json:"item_promotion,omitempty"`
    } `json:"item_promotion,omitempty"`
    

}
*/

type TaobaoPromotionmiscItemActivityGetResponse struct {

    // 单品优惠活动信息。
    ItemPromotion   *ItemPromotion `json:"item_promotion,omitempty"`

}
