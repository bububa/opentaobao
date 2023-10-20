package hotelhstdf

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelHstdfHotelroomstaticGetAPIResponse 根据类型查询静态字段 API返回值
// alitrip.hotel.hstdf.hotelroomstatic.get
//
// 根据类型查询分页静态字段
type AlitripHotelHstdfHotelroomstaticGetAPIResponse struct {
	model.CommonResponse
	AlitripHotelHstdfHotelroomstaticGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripHotelHstdfHotelroomstaticGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripHotelHstdfHotelroomstaticGetAPIResponseModel).Reset()
}

// AlitripHotelHstdfHotelroomstaticGetAPIResponseModel is 根据类型查询静态字段 成功返回结果
type AlitripHotelHstdfHotelroomstaticGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_hotel_hstdf_hotelroomstatic_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// top返回结果
	Result *TopResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripHotelHstdfHotelroomstaticGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripHotelHstdfHotelroomstaticGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripHotelHstdfHotelroomstaticGetAPIResponse)
	},
}

// GetAlitripHotelHstdfHotelroomstaticGetAPIResponse 从 sync.Pool 获取 AlitripHotelHstdfHotelroomstaticGetAPIResponse
func GetAlitripHotelHstdfHotelroomstaticGetAPIResponse() *AlitripHotelHstdfHotelroomstaticGetAPIResponse {
	return poolAlitripHotelHstdfHotelroomstaticGetAPIResponse.Get().(*AlitripHotelHstdfHotelroomstaticGetAPIResponse)
}

// ReleaseAlitripHotelHstdfHotelroomstaticGetAPIResponse 将 AlitripHotelHstdfHotelroomstaticGetAPIResponse 保存到 sync.Pool
func ReleaseAlitripHotelHstdfHotelroomstaticGetAPIResponse(v *AlitripHotelHstdfHotelroomstaticGetAPIResponse) {
	v.Reset()
	poolAlitripHotelHstdfHotelroomstaticGetAPIResponse.Put(v)
}
