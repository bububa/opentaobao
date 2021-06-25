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
    Response *TaobaoPromotionmiscItemActivityGetResponse `json:"taobao_promotionmisc_item_activity_get_response,omitempty"`
}

type TaobaoPromotionmiscItemActivityGetResponse struct {

    // 单品优惠活动信息。
    ItemPromotion   *ItemPromotion `json:"item_promotion,omitempty"`

}
