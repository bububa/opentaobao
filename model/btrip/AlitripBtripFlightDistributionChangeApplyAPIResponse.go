package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionChangeApplyAPIResponse 商旅机票改签申请 API返回值
// alitrip.btrip.flight.distribution.change.apply
//
// 商旅机票改签申请
type AlitripBtripFlightDistributionChangeApplyAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionChangeApplyAPIResponseModel
}

// AlitripBtripFlightDistributionChangeApplyAPIResponseModel is 商旅机票改签申请 成功返回结果
type AlitripBtripFlightDistributionChangeApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_change_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
