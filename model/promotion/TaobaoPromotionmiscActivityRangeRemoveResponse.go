package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
去除活动参与的商品 APIResponse
taobao.promotionmisc.activity.range.remove

去除活动参与的商品
*/
type TaobaoPromotionmiscActivityRangeRemoveAPIResponse struct {
    model.CommonResponse
    Response *TaobaoPromotionmiscActivityRangeRemoveResponse `json:"taobao_promotionmisc_activity_range_remove_response,omitempty"`
}

type TaobaoPromotionmiscActivityRangeRemoveResponse struct {

    // 去除活动参与的商品是否成功。
    IsSuccess   bool `json:"is_success,omitempty"`

}
