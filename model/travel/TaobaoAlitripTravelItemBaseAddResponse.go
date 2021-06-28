package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】度假线路商品发布接口 APIResponse
taobao.alitrip.travel.item.base.add

旅行度假新商品发布接口。目前支持的类目包括：境内跟团游、出境跟团游、境内自由行、出境自由行、境内当地玩乐、境外玩乐套餐、境内邮轮、国际邮轮
*/
type TaobaoAlitripTravelItemBaseAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAlitripTravelItemBaseAddResponse `json:"alitrip_travel_item_base_add_response,omitempty"` 
    TaobaoAlitripTravelItemBaseAddResponse
}

/* model for simplify = false
type TaobaoAlitripTravelItemBaseAddResponse struct {

    // 商品发布结果
    
    TravelItem  *struct {
        TopTravelItem  *TopTravelItem `json:"top_travel_item,omitempty"`
    } `json:"travel_item,omitempty"`
    

}
*/

type TaobaoAlitripTravelItemBaseAddResponse struct {

    // 商品发布结果
    TravelItem   *TopTravelItem `json:"travel_item,omitempty"`

}
