package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionchangecancelAPIResponse 商旅机票改签取消 API返回值
// alitrip.btrip.flight.distribution.change.cancel
//
// 商旅机票改签取消
type AlitripbtripflightdistributionchangecancelAPIResponse struct {
	model.CommonResponse
	AlitripbtripflightdistributionchangecancelAPIResponseModel
}

// AlitripbtripflightdistributionchangecancelAPIResponseModel is 商旅机票改签取消 成功返回结果
type AlitripbtripflightdistributionchangecancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_change_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 改签取消输出参数
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
