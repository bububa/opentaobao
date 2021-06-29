package alitripmerchant

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-酒店列表页搜索 APIResponse
alitrip.merchant.galaxy.hotel.list.search

星河产品=酒店列表页搜索
*/
type AlitripMerchantGalaxyHotelListSearchAPIResponse struct {
    model.CommonResponse
    AlitripMerchantGalaxyHotelListSearchResponse
}

type AlitripMerchantGalaxyHotelListSearchResponse struct {
    XMLName xml.Name `xml:"alitrip_merchant_galaxy_hotel_list_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *PageableResponse `json:"result,omitempty" xml:"result,omitempty"`

    
}
