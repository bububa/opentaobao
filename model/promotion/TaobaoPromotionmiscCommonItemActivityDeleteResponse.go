package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除通用单品优惠活动 APIResponse
taobao.promotionmisc.common.item.activity.delete

删除通用单品优惠活动。
*/
type TaobaoPromotionmiscCommonItemActivityDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPromotionmiscCommonItemActivityDeleteResponse `json:"promotionmisc_common_item_activity_delete_response,omitempty"` 
    TaobaoPromotionmiscCommonItemActivityDeleteResponse
}

/* model for simplify = false
type TaobaoPromotionmiscCommonItemActivityDeleteResponse struct {

    // 是否删除成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoPromotionmiscCommonItemActivityDeleteResponse struct {

    // 是否删除成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
