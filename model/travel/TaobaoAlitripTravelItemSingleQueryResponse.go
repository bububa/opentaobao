package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】度假单个商品查询接口 APIResponse
taobao.alitrip.travel.item.single.query

旅行度假新商品查询接口（单个商品查询） 第三版
*/
type TaobaoAlitripTravelItemSingleQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAlitripTravelItemSingleQueryResponse `json:"taobao_alitrip_travel_item_single_query_response,omitempty"`
}

type TaobaoAlitripTravelItemSingleQueryResponse struct {

    // 商品查询结果
    TravelItem   *PontusTravelFullTravelItem `json:"travel_item,omitempty"`

}
