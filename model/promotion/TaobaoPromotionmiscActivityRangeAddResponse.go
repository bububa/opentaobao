package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
增加活动参与的商品 APIResponse
taobao.promotionmisc.activity.range.add

增加活动参与的商品，部分商品参与的活动，最大支持指定150个商品。
*/
type TaobaoPromotionmiscActivityRangeAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoPromotionmiscActivityRangeAddResponse `json:"taobao_promotionmisc_activity_range_add_response,omitempty"`
}

type TaobaoPromotionmiscActivityRangeAddResponse struct {

    // 增加商品范围是否成功。
    IsSuccess   bool `json:"is_success,omitempty"`

}
