package btrip

import (
	"encoding/xml"

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

// AlitripBtripFlightDistributionRefundPrecalAPIResponseModel is 商旅机票分销-退票费预计算 成功返回结果
type AlitripBtripFlightDistributionRefundPrecalAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_refund_precal_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
