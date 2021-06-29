package hotelhstdf

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
导出一个卖家房型下的所有标准房型 APIResponse
alitrip.hotel.hstdf.shotel.exportsroomtype

导出一个卖家酒店下的所有标准房型
*/
type AlitripHotelHstdfShotelExportsroomtypeAPIResponse struct {
    model.CommonResponse
    AlitripHotelHstdfShotelExportsroomtypeResponse
}

type AlitripHotelHstdfShotelExportsroomtypeResponse struct {
    XMLName xml.Name `xml:"alitrip_hotel_hstdf_shotel_exportsroomtype_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // top返回结果
    
    Result   *TopResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
