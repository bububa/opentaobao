package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionChangeDetailAPIResponse 商旅机票改签详情接口 API返回值
// alitrip.btrip.flight.distribution.change.detail
//
// 商旅机票改签详情接口
type AlitripBtripFlightDistributionChangeDetailAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionChangeDetailAPIResponseModel
}

// AlitripBtripFlightDistributionChangeDetailAPIResponseModel is 商旅机票改签详情接口 成功返回结果
type AlitripBtripFlightDistributionChangeDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_change_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
