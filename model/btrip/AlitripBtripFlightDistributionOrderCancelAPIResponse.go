package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionordercancelAPIResponse 商旅机票分销-取消订单 API返回值
// alitrip.btrip.flight.distribution.order.cancel
//
// 商旅机票分销取消订单
type AlitripbtripflightdistributionordercancelAPIResponse struct {
	model.CommonResponse
	AlitripbtripflightdistributionordercancelAPIResponseModel
}

// AlitripbtripflightdistributionordercancelAPIResponseModel is 商旅机票分销-取消订单 成功返回结果
type AlitripbtripflightdistributionordercancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
