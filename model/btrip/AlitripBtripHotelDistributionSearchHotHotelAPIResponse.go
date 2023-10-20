package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripHotelDistributionSearchHotHotelAPIResponse 商旅酒店api分销-热点酒店 API返回值
// alitrip.btrip.hotel.distribution.search.hot.hotel
//
// 商旅酒店api分销-热点酒店
type AlitripBtripHotelDistributionSearchHotHotelAPIResponse struct {
	model.CommonResponse
	AlitripBtripHotelDistributionSearchHotHotelAPIResponseModel
}

// AlitripBtripHotelDistributionSearchHotHotelAPIResponseModel is 商旅酒店api分销-热点酒店 成功返回结果
type AlitripBtripHotelDistributionSearchHotHotelAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_hotel_distribution_search_hot_hotel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回出参
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
