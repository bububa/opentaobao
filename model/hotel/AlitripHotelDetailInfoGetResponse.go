package hotel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
详情页动态信息接口 APIResponse
alitrip.hotel.detail.info.get

酒店详情页动态信息TOP方法
*/
type AlitripHotelDetailInfoGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_hotel_detail_info_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlitripHotelDetailInfoGetResult `json:"result,omitempty" xml:"