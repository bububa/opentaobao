package hotelhstdf

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据类型查询静态字段 APIResponse
alitrip.hotel.hstdf.hotelroomstatic.get

根据类型查询分页静态字段
*/
type AlitripHotelHstdfHotelroomstaticGetAPIResponse struct {
    model.CommonResponse
    AlitripHotelHstdfHotelroomstaticGetResponse
}

type AlitripHotelHstdfHotelroomstaticGetResponse struct {
    XMLName xml.Name `xml:"alitrip_hotel_hstdf_hotelroomstatic_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // top返回结果
    
    Result   *TopResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
