package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionOrderPayAPIResponse 商旅机票分销-订单支付 API返回值
// alitrip.btrip.flight.distribution.order.pay
//
// 商旅机票分销订单支付
type AlitripBtripFlightDistributionOrderPayAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionOrderPayAPIResponseModel
}

// AlitripBtripFlightDistributionOrderPayAPIResponseModel is 商旅机票分销-订单支付 成功返回结果
type AlitripBtripFlightDistributionOrderPayAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_order_pay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
