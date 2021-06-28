package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】度假线路商品上下架接口 APIResponse
taobao.alitrip.travel.item.shelve

旅行度假新商品发布接口 第三版：度假商品上下架接口
注意：定时上下架功能，目前只支持接送、租车类目。
*/
type TaobaoAlitripTravelItemShelveAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_travel_item_shelve_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品上下架操作是否成功
    
    Result   bool `json:"result,omitempty" xml:"