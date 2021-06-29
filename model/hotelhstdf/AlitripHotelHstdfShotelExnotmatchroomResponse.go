package hotelhstdf

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
导出一个hid下所有未匹配rid的接口 APIResponse
alitrip.hotel.hstdf.shotel.exnotmatchroom

导出一个卖家hid下所有未匹配的rid信息，包括rid，名称、英文名称、床型、窗型、面积、对外系统id
*/
type AlitripHotelHstdfShotelExnotmatchroomAPIResponse struct {
    model.CommonResponse
    AlitripHotelHstdfShotelExnotmatchroomResponse
}

type AlitripHotelHstdfShotelExnotmatchroomResponse struct {
    XMLName xml.Name `xml:"alitrip_hotel_hstdf_shotel_exnotmatchroom_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // top返回结果
    
    Result   *TopResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
