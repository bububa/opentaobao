package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyOrderQueryOrderCountAPIResponse 查询各种状态订单的总数 API返回值
// alitrip.merchant.galaxy.order.query.order.count
//
// 调用查询接口整合各个订单类型总数
type AlitripMerchantGalaxyOrderQueryOrderCountAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyOrderQueryOrderCountAPIResponseModel
}

// AlitripMerchantGalaxyOrderQueryOrderCountAPIResponseModel is 查询各种状态订单的总数 成功返回结果
type AlitripMerchantGalaxyOrderQueryOrderCountAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_order_query_order_count_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyOrderQueryOrderCountResponse `json:"result,omitempty" xml:"result,omitempty"`
}
