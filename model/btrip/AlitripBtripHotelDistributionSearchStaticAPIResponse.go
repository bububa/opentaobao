package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripHotelDistributionSearchStaticAPIResponse 商旅酒店api分销-酒店静态信息接口 API返回值
// alitrip.btrip.hotel.distribution.search.static
//
// 商旅酒店api分销-酒店静态信息接口
type AlitripBtripHotelDistributionSearchStaticAPIResponse struct {
	model.CommonResponse
	AlitripBtripHotelDistributionSearchStaticAPIResponseModel
}

// AlitripBtripHotelDistributionSearchStaticAPIResponseModel is 商旅酒店api分销-酒店静态信息接口 成功返回结果
type AlitripBtripHotelDistributionSearchStaticAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_hotel_distribution_search_static_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回报文
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
