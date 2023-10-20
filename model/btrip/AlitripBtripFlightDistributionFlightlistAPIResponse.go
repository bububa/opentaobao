package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionflightlistAPIResponse 商旅机票航班列表接口 API返回值
// alitrip.btrip.flight.distribution.flightlist
//
// 商旅机票航班列表接口，用于分销询价
type AlitripbtripflightdistributionflightlistAPIResponse struct {
	model.CommonResponse
	AlitripbtripflightdistributionflightlistAPIResponseModel
}

// AlitripbtripflightdistributionflightlistAPIResponseModel is 商旅机票航班列表接口 成功返回结果
type AlitripbtripflightdistributionflightlistAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_flightlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
