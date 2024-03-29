package hotel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelRateGetmixratelistGetAPIResponse 酒店评论接口 API返回值
// alitrip.hotel.rate.getmixratelist.get
//
// 酒店评论接口
type AlitripHotelRateGetmixratelistGetAPIResponse struct {
	model.CommonResponse
	AlitripHotelRateGetmixratelistGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripHotelRateGetmixratelistGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripHotelRateGetmixratelistGetAPIResponseModel).Reset()
}

// AlitripHotelRateGetmixratelistGetAPIResponseModel is 酒店评论接口 成功返回结果
type AlitripHotelRateGetmixratelistGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_hotel_rate_getmixratelist_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlitripHotelRateGetmixratelistGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripHotelRateGetmixratelistGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripHotelRateGetmixratelistGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripHotelRateGetmixratelistGetAPIResponse)
	},
}

// GetAlitripHotelRateGetmixratelistGetAPIResponse 从 sync.Pool 获取 AlitripHotelRateGetmixratelistGetAPIResponse
func GetAlitripHotelRateGetmixratelistGetAPIResponse() *AlitripHotelRateGetmixratelistGetAPIResponse {
	return poolAlitripHotelRateGetmixratelistGetAPIResponse.Get().(*AlitripHotelRateGetmixratelistGetAPIResponse)
}

// ReleaseAlitripHotelRateGetmixratelistGetAPIResponse 将 AlitripHotelRateGetmixratelistGetAPIResponse 保存到 sync.Pool
func ReleaseAlitripHotelRateGetmixratelistGetAPIResponse(v *AlitripHotelRateGetmixratelistGetAPIResponse) {
	v.Reset()
	poolAlitripHotelRateGetmixratelistGetAPIResponse.Put(v)
}
