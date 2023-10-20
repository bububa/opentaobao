package btrip

import (
	"encoding/xml"

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

// AlitripBtripFlightDistributionChangePayAPIResponseModel is 商旅机票改签支付 成功返回结果
type AlitripBtripFlightDistributionChangePayAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_change_pay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
