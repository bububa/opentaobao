package hotel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
详情页静态信息 APIResponse
alitrip.hotel.detail.staticinfo.get

酒店静态信息TOP接口
*/
type AlitripHotelDetailStaticinfoGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_hotel_detail_staticinfo_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlitripHotelDetailStaticinfoGetResult `json:"result,omitempty" xml:"