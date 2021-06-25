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
    Response *TaobaoPromotionmiscMjsActivityGetResponse `json:"taobao_promotionmisc_mjs_activity_get_response,omitempty"`
}

type TaobaoPromotionmiscMjsActivityGetResponse struct {

    // 满就送活动信息。
    MjsPromotion   *MjsPromotion `json:"mjs_promotion,omitempty"`

}
