package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionOrderCreateAPIResponse 商旅机票分销-创建订单 API返回值
// alitrip.btrip.flight.distribution.order.create
//
// 商旅机票分销创建订单接口
type AlitripBtripFlightDistributionOrderCreateAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionOrderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionOrderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripFlightDistributionOrderCreateAPIResponseModel).Reset()
}

// AlitripBtripFlightDistributionOrderCreateAPIResponseModel is 商旅机票分销-创建订单 成功返回结果
type AlitripBtripFlightDistributionOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionOrderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripFlightDistributionOrderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripFlightDistributionOrderCreateAPIResponse)
	},
}

// GetAlitripBtripFlightDistributionOrderCreateAPIResponse 从 sync.Pool 获取 AlitripBtripFlightDistributionOrderCreateAPIResponse
func GetAlitripBtripFlightDistributionOrderCreateAPIResponse() *AlitripBtripFlightDistributionOrderCreateAPIResponse {
	return poolAlitripBtripFlightDistributionOrderCreateAPIResponse.Get().(*AlitripBtripFlightDistributionOrderCreateAPIResponse)
}

// ReleaseAlitripBtripFlightDistributionOrderCreateAPIResponse 将 AlitripBtripFlightDistributionOrderCreateAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripFlightDistributionOrderCreateAPIResponse(v *AlitripBtripFlightDistributionOrderCreateAPIResponse) {
	v.Reset()
	poolAlitripBtripFlightDistributionOrderCreateAPIResponse.Put(v)
}
