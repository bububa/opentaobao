package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionrefundnewdetailAPIResponse 商旅机票退票详情接口V2 API返回值
// alitrip.btrip.flight.distribution.refund.newdetail
//
// 商旅机票退票详情接口V2
type AlitripbtripflightdistributionrefundnewdetailAPIResponse struct {
	model.CommonResponse
	AlitripbtripflightdistributionrefundnewdetailAPIResponseModel
}

// AlitripbtripflightdistributionrefundnewdetailAPIResponseModel is 商旅机票退票详情接口V2 成功返回结果
type AlitripbtripflightdistributionrefundnewdetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_refund_newdetail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
