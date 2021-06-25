package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
（供销）产品级别日历价格库存修改，全量覆盖 APIResponse
taobao.alitrip.travel.product.sku.override

（供销）产品级别日历价格库存修改，全量覆盖
*/
type TaobaoAlitripTravelProductSkuOverrideAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAlitripTravelProductSkuOverrideResponse `json:"taobao_alitrip_travel_product_sku_override_response,omitempty"`
}

type TaobaoAlitripTravelProductSkuOverrideResponse struct {

    // 商品sku修改结果
    TravelItem   *TopTravelItem `json:"travel_item,omitempty"`

}
