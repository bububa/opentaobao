package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionchangenewapplyAPIResponse 商旅机票改签申请V2 API返回值
// alitrip.btrip.flight.distribution.change.newapply
//
// 商旅机票改签申请
type AlitripbtripflightdistributionchangenewapplyAPIResponse struct {
	model.CommonResponse
	AlitripbtripflightdistributionchangenewapplyAPIResponseModel
}

// AlitripbtripflightdistributionchangenewapplyAPIResponseModel is 商旅机票改签申请V2 成功返回结果
type AlitripbtripflightdistributionchangenewapplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_change_newapply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
