package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionChangeNewpayAPIResponse 商旅机票改签支付V2 API返回值
// alitrip.btrip.flight.distribution.change.newpay
//
// 商旅机票改签支付V2
type AlitripBtripFlightDistributionChangeNewpayAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionChangeNewpayAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionChangeNewpayAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripFlightDistributionChangeNewpayAPIResponseModel).Reset()
}

// AlitripBtripFlightDistributionChangeNewpayAPIResponseModel is 商旅机票改签支付V2 成功返回结果
type AlitripBtripFlightDistributionChangeNewpayAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_change_newpay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionChangeNewpayAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripFlightDistributionChangeNewpayAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripFlightDistributionChangeNewpayAPIResponse)
	},
}

// GetAlitripBtripFlightDistributionChangeNewpayAPIResponse 从 sync.Pool 获取 AlitripBtripFlightDistributionChangeNewpayAPIResponse
func GetAlitripBtripFlightDistributionChangeNewpayAPIResponse() *AlitripBtripFlightDistributionChangeNewpayAPIResponse {
	return poolAlitripBtripFlightDistributionChangeNewpayAPIResponse.Get().(*AlitripBtripFlightDistributionChangeNewpayAPIResponse)
}

// ReleaseAlitripBtripFlightDistributionChangeNewpayAPIResponse 将 AlitripBtripFlightDistributionChangeNewpayAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripFlightDistributionChangeNewpayAPIResponse(v *AlitripBtripFlightDistributionChangeNewpayAPIResponse) {
	v.Reset()
	poolAlitripBtripFlightDistributionChangeNewpayAPIResponse.Put(v)
}
