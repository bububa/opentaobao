package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionRefundNewdetailAPIResponse 商旅机票退票详情接口V2 API返回值
// alitrip.btrip.flight.distribution.refund.newdetail
//
// 商旅机票退票详情接口V2
type AlitripBtripFlightDistributionRefundNewdetailAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionRefundNewdetailAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionRefundNewdetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripFlightDistributionRefundNewdetailAPIResponseModel).Reset()
}

// AlitripBtripFlightDistributionRefundNewdetailAPIResponseModel is 商旅机票退票详情接口V2 成功返回结果
type AlitripBtripFlightDistributionRefundNewdetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_refund_newdetail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionRefundNewdetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripFlightDistributionRefundNewdetailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripFlightDistributionRefundNewdetailAPIResponse)
	},
}

// GetAlitripBtripFlightDistributionRefundNewdetailAPIResponse 从 sync.Pool 获取 AlitripBtripFlightDistributionRefundNewdetailAPIResponse
func GetAlitripBtripFlightDistributionRefundNewdetailAPIResponse() *AlitripBtripFlightDistributionRefundNewdetailAPIResponse {
	return poolAlitripBtripFlightDistributionRefundNewdetailAPIResponse.Get().(*AlitripBtripFlightDistributionRefundNewdetailAPIResponse)
}

// ReleaseAlitripBtripFlightDistributionRefundNewdetailAPIResponse 将 AlitripBtripFlightDistributionRefundNewdetailAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripFlightDistributionRefundNewdetailAPIResponse(v *AlitripBtripFlightDistributionRefundNewdetailAPIResponse) {
	v.Reset()
	poolAlitripBtripFlightDistributionRefundNewdetailAPIResponse.Put(v)
}
