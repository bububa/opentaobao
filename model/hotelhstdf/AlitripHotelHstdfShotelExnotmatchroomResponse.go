package hotelhstdf

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
导出一个hid下所有未匹配rid的接口 API返回值 
alitrip.hotel.hstdf.shotel.exnotmatchroom

导出一个卖家hid下所有未匹配的rid信息，包括rid，名称、英文名称、床型、窗型、面积、对外系统id
*/
type AlitripHotelHstdfShotelExnotmatchroomAPIResponse struct {
    model.CommonResponse
    AlitripHotelHstdfShotelExnotmatchroomResponse
}

// 导出一个hid下所有未匹配rid的接口 成功返回结果
type AlitripHotelHstdfShotelExnotmatchroomResponse struct {
    XMLName xml.Name `xml:"alitrip_hotel_hstdf_shotel_exnotmatchroom_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // top返回结果
    Result   *TopResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
