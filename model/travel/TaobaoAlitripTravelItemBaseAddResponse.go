package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】度假线路商品发布接口 APIResponse
taobao.alitrip.travel.item.base.add

旅行度假新商品发布接口。目前支持的类目包括：境内跟团游、出境跟团游、境内自由行、出境自由行、境内当地玩乐、境外玩乐套餐、境内邮轮、国际邮轮
*/
type TaobaoAlitripTravelItemBaseAddAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripTravelItemBaseAddResponse
}

type TaobaoAlitripTravelItemBaseAddResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_item_base_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 商品发布结果
    
    TravelItem   *TopTravelItem `json:"travel_item,omitempty" xml:"travel_item,omitempty"`

    
}
