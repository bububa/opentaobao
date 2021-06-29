package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索酒店订单列表 APIResponse
alitrip.btrip.hotel.order.search

企业获取商旅酒店订单数据
*/
type AlitripBtripHotelOrderSearchAPIResponse struct {
    model.CommonResponse
    AlitripBtripHotelOrderSearchResponse
}

type AlitripBtripHotelOrderSearchResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_hotel_order_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *BtriphomeResultSupport `json:"result,omitempty" xml:"result,omitempty"`

    
}
