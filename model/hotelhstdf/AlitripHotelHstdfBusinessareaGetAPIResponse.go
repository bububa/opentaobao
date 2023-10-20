package hotelhstdf

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelHstdfBusinessareaGetAPIResponse 根据城市查询商圈 API返回值
// alitrip.hotel.hstdf.businessarea.get
//
// 根据cityId分页查询商圈信息
type AlitripHotelHstdfBusinessareaGetAPIResponse struct {
	model.CommonResponse
	AlitripHotelHstdfBusinessareaGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripHotelHstdfBusinessareaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripHotelHstdfBusinessareaGetAPIResponseModel).Reset()
}

// AlitripHotelHstdfBusinessareaGetAPIResponseModel is 根据城市查询商圈 成功返回结果
type AlitripHotelHstdfBusinessareaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_hotel_hstdf_businessarea_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// top返回结果
	Result *TopStdResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripHotelHstdfBusinessareaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripHotelHstdfBusinessareaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripHotelHstdfBusinessareaGetAPIResponse)
	},
}

// GetAlitripHotelHstdfBusinessareaGetAPIResponse 从 sync.Pool 获取 AlitripHotelHstdfBusinessareaGetAPIResponse
func GetAlitripHotelHstdfBusinessareaGetAPIResponse() *AlitripHotelHstdfBusinessareaGetAPIResponse {
	return poolAlitripHotelHstdfBusinessareaGetAPIResponse.Get().(*AlitripHotelHstdfBusinessareaGetAPIResponse)
}

// ReleaseAlitripHotelHstdfBusinessareaGetAPIResponse 将 AlitripHotelHstdfBusinessareaGetAPIResponse 保存到 sync.Pool
func ReleaseAlitripHotelHstdfBusinessareaGetAPIResponse(v *AlitripHotelHstdfBusinessareaGetAPIResponse) {
	v.Reset()
	poolAlitripHotelHstdfBusinessareaGetAPIResponse.Put(v)
}
