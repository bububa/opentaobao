package hotel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店搜索List接口 APIResponse
alitrip.hotel.search.list.get

酒店搜索List接口
*/
type AlitripHotelSearchListGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_hotel_search_list_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlitripHotelSearchListGetResult `json:"result,omitempty" xml:"