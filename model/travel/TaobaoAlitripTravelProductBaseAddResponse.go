package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
供应商新增产品API APIResponse
taobao.alitrip.travel.product.base.add

飞猪供销平台供应商可通过该API发布新产品
*/
type TaobaoAlitripTravelProductBaseAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAlitripTravelProductBaseAddResponse `json:"taobao_alitrip_travel_product_base_add_response,omitempty"`
}

type TaobaoAlitripTravelProductBaseAddResponse struct {

    // 商品发布结果
    TravelItem   *TopTravelItem `json:"travel_item,omitempty"`

}
