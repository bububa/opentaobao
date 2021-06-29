package hotelhstdf

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
匹配标准房型以及卖家房型 APIResponse
alitrip.hotel.hstdf.shotel.matchsroomself

匹配卖家房型以及标准房型，返回匹配结果
*/
type AlitripHotelHstdfShotelMatchsroomselfAPIResponse struct {
    model.CommonResponse
    AlitripHotelHstdfShotelMatchsroomselfResponse
}

type AlitripHotelHstdfShotelMatchsroomselfResponse struct {
    XMLName xml.Name `xml:"alitrip_hotel_hstdf_shotel_matchsroomself_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // error_code
    
    Errorcode   string `json:"errorcode,omitempty" xml:"errorcode,omitempty"`

    
    // error_msg
    
    Errormsg   string `json:"errormsg,omitempty" xml:"errormsg,omitempty"`

    
    // 是否成功
    
    Status   bool `json:"status,omitempty" xml:"status,omitempty"`

    
}
