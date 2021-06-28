package hotel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
详情页动态信息接口 APIResponse
alitrip.hotel.detail.info.get

酒店详情页动态信息TOP方法
*/
type AlitripHotelDetailInfoGetAPIResponse struct {
    model.CommonResponse
    // Response *AlitripHotelDetailInfoGetResponse `json:"alitrip_hotel_detail_info_get_response,omitempty"` 
    AlitripHotelDetailInfoGetResponse
}

/* model for simplify = false
type AlitripHotelDetailInfoGetResponse struct {

    // result
    
    Result  *struct {
        AlitripHotelDetailInfoGetResult  *AlitripHotelDetailInfoGetResult `json:"alitrip_hotel_detail_info_get_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlitripHotelDetailInfoGetResponse struct {

    // result
    Result   *AlitripHotelDetailInfoGetResult `json:"result,omitempty"`

}
