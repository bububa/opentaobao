package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionRefundPrecalAPIResponse 商旅机票分销-退票费预计算 API返回值
// alitrip.btrip.flight.distribution.refund.precal
//
// 商旅机票分销-退票费预计算
type AlitripBtripFlightDistributionRefundPrecalAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionRefundPrecalAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionRefundPrecalAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripFlightDistributionRefundPrecalAPIResponseModel).Reset()
}

// AlitripBtripFlightDistributionRefundPrecalAPIResponseModel is 商旅机票分销-退票费预计算 成功返回结果
type AlitripBtripFlightDistributionRefundPrecalAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_refund_precal_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionRefundPrecalAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripFlightDistributionRefundPrecalAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripFlightDistributionRefundPrecalAPIResponse)
	},
}

// GetAlitripBtripFlightDistributionRefundPrecalAPIResponse 从 sync.Pool 获取 AlitripBtripFlightDistributionRefundPrecalAPIResponse
func GetAlitripBtripFlightDistributionRefundPrecalAPIResponse() *AlitripBtripFlightDistributionRefundPrecalAPIResponse {
	return poolAlitripBtripFlightDistributionRefundPrecalAPIResponse.Get().(*AlitripBtripFlightDistributionRefundPrecalAPIResponse)
}

// ReleaseAlitripBtripFlightDistributionRefundPrecalAPIResponse 将 AlitripBtripFlightDistributionRefundPrecalAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripFlightDistributionRefundPrecalAPIResponse(v *AlitripBtripFlightDistributionRefundPrecalAPIResponse) {
	v.Reset()
	poolAlitripBtripFlightDistributionRefundPrecalAPIResponse.Put(v)
}
