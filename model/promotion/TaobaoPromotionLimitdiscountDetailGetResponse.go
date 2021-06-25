package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
限时打折详情查询 APIResponse
taobao.promotion.limitdiscount.detail.get

限时打折详情查询。查询出指定限时打折的对应商品记录信息。
*/
type TaobaoPromotionLimitdiscountDetailGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoPromotionLimitdiscountDetailGetResponse `json:"taobao_promotion_limitdiscount_detail_get_response,omitempty"`
}

type TaobaoPromotionLimitdiscountDetailGetResponse struct {

    // 限时打折对应的商品详情列表。
    ItemDiscountDetailList   []LimitDiscountDetail `json:"item_discount_detail_list,omitempty"`

}
