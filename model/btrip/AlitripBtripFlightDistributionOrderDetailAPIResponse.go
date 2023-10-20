package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionOrderDetailAPIResponse 商旅机票分销订单详情接口 API返回值
// alitrip.btrip.flight.distribution.order.detail
//
// 商旅机票分销订单详情接口
type AlitripBtripFlightDistributionOrderDetailAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionOrderDetailAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionOrderDetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripFlightDistributionOrderDetailAPIResponseModel).Reset()
}

// AlitripBtripFlightDistributionOrderDetailAPIResponseModel is 商旅机票分销订单详情接口 成功返回结果
type AlitripBtripFlightDistributionOrderDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_order_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionOrderDetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripFlightDistributionOrderDetailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripFlightDistributionOrderDetailAPIResponse)
	},
}

// GetAlitripBtripFlightDistributionOrderDetailAPIResponse 从 sync.Pool 获取 AlitripBtripFlightDistributionOrderDetailAPIResponse
func GetAlitripBtripFlightDistributionOrderDetailAPIResponse() *AlitripBtripFlightDistributionOrderDetailAPIResponse {
	return poolAlitripBtripFlightDistributionOrderDetailAPIResponse.Get().(*AlitripBtripFlightDistributionOrderDetailAPIResponse)
}

// ReleaseAlitripBtripFlightDistributionOrderDetailAPIResponse 将 AlitripBtripFlightDistributionOrderDetailAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripFlightDistributionOrderDetailAPIResponse(v *AlitripBtripFlightDistributionOrderDetailAPIResponse) {
	v.Reset()
	poolAlitripBtripFlightDistributionOrderDetailAPIResponse.Put(v)
}
