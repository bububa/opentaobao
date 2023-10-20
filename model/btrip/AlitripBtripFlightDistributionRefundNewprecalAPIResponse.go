package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionRefundNewprecalAPIResponse 商旅机票分销-退票费预计算 API返回值
// alitrip.btrip.flight.distribution.refund.newprecal
//
// 商旅机票分销-退票费预计算
type AlitripBtripFlightDistributionRefundNewprecalAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionRefundNewprecalAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionRefundNewprecalAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripFlightDistributionRefundNewprecalAPIResponseModel).Reset()
}

// AlitripBtripFlightDistributionRefundNewprecalAPIResponseModel is 商旅机票分销-退票费预计算 成功返回结果
type AlitripBtripFlightDistributionRefundNewprecalAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_refund_newprecal_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionRefundNewprecalAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripFlightDistributionRefundNewprecalAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripFlightDistributionRefundNewprecalAPIResponse)
	},
}

// GetAlitripBtripFlightDistributionRefundNewprecalAPIResponse 从 sync.Pool 获取 AlitripBtripFlightDistributionRefundNewprecalAPIResponse
func GetAlitripBtripFlightDistributionRefundNewprecalAPIResponse() *AlitripBtripFlightDistributionRefundNewprecalAPIResponse {
	return poolAlitripBtripFlightDistributionRefundNewprecalAPIResponse.Get().(*AlitripBtripFlightDistributionRefundNewprecalAPIResponse)
}

// ReleaseAlitripBtripFlightDistributionRefundNewprecalAPIResponse 将 AlitripBtripFlightDistributionRefundNewprecalAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripFlightDistributionRefundNewprecalAPIResponse(v *AlitripBtripFlightDistributionRefundNewprecalAPIResponse) {
	v.Reset()
	poolAlitripBtripFlightDistributionRefundNewprecalAPIResponse.Put(v)
}
