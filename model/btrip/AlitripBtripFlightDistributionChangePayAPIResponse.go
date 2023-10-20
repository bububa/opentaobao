package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionChangePayAPIResponse 商旅机票改签支付 API返回值
// alitrip.btrip.flight.distribution.change.pay
//
// 改签订单支付
type AlitripBtripFlightDistributionChangePayAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionChangePayAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionChangePayAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripFlightDistributionChangePayAPIResponseModel).Reset()
}

// AlitripBtripFlightDistributionChangePayAPIResponseModel is 商旅机票改签支付 成功返回结果
type AlitripBtripFlightDistributionChangePayAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_change_pay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionChangePayAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripFlightDistributionChangePayAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripFlightDistributionChangePayAPIResponse)
	},
}

// GetAlitripBtripFlightDistributionChangePayAPIResponse 从 sync.Pool 获取 AlitripBtripFlightDistributionChangePayAPIResponse
func GetAlitripBtripFlightDistributionChangePayAPIResponse() *AlitripBtripFlightDistributionChangePayAPIResponse {
	return poolAlitripBtripFlightDistributionChangePayAPIResponse.Get().(*AlitripBtripFlightDistributionChangePayAPIResponse)
}

// ReleaseAlitripBtripFlightDistributionChangePayAPIResponse 将 AlitripBtripFlightDistributionChangePayAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripFlightDistributionChangePayAPIResponse(v *AlitripBtripFlightDistributionChangePayAPIResponse) {
	v.Reset()
	poolAlitripBtripFlightDistributionChangePayAPIResponse.Put(v)
}
