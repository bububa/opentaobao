package hotel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
城市Suggest接口 APIResponse
alitrip.hotel.search.citysuggest.get

城市Suggest接口
*/
type AlitripHotelSearchCitysuggestGetAPIResponse struct {
    model.CommonResponse
    Response *AlitripHotelSearchCitysuggestGetResponse `json:"alitrip_hotel_search_citysuggest_get_response,omitempty"`
}

type AlitripHotelSearchCitysuggestGetResponse struct {

    // result
    Result   *AlitripHotelSearchCitysuggestGetResult `json:"result,omitempty"`

}
