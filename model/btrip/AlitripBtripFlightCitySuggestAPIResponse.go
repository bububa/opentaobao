package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightCitySuggestAPIResponse 机票城市搜索 API返回值
// alitrip.btrip.flight.city.suggest
//
// 提供机票城市搜索接口，提高OA用户对接效率
type AlitripBtripFlightCitySuggestAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightCitySuggestAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripFlightCitySuggestAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripFlightCitySuggestAPIResponseModel).Reset()
}

// AlitripBtripFlightCitySuggestAPIResponseModel is 机票城市搜索 成功返回结果
type AlitripBtripFlightCitySuggestAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_city_suggest_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *BtripApplyResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripFlightCitySuggestAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripFlightCitySuggestAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripFlightCitySuggestAPIResponse)
	},
}

// GetAlitripBtripFlightCitySuggestAPIResponse 从 sync.Pool 获取 AlitripBtripFlightCitySuggestAPIResponse
func GetAlitripBtripFlightCitySuggestAPIResponse() *AlitripBtripFlightCitySuggestAPIResponse {
	return poolAlitripBtripFlightCitySuggestAPIResponse.Get().(*AlitripBtripFlightCitySuggestAPIResponse)
}

// ReleaseAlitripBtripFlightCitySuggestAPIResponse 将 AlitripBtripFlightCitySuggestAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripFlightCitySuggestAPIResponse(v *AlitripBtripFlightCitySuggestAPIResponse) {
	v.Reset()
	poolAlitripBtripFlightCitySuggestAPIResponse.Put(v)
}
