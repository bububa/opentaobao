package hotel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelSearchCitysuggestGetAPIResponse 城市Suggest接口 API返回值
// alitrip.hotel.search.citysuggest.get
//
// 城市Suggest接口
type AlitripHotelSearchCitysuggestGetAPIResponse struct {
	model.CommonResponse
	AlitripHotelSearchCitysuggestGetAPIResponseModel
}

// AlitripHotelSearchCitysuggestGetAPIResponseModel is 城市Suggest接口 成功返回结果
type AlitripHotelSearchCitysuggestGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_hotel_search_citysuggest_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlitripHotelSearchCitysuggestGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
