package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionOrderCancelAPIResponse 商旅机票分销-取消订单 API返回值
// alitrip.btrip.flight.distribution.order.cancel
//
// 商旅机票分销取消订单
type AlitripBtripFlightDistributionOrderCancelAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionOrderCancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionOrderCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripFlightDistributionOrderCancelAPIResponseModel).Reset()
}

// AlitripBtripFlightDistributionOrderCancelAPIResponseModel is 商旅机票分销-取消订单 成功返回结果
type AlitripBtripFlightDistributionOrderCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionOrderCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripFlightDistributionOrderCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripFlightDistributionOrderCancelAPIResponse)
	},
}

// GetAlitripBtripFlightDistributionOrderCancelAPIResponse 从 sync.Pool 获取 AlitripBtripFlightDistributionOrderCancelAPIResponse
func GetAlitripBtripFlightDistributionOrderCancelAPIResponse() *AlitripBtripFlightDistributionOrderCancelAPIResponse {
	return poolAlitripBtripFlightDistributionOrderCancelAPIResponse.Get().(*AlitripBtripFlightDistributionOrderCancelAPIResponse)
}

// ReleaseAlitripBtripFlightDistributionOrderCancelAPIResponse 将 AlitripBtripFlightDistributionOrderCancelAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripFlightDistributionOrderCancelAPIResponse(v *AlitripBtripFlightDistributionOrderCancelAPIResponse) {
	v.Reset()
	poolAlitripBtripFlightDistributionOrderCancelAPIResponse.Put(v)
}
