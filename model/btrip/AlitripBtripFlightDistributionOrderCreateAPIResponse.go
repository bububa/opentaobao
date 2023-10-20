package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionordercreateAPIResponse 商旅机票分销-创建订单 API返回值
// alitrip.btrip.flight.distribution.order.create
//
// 商旅机票分销创建订单接口
type AlitripbtripflightdistributionordercreateAPIResponse struct {
	model.CommonResponse
	AlitripbtripflightdistributionordercreateAPIResponseModel
}

// AlitripbtripflightdistributionordercreateAPIResponseModel is 商旅机票分销-创建订单 成功返回结果
type AlitripbtripflightdistributionordercreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
