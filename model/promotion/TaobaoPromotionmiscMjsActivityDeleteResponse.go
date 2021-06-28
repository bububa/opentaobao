package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除满就送活动 APIResponse
taobao.promotionmisc.mjs.activity.delete

删除满就送活动
*/
type TaobaoPromotionmiscMjsActivityDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPromotionmiscMjsActivityDeleteResponse `json:"promotionmisc_mjs_activity_delete_response,omitempty"` 
    TaobaoPromotionmiscMjsActivityDeleteResponse
}

/* model for simplify = false
type TaobaoPromotionmiscMjsActivityDeleteResponse struct {

    // 是否成功删除活动。
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoPromotionmiscMjsActivityDeleteResponse struct {

    // 是否成功删除活动。
    IsSuccess   bool `json:"is_success,omitempty"`

}
