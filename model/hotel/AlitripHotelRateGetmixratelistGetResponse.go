package hotel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
酒店评论接口 APIResponse
alitrip.hotel.rate.getmixratelist.get

酒店评论接口
*/
type AlitripHotelRateGetmixratelistGetAPIResponse struct {
    model.CommonResponse
    // Response *AlitripHotelRateGetmixratelistGetResponse `json:"alitrip_hotel_rate_getmixratelist_get_response,omitempty"` 
    AlitripHotelRateGetmixratelistGetResponse
}

/* model for simplify = false
type AlitripHotelRateGetmixratelistGetResponse struct {

    // result
    
    Result  *struct {
        AlitripHotelRateGetmixratelistGetResult  *AlitripHotelRateGetmixratelistGetResult `json:"alitrip_hotel_rate_getmixratelist_get_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlitripHotelRateGetmixratelistGetResponse struct {

    // result
    Result   *AlitripHotelRateGetmixratelistGetResult `json:"result,omitempty"`

}
