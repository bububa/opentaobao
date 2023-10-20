package hotelhstdf

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlitripHotelHstdfShotelRoomtypeMappingsListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripHotelHstdfShotelRoomtypeMappingsListAPIResponseModel).Reset()
}

// AlitripHotelHstdfShotelRoomtypeMappingsListAPIResponseModel is 根据HID获取所有卖家房型匹配关系 成功返回结果
type AlitripHotelHstdfShotelRoomtypeMappingsListAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_hotel_hstdf_shotel_roomtype_mappings_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// top返回结果
	Result *TopResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripHotelHstdfShotelRoomtypeMappingsListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripHotelHstdfShotelRoomtypeMappingsListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripHotelHstdfShotelRoomtypeMappingsListAPIResponse)
	},
}

// GetAlitripHotelHstdfShotelRoomtypeMappingsListAPIResponse 从 sync.Pool 获取 AlitripHotelHstdfShotelRoomtypeMappingsListAPIResponse
func GetAlitripHotelHstdfShotelRoomtypeMappingsListAPIResponse() *AlitripHotelHstdfShotelRoomtypeMappingsListAPIResponse {
	return poolAlitripHotelHstdfShotelRoomtypeMappingsListAPIResponse.Get().(*AlitripHotelHstdfShotelRoomtypeMappingsListAPIResponse)
}

// ReleaseAlitripHotelHstdfShotelRoomtypeMappingsListAPIResponse 将 AlitripHotelHstdfShotelRoomtypeMappingsListAPIResponse 保存到 sync.Pool
func ReleaseAlitripHotelHstdfShotelRoomtypeMappingsListAPIResponse(v *AlitripHotelHstdfShotelRoomtypeMappingsListAPIResponse) {
	v.Reset()
	poolAlitripHotelHstdfShotelRoomtypeMappingsListAPIResponse.Put(v)
}
