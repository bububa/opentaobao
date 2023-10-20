package hotelhstdf

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelHstdfShotelRoomtypeMappingsListAPIResponse 根据HID获取所有卖家房型匹配关系 API返回值
// alitrip.hotel.hstdf.shotel.roomtype.mappings.list
//
// 根据HID获取所有卖家房型匹配关系
type AlitripHotelHstdfShotelRoomtypeMappingsListAPIResponse struct {
	model.CommonResponse
	AlitripHotelHstdfShotelRoomtypeMappingsListAPIResponseModel
}

// AlitripHotelHstdfShotelRoomtypeMappingsListAPIResponseModel is 根据HID获取所有卖家房型匹配关系 成功返回结果
type AlitripHotelHstdfShotelRoomtypeMappingsListAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_hotel_hstdf_shotel_roomtype_mappings_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// top返回结果
	Result *TopResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
