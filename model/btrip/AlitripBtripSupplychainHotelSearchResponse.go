package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】酒店订单查询 API返回值 
alitrip.btrip.supplychain.hotel.search

【商旅】酒店订单查询
*/
type AlitripBtripSupplychainHotelSearchAPIResponse struct {
    model.CommonResponse
    AlitripBtripSupplychainHotelSearchResponse
}

// 【商旅】酒店订单查询 成功返回结果
type AlitripBtripSupplychainHotelSearchResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_supplychain_hotel_search_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 出参
    Result   *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
