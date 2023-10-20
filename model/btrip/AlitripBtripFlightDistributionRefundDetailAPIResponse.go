package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionRefundDetailAPIResponse 商旅机票退票详情接口 API返回值
// alitrip.btrip.flight.distribution.refund.detail
//
// 商旅机票分销退票详情
type AlitripBtripFlightDistributionRefundDetailAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionRefundDetailAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionRefundDetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripFlightDistributionRefundDetailAPIResponseModel).Reset()
}

// AlitripBtripFlightDistributionRefundDetailAPIResponseModel is 商旅机票退票详情接口 成功返回结果
type AlitripBtripFlightDistributionRefundDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_refund_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionRefundDetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripFlightDistributionRefundDetailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripFlightDistributionRefundDetailAPIResponse)
	},
}

// GetAlitripBtripFlightDistributionRefundDetailAPIResponse 从 sync.Pool 获取 AlitripBtripFlightDistributionRefundDetailAPIResponse
func GetAlitripBtripFlightDistributionRefundDetailAPIResponse() *AlitripBtripFlightDistributionRefundDetailAPIResponse {
	return poolAlitripBtripFlightDistributionRefundDetailAPIResponse.Get().(*AlitripBtripFlightDistributionRefundDetailAPIResponse)
}

// ReleaseAlitripBtripFlightDistributionRefundDetailAPIResponse 将 AlitripBtripFlightDistributionRefundDetailAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripFlightDistributionRefundDetailAPIResponse(v *AlitripBtripFlightDistributionRefundDetailAPIResponse) {
	v.Reset()
	poolAlitripBtripFlightDistributionRefundDetailAPIResponse.Put(v)
}
