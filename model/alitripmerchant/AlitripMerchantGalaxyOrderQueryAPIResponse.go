package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyOrderQueryAPIResponse 星河-单个订单详细信息查询 API返回值
// alitrip.merchant.galaxy.order.query
//
// 为用户提供酒店订单的详细信息查询能力
type AlitripMerchantGalaxyOrderQueryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyOrderQueryAPIResponseModel
}

// AlitripMerchantGalaxyOrderQueryAPIResponseModel is 星河-单个订单详细信息查询 成功返回结果
type AlitripMerchantGalaxyOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyOrderQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}
