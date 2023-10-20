package car

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelCrsdriverArrangeAPIResponse CRS接送机商家派司机接口 API返回值
// alitrip.travel.crsdriver.arrange
//
// 提供给CRS接送机商家派司机的API
type AlitripTravelCrsdriverArrangeAPIResponse struct {
	model.CommonResponse
	AlitripTravelCrsdriverArrangeAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTravelCrsdriverArrangeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTravelCrsdriverArrangeAPIResponseModel).Reset()
}

// AlitripTravelCrsdriverArrangeAPIResponseModel is CRS接送机商家派司机接口 成功返回结果
type AlitripTravelCrsdriverArrangeAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_crsdriver_arrange_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果code
	MessageCode int64 `json:"message_code,omitempty" xml:"message_code,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTravelCrsdriverArrangeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.MessageCode = 0
}

var poolAlitripTravelCrsdriverArrangeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTravelCrsdriverArrangeAPIResponse)
	},
}

// GetAlitripTravelCrsdriverArrangeAPIResponse 从 sync.Pool 获取 AlitripTravelCrsdriverArrangeAPIResponse
func GetAlitripTravelCrsdriverArrangeAPIResponse() *AlitripTravelCrsdriverArrangeAPIResponse {
	return poolAlitripTravelCrsdriverArrangeAPIResponse.Get().(*AlitripTravelCrsdriverArrangeAPIResponse)
}

// ReleaseAlitripTravelCrsdriverArrangeAPIResponse 将 AlitripTravelCrsdriverArrangeAPIResponse 保存到 sync.Pool
func ReleaseAlitripTravelCrsdriverArrangeAPIResponse(v *AlitripTravelCrsdriverArrangeAPIResponse) {
	v.Reset()
	poolAlitripTravelCrsdriverArrangeAPIResponse.Put(v)
}
