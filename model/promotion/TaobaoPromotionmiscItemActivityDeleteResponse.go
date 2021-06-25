package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除无条件单品优惠活动 APIResponse
taobao.promotionmisc.item.activity.delete

删除无条件单品优惠活动
*/
type TaobaoPromotionmiscItemActivityDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoPromotionmiscItemActivityDeleteResponse `json:"taobao_promotionmisc_item_activity_delete_response,omitempty"`
}

type TaobaoPromotionmiscItemActivityDeleteResponse struct {

    // 是否成功删除活动。
    IsSuccess   bool `json:"is_success,omitempty"`

}
