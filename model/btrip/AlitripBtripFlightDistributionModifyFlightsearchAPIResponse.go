package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionModifyFlightsearchAPIResponse 改签航班列表 API返回值
// alitrip.btrip.flight.distribution.modify.flightsearch
//
// 商旅分销改签航班列表
type AlitripBtripFlightDistributionModifyFlightsearchAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionModifyFlightsearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionModifyFlightsearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripFlightDistributionModifyFlightsearchAPIResponseModel).Reset()
}

// AlitripBtripFlightDistributionModifyFlightsearchAPIResponseModel is 改签航班列表 成功返回结果
type AlitripBtripFlightDistributionModifyFlightsearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_modify_flightsearch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionModifyFlightsearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripFlightDistributionModifyFlightsearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripFlightDistributionModifyFlightsearchAPIResponse)
	},
}

// GetAlitripBtripFlightDistributionModifyFlightsearchAPIResponse 从 sync.Pool 获取 AlitripBtripFlightDistributionModifyFlightsearchAPIResponse
func GetAlitripBtripFlightDistributionModifyFlightsearchAPIResponse() *AlitripBtripFlightDistributionModifyFlightsearchAPIResponse {
	return poolAlitripBtripFlightDistributionModifyFlightsearchAPIResponse.Get().(*AlitripBtripFlightDistributionModifyFlightsearchAPIResponse)
}

// ReleaseAlitripBtripFlightDistributionModifyFlightsearchAPIResponse 将 AlitripBtripFlightDistributionModifyFlightsearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripFlightDistributionModifyFlightsearchAPIResponse(v *AlitripBtripFlightDistributionModifyFlightsearchAPIResponse) {
	v.Reset()
	poolAlitripBtripFlightDistributionModifyFlightsearchAPIResponse.Put(v)
}
