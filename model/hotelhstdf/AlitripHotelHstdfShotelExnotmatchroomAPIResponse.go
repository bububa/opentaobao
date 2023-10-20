package hotelhstdf

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelHstdfShotelExnotmatchroomAPIResponse 导出一个hid下所有未匹配rid的接口 API返回值
// alitrip.hotel.hstdf.shotel.exnotmatchroom
//
// 导出一个卖家hid下所有未匹配的rid信息，包括rid，名称、英文名称、床型、窗型、面积、对外系统id
type AlitripHotelHstdfShotelExnotmatchroomAPIResponse struct {
	model.CommonResponse
	AlitripHotelHstdfShotelExnotmatchroomAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripHotelHstdfShotelExnotmatchroomAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripHotelHstdfShotelExnotmatchroomAPIResponseModel).Reset()
}

// AlitripHotelHstdfShotelExnotmatchroomAPIResponseModel is 导出一个hid下所有未匹配rid的接口 成功返回结果
type AlitripHotelHstdfShotelExnotmatchroomAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_hotel_hstdf_shotel_exnotmatchroom_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// top返回结果
	Result *TopResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripHotelHstdfShotelExnotmatchroomAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripHotelHstdfShotelExnotmatchroomAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripHotelHstdfShotelExnotmatchroomAPIResponse)
	},
}

// GetAlitripHotelHstdfShotelExnotmatchroomAPIResponse 从 sync.Pool 获取 AlitripHotelHstdfShotelExnotmatchroomAPIResponse
func GetAlitripHotelHstdfShotelExnotmatchroomAPIResponse() *AlitripHotelHstdfShotelExnotmatchroomAPIResponse {
	return poolAlitripHotelHstdfShotelExnotmatchroomAPIResponse.Get().(*AlitripHotelHstdfShotelExnotmatchroomAPIResponse)
}

// ReleaseAlitripHotelHstdfShotelExnotmatchroomAPIResponse 将 AlitripHotelHstdfShotelExnotmatchroomAPIResponse 保存到 sync.Pool
func ReleaseAlitripHotelHstdfShotelExnotmatchroomAPIResponse(v *AlitripHotelHstdfShotelExnotmatchroomAPIResponse) {
	v.Reset()
	poolAlitripHotelHstdfShotelExnotmatchroomAPIResponse.Put(v)
}
