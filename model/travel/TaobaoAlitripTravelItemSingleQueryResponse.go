package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】度假单个商品查询接口 APIResponse
taobao.alitrip.travel.item.single.query

旅行度假新商品查询接口（单个商品查询） 第三版
*/
type TaobaoAlitripTravelItemSingleQueryAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripTravelItemSingleQueryResponse
}

type TaobaoAlitripTravelItemSingleQueryResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_item_single_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 商品查询结果
    
    TravelItem   *PontusTravelFullTravelItem `json:"travel_item,omitempty" xml:"travel_item,omitempty"`

    
}
