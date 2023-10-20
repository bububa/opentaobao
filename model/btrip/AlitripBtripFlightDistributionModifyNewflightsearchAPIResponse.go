package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionmodifynewflightsearchAPIResponse 改签航班列表V2 API返回值
// alitrip.btrip.flight.distribution.modify.newflightsearch
//
// 改签航班列表V2
type AlitripbtripflightdistributionmodifynewflightsearchAPIResponse struct {
	model.CommonResponse
	AlitripbtripflightdistributionmodifynewflightsearchAPIResponseModel
}

// AlitripbtripflightdistributionmodifynewflightsearchAPIResponseModel is 改签航班列表V2 成功返回结果
type AlitripbtripflightdistributionmodifynewflightsearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_modify_newflightsearch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应参数
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
