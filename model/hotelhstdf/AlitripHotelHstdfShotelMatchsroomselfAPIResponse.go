package hotelhstdf

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelHstdfShotelMatchsroomselfAPIResponse 匹配标准房型以及卖家房型 API返回值
// alitrip.hotel.hstdf.shotel.matchsroomself
//
// 匹配卖家房型以及标准房型，返回匹配结果
type AlitripHotelHstdfShotelMatchsroomselfAPIResponse struct {
	model.CommonResponse
	AlitripHotelHstdfShotelMatchsroomselfAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripHotelHstdfShotelMatchsroomselfAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripHotelHstdfShotelMatchsroomselfAPIResponseModel).Reset()
}

// AlitripHotelHstdfShotelMatchsroomselfAPIResponseModel is 匹配标准房型以及卖家房型 成功返回结果
type AlitripHotelHstdfShotelMatchsroomselfAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_hotel_hstdf_shotel_matchsroomself_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// error_code
	Errorcode string `json:"errorcode,omitempty" xml:"errorcode,omitempty"`
	// error_msg
	Errormsg string `json:"errormsg,omitempty" xml:"errormsg,omitempty"`
	// 是否成功
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
}

// Reset 清空结构体
func (m *AlitripHotelHstdfShotelMatchsroomselfAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errorcode = ""
	m.Errormsg = ""
	m.Status = false
}

var poolAlitripHotelHstdfShotelMatchsroomselfAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripHotelHstdfShotelMatchsroomselfAPIResponse)
	},
}

// GetAlitripHotelHstdfShotelMatchsroomselfAPIResponse 从 sync.Pool 获取 AlitripHotelHstdfShotelMatchsroomselfAPIResponse
func GetAlitripHotelHstdfShotelMatchsroomselfAPIResponse() *AlitripHotelHstdfShotelMatchsroomselfAPIResponse {
	return poolAlitripHotelHstdfShotelMatchsroomselfAPIResponse.Get().(*AlitripHotelHstdfShotelMatchsroomselfAPIResponse)
}

// ReleaseAlitripHotelHstdfShotelMatchsroomselfAPIResponse 将 AlitripHotelHstdfShotelMatchsroomselfAPIResponse 保存到 sync.Pool
func ReleaseAlitripHotelHstdfShotelMatchsroomselfAPIResponse(v *AlitripHotelHstdfShotelMatchsroomselfAPIResponse) {
	v.Reset()
	poolAlitripHotelHstdfShotelMatchsroomselfAPIResponse.Put(v)
}
