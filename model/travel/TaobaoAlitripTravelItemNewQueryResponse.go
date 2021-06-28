package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】新版度假单个商品查询接口 APIResponse
taobao.alitrip.travel.item.new.query

新版旅行度假新商品查询接口（单个商品查询）
*/
type TaobaoAlitripTravelItemNewQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAlitripTravelItemNewQueryResponse `json:"alitrip_travel_item_new_query_response,omitempty"` 
    TaobaoAlitripTravelItemNewQueryResponse
}

/* model for simplify = false
type TaobaoAlitripTravelItemNewQueryResponse struct {

    // 商品查询结果
    
    TravelItem  *struct {
        FullTravelItem  *FullTravelItem `json:"full_travel_item,omitempty"`
    } `json:"travel_item,omitempty"`
    

}
*/

type TaobaoAlitripTravelItemNewQueryResponse struct {

    // 商品查询结果
    TravelItem   *FullTravelItem `json:"travel_item,omitempty"`

}
