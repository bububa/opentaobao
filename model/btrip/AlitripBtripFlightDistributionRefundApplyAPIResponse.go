package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionRefundApplyAPIResponse 商旅机票分销-退票申请 API返回值
// alitrip.btrip.flight.distribution.refund.apply
//
// 商旅机票分销-退票申请
type AlitripBtripFlightDistributionRefundApplyAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionRefundApplyAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionRefundApplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripFlightDistributionRefundApplyAPIResponseModel).Reset()
}

// AlitripBtripFlightDistributionRefundApplyAPIResponseModel is 商旅机票分销-退票申请 成功返回结果
type AlitripBtripFlightDistributionRefundApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_refund_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionRefundApplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripFlightDistributionRefundApplyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripFlightDistributionRefundApplyAPIResponse)
	},
}

// GetAlitripBtripFlightDistributionRefundApplyAPIResponse 从 sync.Pool 获取 AlitripBtripFlightDistributionRefundApplyAPIResponse
func GetAlitripBtripFlightDistributionRefundApplyAPIResponse() *AlitripBtripFlightDistributionRefundApplyAPIResponse {
	return poolAlitripBtripFlightDistributionRefundApplyAPIResponse.Get().(*AlitripBtripFlightDistributionRefundApplyAPIResponse)
}

// ReleaseAlitripBtripFlightDistributionRefundApplyAPIResponse 将 AlitripBtripFlightDistributionRefundApplyAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripFlightDistributionRefundApplyAPIResponse(v *AlitripBtripFlightDistributionRefundApplyAPIResponse) {
	v.Reset()
	poolAlitripBtripFlightDistributionRefundApplyAPIResponse.Put(v)
}
