package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionRefundNewapplyAPIResponse 商旅机票分销-退票申请 API返回值
// alitrip.btrip.flight.distribution.refund.newapply
//
// 商旅机票分销-退票申请
type AlitripBtripFlightDistributionRefundNewapplyAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionRefundNewapplyAPIResponseModel
}

// AlitripBtripFlightDistributionRefundNewapplyAPIResponseModel is 商旅机票分销-退票申请 成功返回结果
type AlitripBtripFlightDistributionRefundNewapplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_refund_newapply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
