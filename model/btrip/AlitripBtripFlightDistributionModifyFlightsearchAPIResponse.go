package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionmodifyflightsearchAPIResponse 改签航班列表 API返回值
// alitrip.btrip.flight.distribution.modify.flightsearch
//
// 商旅分销改签航班列表
type AlitripbtripflightdistributionmodifyflightsearchAPIResponse struct {
	model.CommonResponse
	AlitripbtripflightdistributionmodifyflightsearchAPIResponseModel
}

// AlitripbtripflightdistributionmodifyflightsearchAPIResponseModel is 改签航班列表 成功返回结果
type AlitripbtripflightdistributionmodifyflightsearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_modify_flightsearch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
