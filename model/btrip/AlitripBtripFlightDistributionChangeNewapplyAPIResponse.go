package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionChangeNewapplyAPIResponse 商旅机票改签申请V2 API返回值
// alitrip.btrip.flight.distribution.change.newapply
//
// 商旅机票改签申请
type AlitripBtripFlightDistributionChangeNewapplyAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionChangeNewapplyAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionChangeNewapplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripFlightDistributionChangeNewapplyAPIResponseModel).Reset()
}

// AlitripBtripFlightDistributionChangeNewapplyAPIResponseModel is 商旅机票改签申请V2 成功返回结果
type AlitripBtripFlightDistributionChangeNewapplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_change_newapply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionChangeNewapplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripFlightDistributionChangeNewapplyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripFlightDistributionChangeNewapplyAPIResponse)
	},
}

// GetAlitripBtripFlightDistributionChangeNewapplyAPIResponse 从 sync.Pool 获取 AlitripBtripFlightDistributionChangeNewapplyAPIResponse
func GetAlitripBtripFlightDistributionChangeNewapplyAPIResponse() *AlitripBtripFlightDistributionChangeNewapplyAPIResponse {
	return poolAlitripBtripFlightDistributionChangeNewapplyAPIResponse.Get().(*AlitripBtripFlightDistributionChangeNewapplyAPIResponse)
}

// ReleaseAlitripBtripFlightDistributionChangeNewapplyAPIResponse 将 AlitripBtripFlightDistributionChangeNewapplyAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripFlightDistributionChangeNewapplyAPIResponse(v *AlitripBtripFlightDistributionChangeNewapplyAPIResponse) {
	v.Reset()
	poolAlitripBtripFlightDistributionChangeNewapplyAPIResponse.Put(v)
}
