package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionChangeNewcancelAPIResponse 商旅机票改签取消 API返回值
// alitrip.btrip.flight.distribution.change.newcancel
//
// 商旅机票改签取消
type AlitripBtripFlightDistributionChangeNewcancelAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionChangeNewcancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionChangeNewcancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripFlightDistributionChangeNewcancelAPIResponseModel).Reset()
}

// AlitripBtripFlightDistributionChangeNewcancelAPIResponseModel is 商旅机票改签取消 成功返回结果
type AlitripBtripFlightDistributionChangeNewcancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_change_newcancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 改签取消输出参数
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionChangeNewcancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripFlightDistributionChangeNewcancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripFlightDistributionChangeNewcancelAPIResponse)
	},
}

// GetAlitripBtripFlightDistributionChangeNewcancelAPIResponse 从 sync.Pool 获取 AlitripBtripFlightDistributionChangeNewcancelAPIResponse
func GetAlitripBtripFlightDistributionChangeNewcancelAPIResponse() *AlitripBtripFlightDistributionChangeNewcancelAPIResponse {
	return poolAlitripBtripFlightDistributionChangeNewcancelAPIResponse.Get().(*AlitripBtripFlightDistributionChangeNewcancelAPIResponse)
}

// ReleaseAlitripBtripFlightDistributionChangeNewcancelAPIResponse 将 AlitripBtripFlightDistributionChangeNewcancelAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripFlightDistributionChangeNewcancelAPIResponse(v *AlitripBtripFlightDistributionChangeNewcancelAPIResponse) {
	v.Reset()
	poolAlitripBtripFlightDistributionChangeNewcancelAPIResponse.Put(v)
}
