package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionChangeApplyAPIResponse 商旅机票改签申请 API返回值
// alitrip.btrip.flight.distribution.change.apply
//
// 商旅机票改签申请
type AlitripBtripFlightDistributionChangeApplyAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionChangeApplyAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionChangeApplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripFlightDistributionChangeApplyAPIResponseModel).Reset()
}

// AlitripBtripFlightDistributionChangeApplyAPIResponseModel is 商旅机票改签申请 成功返回结果
type AlitripBtripFlightDistributionChangeApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_change_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionChangeApplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripFlightDistributionChangeApplyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripFlightDistributionChangeApplyAPIResponse)
	},
}

// GetAlitripBtripFlightDistributionChangeApplyAPIResponse 从 sync.Pool 获取 AlitripBtripFlightDistributionChangeApplyAPIResponse
func GetAlitripBtripFlightDistributionChangeApplyAPIResponse() *AlitripBtripFlightDistributionChangeApplyAPIResponse {
	return poolAlitripBtripFlightDistributionChangeApplyAPIResponse.Get().(*AlitripBtripFlightDistributionChangeApplyAPIResponse)
}

// ReleaseAlitripBtripFlightDistributionChangeApplyAPIResponse 将 AlitripBtripFlightDistributionChangeApplyAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripFlightDistributionChangeApplyAPIResponse(v *AlitripBtripFlightDistributionChangeApplyAPIResponse) {
	v.Reset()
	poolAlitripBtripFlightDistributionChangeApplyAPIResponse.Put(v)
}
