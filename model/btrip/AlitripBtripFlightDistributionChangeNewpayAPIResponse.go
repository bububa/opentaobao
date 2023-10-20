package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionchangenewpayAPIResponse 商旅机票改签支付V2 API返回值
// alitrip.btrip.flight.distribution.change.newpay
//
// 商旅机票改签支付V2
type AlitripbtripflightdistributionchangenewpayAPIResponse struct {
	model.CommonResponse
	AlitripbtripflightdistributionchangenewpayAPIResponseModel
}

// AlitripbtripflightdistributionchangenewpayAPIResponseModel is 商旅机票改签支付V2 成功返回结果
type AlitripbtripflightdistributionchangenewpayAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_change_newpay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
