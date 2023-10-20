package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionOrderNewpayAPIResponse 商旅机票分销-订单支付V2 API返回值
// alitrip.btrip.flight.distribution.order.newpay
//
// 商旅机票分销-订单支付V2
type AlitripBtripFlightDistributionOrderNewpayAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionOrderNewpayAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionOrderNewpayAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripFlightDistributionOrderNewpayAPIResponseModel).Reset()
}

// AlitripBtripFlightDistributionOrderNewpayAPIResponseModel is 商旅机票分销-订单支付V2 成功返回结果
type AlitripBtripFlightDistributionOrderNewpayAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_order_newpay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionOrderNewpayAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripFlightDistributionOrderNewpayAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripFlightDistributionOrderNewpayAPIResponse)
	},
}

// GetAlitripBtripFlightDistributionOrderNewpayAPIResponse 从 sync.Pool 获取 AlitripBtripFlightDistributionOrderNewpayAPIResponse
func GetAlitripBtripFlightDistributionOrderNewpayAPIResponse() *AlitripBtripFlightDistributionOrderNewpayAPIResponse {
	return poolAlitripBtripFlightDistributionOrderNewpayAPIResponse.Get().(*AlitripBtripFlightDistributionOrderNewpayAPIResponse)
}

// ReleaseAlitripBtripFlightDistributionOrderNewpayAPIResponse 将 AlitripBtripFlightDistributionOrderNewpayAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripFlightDistributionOrderNewpayAPIResponse(v *AlitripBtripFlightDistributionOrderNewpayAPIResponse) {
	v.Reset()
	poolAlitripBtripFlightDistributionOrderNewpayAPIResponse.Put(v)
}
