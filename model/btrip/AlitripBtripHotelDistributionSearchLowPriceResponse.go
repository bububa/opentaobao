package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅酒店api分销-酒店最低价 APIResponse
alitrip.btrip.hotel.distribution.search.low.price

商旅酒店api分销-酒店最低价
*/
type AlitripBtripHotelDistributionSearchLowPriceAPIResponse struct {
    model.CommonResponse
    AlitripBtripHotelDistributionSearchLowPriceResponse
}

type AlitripBtripHotelDistributionSearchLowPriceResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_hotel_distribution_search_low_price_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回出参
    
    Result   *HisvResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
