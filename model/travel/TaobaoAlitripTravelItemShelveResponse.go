package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】度假线路商品上下架接口 API返回值 
taobao.alitrip.travel.item.shelve

旅行度假新商品发布接口 第三版：度假商品上下架接口
注意：定时上下架功能，目前只支持接送、租车类目。
*/
type TaobaoAlitripTravelItemShelveAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripTravelItemShelveResponse
}

// 【API3.0】度假线路商品上下架接口 成功返回结果
type TaobaoAlitripTravelItemShelveResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_item_shelve_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 商品上下架操作是否成功
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
}
