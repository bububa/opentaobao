package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionRefundNewapplyAPIResponse 商旅机票分销-退票申请 API返回值
// alitrip.btrip.flight.distribution.refund.newapply
//
// 商旅机票分销-退票申请
type AlitripBtripFlightDistributionRefundNewapplyAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionRefundNewapplyAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionRefundNewapplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripFlightDistributionRefundNewapplyAPIResponseModel).Reset()
}

// AlitripBtripFlightDistributionRefundNewapplyAPIResponseModel is 商旅机票分销-退票申请 成功返回结果
type AlitripBtripFlightDistributionRefundNewapplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_refund_newapply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionRefundNewapplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripFlightDistributionRefundNewapplyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripFlightDistributionRefundNewapplyAPIResponse)
	},
}

// GetAlitripBtripFlightDistributionRefundNewapplyAPIResponse 从 sync.Pool 获取 AlitripBtripFlightDistributionRefundNewapplyAPIResponse
func GetAlitripBtripFlightDistributionRefundNewapplyAPIResponse() *AlitripBtripFlightDistributionRefundNewapplyAPIResponse {
	return poolAlitripBtripFlightDistributionRefundNewapplyAPIResponse.Get().(*AlitripBtripFlightDistributionRefundNewapplyAPIResponse)
}

// ReleaseAlitripBtripFlightDistributionRefundNewapplyAPIResponse 将 AlitripBtripFlightDistributionRefundNewapplyAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripFlightDistributionRefundNewapplyAPIResponse(v *AlitripBtripFlightDistributionRefundNewapplyAPIResponse) {
	v.Reset()
	poolAlitripBtripFlightDistributionRefundNewapplyAPIResponse.Put(v)
}
