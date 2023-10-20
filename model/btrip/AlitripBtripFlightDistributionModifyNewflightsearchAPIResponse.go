package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionModifyNewflightsearchAPIResponse 改签航班列表V2 API返回值
// alitrip.btrip.flight.distribution.modify.newflightsearch
//
// 改签航班列表V2
type AlitripBtripFlightDistributionModifyNewflightsearchAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionModifyNewflightsearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionModifyNewflightsearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripFlightDistributionModifyNewflightsearchAPIResponseModel).Reset()
}

// AlitripBtripFlightDistributionModifyNewflightsearchAPIResponseModel is 改签航班列表V2 成功返回结果
type AlitripBtripFlightDistributionModifyNewflightsearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_modify_newflightsearch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应参数
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionModifyNewflightsearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripFlightDistributionModifyNewflightsearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripFlightDistributionModifyNewflightsearchAPIResponse)
	},
}

// GetAlitripBtripFlightDistributionModifyNewflightsearchAPIResponse 从 sync.Pool 获取 AlitripBtripFlightDistributionModifyNewflightsearchAPIResponse
func GetAlitripBtripFlightDistributionModifyNewflightsearchAPIResponse() *AlitripBtripFlightDistributionModifyNewflightsearchAPIResponse {
	return poolAlitripBtripFlightDistributionModifyNewflightsearchAPIResponse.Get().(*AlitripBtripFlightDistributionModifyNewflightsearchAPIResponse)
}

// ReleaseAlitripBtripFlightDistributionModifyNewflightsearchAPIResponse 将 AlitripBtripFlightDistributionModifyNewflightsearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripFlightDistributionModifyNewflightsearchAPIResponse(v *AlitripBtripFlightDistributionModifyNewflightsearchAPIResponse) {
	v.Reset()
	poolAlitripBtripFlightDistributionModifyNewflightsearchAPIResponse.Put(v)
}
