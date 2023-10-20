package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightdistributionordernewcreateAPIResponse 商旅机票分销-创建订单V2 API返回值
// alitrip.btrip.flight.distribution.order.newcreate
//
// 商旅机票分销-创建订单V2
type AlitripbtripflightdistributionordernewcreateAPIResponse struct {
	model.CommonResponse
	AlitripbtripflightdistributionordernewcreateAPIResponseModel
}

// AlitripbtripflightdistributionordernewcreateAPIResponseModel is 商旅机票分销-创建订单V2 成功返回结果
type AlitripbtripflightdistributionordernewcreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_order_newcreate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
