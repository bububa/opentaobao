package hotelhstdf

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelHstdfPoilocationGetAPIResponse 根据平台城市id分页查询poi location API返回值
// alitrip.hotel.hstdf.poilocation.get
//
// 根据平台城市id分页查询poi location
type AlitripHotelHstdfPoilocationGetAPIResponse struct {
	model.CommonResponse
	AlitripHotelHstdfPoilocationGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripHotelHstdfPoilocationGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripHotelHstdfPoilocationGetAPIResponseModel).Reset()
}

// AlitripHotelHstdfPoilocationGetAPIResponseModel is 根据平台城市id分页查询poi location 成功返回结果
type AlitripHotelHstdfPoilocationGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_hotel_hstdf_poilocation_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// top返回结果
	Result *TopStdResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripHotelHstdfPoilocationGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripHotelHstdfPoilocationGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripHotelHstdfPoilocationGetAPIResponse)
	},
}

// GetAlitripHotelHstdfPoilocationGetAPIResponse 从 sync.Pool 获取 AlitripHotelHstdfPoilocationGetAPIResponse
func GetAlitripHotelHstdfPoilocationGetAPIResponse() *AlitripHotelHstdfPoilocationGetAPIResponse {
	return poolAlitripHotelHstdfPoilocationGetAPIResponse.Get().(*AlitripHotelHstdfPoilocationGetAPIResponse)
}

// ReleaseAlitripHotelHstdfPoilocationGetAPIResponse 将 AlitripHotelHstdfPoilocationGetAPIResponse 保存到 sync.Pool
func ReleaseAlitripHotelHstdfPoilocationGetAPIResponse(v *AlitripHotelHstdfPoilocationGetAPIResponse) {
	v.Reset()
	poolAlitripHotelHstdfPoilocationGetAPIResponse.Put(v)
}
