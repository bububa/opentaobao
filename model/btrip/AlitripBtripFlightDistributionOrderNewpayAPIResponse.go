package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionordernewpayAPIResponse 商旅机票分销-订单支付V2 API返回值
// alitrip.btrip.flight.distribution.order.newpay
//
// 商旅机票分销-订单支付V2
type AlitripbtripflightdistributionordernewpayAPIResponse struct {
	model.CommonResponse
	AlitripbtripflightdistributionordernewpayAPIResponseModel
}

// AlitripbtripflightdistributionordernewpayAPIResponseModel is 商旅机票分销-订单支付V2 成功返回结果
type AlitripbtripflightdistributionordernewpayAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_order_newpay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
