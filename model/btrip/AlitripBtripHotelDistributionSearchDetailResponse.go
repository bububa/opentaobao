package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅酒店api分销-详情报价接口 APIResponse
alitrip.btrip.hotel.distribution.search.detail

商旅酒店api分销-详情报价接口
*/
type AlitripBtripHotelDistributionSearchDetailAPIResponse struct {
    model.CommonResponse
    AlitripBtripHotelDistributionSearchDetailResponse
}

type AlitripBtripHotelDistributionSearchDetailResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_hotel_distribution_search_detail_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 详情报价回参
    
    Result   *HisvResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
