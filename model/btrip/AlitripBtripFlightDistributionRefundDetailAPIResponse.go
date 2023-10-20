package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionrefunddetailAPIResponse 商旅机票退票详情接口 API返回值
// alitrip.btrip.flight.distribution.refund.detail
//
// 商旅机票分销退票详情
type AlitripbtripflightdistributionrefunddetailAPIResponse struct {
	model.CommonResponse
	AlitripbtripflightdistributionrefunddetailAPIResponseModel
}

// AlitripbtripflightdistributionrefunddetailAPIResponseModel is 商旅机票退票详情接口 成功返回结果
type AlitripbtripflightdistributionrefunddetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_refund_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
