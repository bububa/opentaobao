package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionChangeNewqueryAPIResponse 改签航班询价V2 API返回值
// alitrip.btrip.flight.distribution.change.newquery
//
// 商旅机票改签航班询价
type AlitripBtripFlightDistributionChangeNewqueryAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionChangeNewqueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionChangeNewqueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripFlightDistributionChangeNewqueryAPIResponseModel).Reset()
}

// AlitripBtripFlightDistributionChangeNewqueryAPIResponseModel is 改签航班询价V2 成功返回结果
type AlitripBtripFlightDistributionChangeNewqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_change_newquery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionChangeNewqueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripFlightDistributionChangeNewqueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripFlightDistributionChangeNewqueryAPIResponse)
	},
}

// GetAlitripBtripFlightDistributionChangeNewqueryAPIResponse 从 sync.Pool 获取 AlitripBtripFlightDistributionChangeNewqueryAPIResponse
func GetAlitripBtripFlightDistributionChangeNewqueryAPIResponse() *AlitripBtripFlightDistributionChangeNewqueryAPIResponse {
	return poolAlitripBtripFlightDistributionChangeNewqueryAPIResponse.Get().(*AlitripBtripFlightDistributionChangeNewqueryAPIResponse)
}

// ReleaseAlitripBtripFlightDistributionChangeNewqueryAPIResponse 将 AlitripBtripFlightDistributionChangeNewqueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripFlightDistributionChangeNewqueryAPIResponse(v *AlitripBtripFlightDistributionChangeNewqueryAPIResponse) {
	v.Reset()
	poolAlitripBtripFlightDistributionChangeNewqueryAPIResponse.Put(v)
}
