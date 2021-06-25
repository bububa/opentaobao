package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品优惠详情查询 APIResponse
taobao.ump.promotion.get

商品优惠详情查询，可查询商品设置的详细优惠。包括限时折扣，满就送等官方优惠以及第三方优惠。
*/
type TaobaoUmpPromotionGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoUmpPromotionGetResponse `json:"taobao_ump_promotion_get_response,omitempty"`
}

type TaobaoUmpPromotionGetResponse struct {

    // 优惠详细信息
    Promotions   *PromotionDisplayTop `json:"promotions,omitempty"`

}
