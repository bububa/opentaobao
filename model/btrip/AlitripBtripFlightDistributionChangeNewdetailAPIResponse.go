package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionChangeNewdetailAPIResponse 商旅机票改签详情接口 API返回值
// alitrip.btrip.flight.distribution.change.newdetail
//
// 商旅机票改签详情接口
type AlitripBtripFlightDistributionChangeNewdetailAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionChangeNewdetailAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionChangeNewdetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripFlightDistributionChangeNewdetailAPIResponseModel).Reset()
}

// AlitripBtripFlightDistributionChangeNewdetailAPIResponseModel is 商旅机票改签详情接口 成功返回结果
type AlitripBtripFlightDistributionChangeNewdetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_change_newdetail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionChangeNewdetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripFlightDistributionChangeNewdetailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripFlightDistributionChangeNewdetailAPIResponse)
	},
}

// GetAlitripBtripFlightDistributionChangeNewdetailAPIResponse 从 sync.Pool 获取 AlitripBtripFlightDistributionChangeNewdetailAPIResponse
func GetAlitripBtripFlightDistributionChangeNewdetailAPIResponse() *AlitripBtripFlightDistributionChangeNewdetailAPIResponse {
	return poolAlitripBtripFlightDistributionChangeNewdetailAPIResponse.Get().(*AlitripBtripFlightDistributionChangeNewdetailAPIResponse)
}

// ReleaseAlitripBtripFlightDistributionChangeNewdetailAPIResponse 将 AlitripBtripFlightDistributionChangeNewdetailAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripFlightDistributionChangeNewdetailAPIResponse(v *AlitripBtripFlightDistributionChangeNewdetailAPIResponse) {
	v.Reset()
	poolAlitripBtripFlightDistributionChangeNewdetailAPIResponse.Put(v)
}
