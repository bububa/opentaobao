package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】新版度假单个商品查询接口 APIResponse
taobao.alitrip.travel.item.new.query

新版旅行度假新商品查询接口（单个商品查询）
*/
type TaobaoAlitripTravelItemNewQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_travel_item_new_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品查询结果
    
    TravelItem   *FullTravelItem `json:"travel_item,omitempty" xml:"