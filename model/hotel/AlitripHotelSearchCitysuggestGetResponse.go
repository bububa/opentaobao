package hotel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
城市Suggest接口 APIResponse
alitrip.hotel.search.citysuggest.get

城市Suggest接口
*/
type AlitripHotelSearchCitysuggestGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_hotel_search_citysuggest_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlitripHotelSearchCitysuggestGetResult `json:"result,omitempty" xml:"