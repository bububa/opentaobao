package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】商品级别日历价格库存修改，全量覆盖 APIResponse
taobao.alitrip.travel.item.sku.override

旅行度假新商品日历价格库存信息修改接口 第三版。提供商家通过TOP API方式修改商品sku信息。
*/
type TaobaoAlitripTravelItemSkuOverrideAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAlitripTravelItemSkuOverrideResponse `json:"taobao_alitrip_travel_item_sku_override_response,omitempty"`
}

type TaobaoAlitripTravelItemSkuOverrideResponse struct {

    // 商品sku修改结果
    TravelItem   *TopTravelItem `json:"travel_item,omitempty"`

}
