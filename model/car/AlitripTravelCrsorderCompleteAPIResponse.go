package car

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelCrsorderCompleteAPIResponse CRS接送机商家服务完成接口 API返回值
// alitrip.travel.crsorder.complete
//
// 提供给CRS接送机商家的服务完成回调接口
type AlitripTravelCrsorderCompleteAPIResponse struct {
	model.CommonResponse
	AlitripTravelCrsorderCompleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTravelCrsorderCompleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTravelCrsorderCompleteAPIResponseModel).Reset()
}

// AlitripTravelCrsorderCompleteAPIResponseModel is CRS接送机商家服务完成接口 成功返回结果
type AlitripTravelCrsorderCompleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_crsorder_complete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果code
	MessageCode int64 `json:"message_code,omitempty" xml:"message_code,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTravelCrsorderCompleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.MessageCode = 0
}

var poolAlitripTravelCrsorderCompleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTravelCrsorderCompleteAPIResponse)
	},
}

// GetAlitripTravelCrsorderCompleteAPIResponse 从 sync.Pool 获取 AlitripTravelCrsorderCompleteAPIResponse
func GetAlitripTravelCrsorderCompleteAPIResponse() *AlitripTravelCrsorderCompleteAPIResponse {
	return poolAlitripTravelCrsorderCompleteAPIResponse.Get().(*AlitripTravelCrsorderCompleteAPIResponse)
}

// ReleaseAlitripTravelCrsorderCompleteAPIResponse 将 AlitripTravelCrsorderCompleteAPIResponse 保存到 sync.Pool
func ReleaseAlitripTravelCrsorderCompleteAPIResponse(v *AlitripTravelCrsorderCompleteAPIResponse) {
	v.Reset()
	poolAlitripTravelCrsorderCompleteAPIResponse.Put(v)
}
