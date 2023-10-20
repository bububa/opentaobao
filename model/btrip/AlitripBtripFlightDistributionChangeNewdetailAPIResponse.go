package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionChangeNewdetailAPIResponse 商旅机票改签详情接口 API返回值
// alitrip.btrip.flight.distribution.change.newdetail
//
// 商旅机票改签详情接口
type AlitripBtripFlightDistributionChangeNewdetailAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionChangeNewdetailAPIResponseModel
}

// AlitripBtripFlightDistributionChangeNewdetailAPIResponseModel is 商旅机票改签详情接口 成功返回结果
type AlitripBtripFlightDistributionChangeNewdetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_change_newdetail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
