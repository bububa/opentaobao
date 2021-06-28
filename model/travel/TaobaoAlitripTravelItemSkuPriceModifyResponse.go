package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】日期级别日历价格库存修改，增量维护 APIResponse
taobao.alitrip.travel.item.sku.price.modify

【API3.0】日期级别日历价格库存增量维护
*/
type TaobaoAlitripTravelItemSkuPriceModifyAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAlitripTravelItemSkuPriceModifyResponse `json:"alitrip_travel_item_sku_price_modify_response,omitempty"` 
    TaobaoAlitripTravelItemSkuPriceModifyResponse
}

/* model for simplify = false
type TaobaoAlitripTravelItemSkuPriceModifyResponse struct {

    // 日期级别日历价格库存增量维护
    
    TravelItem  *struct {
        TopTravelItem  *TopTravelItem `json:"top_travel_item,omitempty"`
    } `json:"travel_item,omitempty"`
    

}
*/

type TaobaoAlitripTravelItemSkuPriceModifyResponse struct {

    // 日期级别日历价格库存增量维护
    TravelItem   *TopTravelItem `json:"travel_item,omitempty"`

}
