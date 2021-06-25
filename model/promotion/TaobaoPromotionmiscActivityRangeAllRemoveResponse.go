package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
清空活动参与的商品 APIResponse
taobao.promotionmisc.activity.range.all.remove

清空活动参与的商品
*/
type TaobaoPromotionmiscActivityRangeAllRemoveAPIResponse struct {
    model.CommonResponse
    Response *TaobaoPromotionmiscActivityRangeAllRemoveResponse `json:"taobao_promotionmisc_activity_range_all_remove_response,omitempty"`
}

type TaobaoPromotionmiscActivityRangeAllRemoveResponse struct {

    // 清空活动参与商品是否成功。
    IsSuccess   bool `json:"is_success,omitempty"`

}
