package hotel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
详情页静态信息 APIResponse
alitrip.hotel.detail.staticinfo.get

酒店静态信息TOP接口
*/
type AlitripHotelDetailStaticinfoGetAPIResponse struct {
    model.CommonResponse
    // Response *AlitripHotelDetailStaticinfoGetResponse `json:"alitrip_hotel_detail_staticinfo_get_response,omitempty"` 
    AlitripHotelDetailStaticinfoGetResponse
}

/* model for simplify = false
type AlitripHotelDetailStaticinfoGetResponse struct {

    // result
    
    Result  *struct {
        AlitripHotelDetailStaticinfoGetResult  *AlitripHotelDetailStaticinfoGetResult `json:"alitrip_hotel_detail_staticinfo_get_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlitripHotelDetailStaticinfoGetResponse struct {

    // result
    Result   *AlitripHotelDetailStaticinfoGetResult `json:"result,omitempty"`

}
