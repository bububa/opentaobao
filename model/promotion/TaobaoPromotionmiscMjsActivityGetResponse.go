package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询满就送活动 APIResponse
taobao.promotionmisc.mjs.activity.get

查询满就送活动
*/
type TaobaoPromotionmiscMjsActivityGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPromotionmiscMjsActivityGetResponse `json:"promotionmisc_mjs_activity_get_response,omitempty"` 
    TaobaoPromotionmiscMjsActivityGetResponse
}

/* model for simplify = false
type TaobaoPromotionmiscMjsActivityGetResponse struct {

    // 满就送活动信息。
    
    MjsPromotion  *struct {
        MjsPromotion  *MjsPromotion `json:"mjs_promotion,omitempty"`
    } `json:"mjs_promotion,omitempty"`
    

}
*/

type TaobaoPromotionmiscMjsActivityGetResponse struct {

    // 满就送活动信息。
    MjsPromotion   *MjsPromotion `json:"mjs_promotion,omitempty"`

}
