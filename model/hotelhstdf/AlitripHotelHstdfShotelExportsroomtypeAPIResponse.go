package hotelhstdf

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelHstdfShotelExportsroomtypeAPIResponse 导出一个卖家房型下的所有标准房型 API返回值
// alitrip.hotel.hstdf.shotel.exportsroomtype
//
// 导出一个卖家酒店下的所有标准房型
type AlitripHotelHstdfShotelExportsroomtypeAPIResponse struct {
	model.CommonResponse
	AlitripHotelHstdfShotelExportsroomtypeAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripHotelHstdfShotelExportsroomtypeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripHotelHstdfShotelExportsroomtypeAPIResponseModel).Reset()
}

// AlitripHotelHstdfShotelExportsroomtypeAPIResponseModel is 导出一个卖家房型下的所有标准房型 成功返回结果
type AlitripHotelHstdfShotelExportsroomtypeAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_hotel_hstdf_shotel_exportsroomtype_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// top返回结果
	Result *TopResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripHotelHstdfShotelExportsroomtypeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripHotelHstdfShotelExportsroomtypeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripHotelHstdfShotelExportsroomtypeAPIResponse)
	},
}

// GetAlitripHotelHstdfShotelExportsroomtypeAPIResponse 从 sync.Pool 获取 AlitripHotelHstdfShotelExportsroomtypeAPIResponse
func GetAlitripHotelHstdfShotelExportsroomtypeAPIResponse() *AlitripHotelHstdfShotelExportsroomtypeAPIResponse {
	return poolAlitripHotelHstdfShotelExportsroomtypeAPIResponse.Get().(*AlitripHotelHstdfShotelExportsroomtypeAPIResponse)
}

// ReleaseAlitripHotelHstdfShotelExportsroomtypeAPIResponse 将 AlitripHotelHstdfShotelExportsroomtypeAPIResponse 保存到 sync.Pool
func ReleaseAlitripHotelHstdfShotelExportsroomtypeAPIResponse(v *AlitripHotelHstdfShotelExportsroomtypeAPIResponse) {
	v.Reset()
	poolAlitripHotelHstdfShotelExportsroomtypeAPIResponse.Put(v)
}
