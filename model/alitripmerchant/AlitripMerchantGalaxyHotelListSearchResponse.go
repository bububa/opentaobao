package alitripmerchant

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-酒店列表页搜索 API返回值 
alitrip.merchant.galaxy.hotel.list.search

星河产品=酒店列表页搜索
*/
type AlitripMerchantGalaxyHotelListSearchAPIResponse struct {
    model.CommonResponse
    AlitripMerchantGalaxyHotelListSearchResponse
}

// 星河-酒店列表页搜索 成功返回结果
type AlitripMerchantGalaxyHotelListSearchResponse struct {
    XMLName xml.Name `xml:"alitrip_merchant_galaxy_hotel_list_search_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *PageableResponse `json:"result,omitempty" xml:"result,omitempty"`
}
