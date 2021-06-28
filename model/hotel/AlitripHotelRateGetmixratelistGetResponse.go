package hotel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店评论接口 APIResponse
alitrip.hotel.rate.getmixratelist.get

酒店评论接口
*/
type AlitripHotelRateGetmixratelistGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_hotel_rate_getmixratelist_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlitripHotelRateGetmixratelistGetResult `json:"result,omitempty" xml:"