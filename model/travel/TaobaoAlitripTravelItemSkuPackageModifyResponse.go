package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】套餐级别日历价格库存增删操作 APIResponse
taobao.alitrip.travel.item.sku.package.modify

【API3.0】套餐级别日历价格库存增删操作
*/
type TaobaoAlitripTravelItemSkuPackageModifyAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAlitripTravelItemSkuPackageModifyResponse `json:"alitrip_travel_item_sku_package_modify_response,omitempty"` 
    TaobaoAlitripTravelItemSkuPackageModifyResponse
}

/* model for simplify = false
type TaobaoAlitripTravelItemSkuPackageModifyResponse struct {

    // 商品sku修改结果
    
    TravelItem  *struct {
        TopTravelItem  *TopTravelItem `json:"top_travel_item,omitempty"`
    } `json:"travel_item,omitempty"`
    

}
*/

type TaobaoAlitripTravelItemSkuPackageModifyResponse struct {

    // 商品sku修改结果
    TravelItem   *TopTravelItem `json:"travel_item,omitempty"`

}
