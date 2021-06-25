package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
供应商编辑产品API APIResponse
taobao.alitrip.travel.product.base.modify

飞猪供销平台供应商可通过该API编辑产品
*/
type TaobaoAlitripTravelProductBaseModifyAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAlitripTravelProductBaseModifyResponse `json:"taobao_alitrip_travel_product_base_modify_response,omitempty"`
}

type TaobaoAlitripTravelProductBaseModifyResponse struct {

    // 商品修改结果
    TravelItem   *TopTravelItem `json:"travel_item,omitempty"`

}
