package hotelhstdf

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据HID获取所有卖家房型匹配关系 APIResponse
alitrip.hotel.hstdf.shotel.roomtype.mappings.list

根据HID获取所有卖家房型匹配关系
*/
type AlitripHotelHstdfShotelRoomtypeMappingsListAPIResponse struct {
    model.CommonResponse
    AlitripHotelHstdfShotelRoomtypeMappingsListResponse
}

type AlitripHotelHstdfShotelRoomtypeMappingsListResponse struct {
    XMLName xml.Name `xml:"alitrip_hotel_hstdf_shotel_roomtype_mappings_list_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // top返回结果
    
    Result   *TopResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
