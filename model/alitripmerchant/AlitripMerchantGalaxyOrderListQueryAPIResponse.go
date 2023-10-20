package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyorderlistqueryAPIResponse 星河-订单列表查询 API返回值
// alitrip.merchant.galaxy.order.list.query
//
// 为C端用户提供酒店预订订单列表查询服务，包括订单支付状态、订单日期
type AlitripmerchantgalaxyorderlistqueryAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyorderlistqueryAPIResponseModel
}

// AlitripmerchantgalaxyorderlistqueryAPIResponseModel is 星河-订单列表查询 成功返回结果
type AlitripmerchantgalaxyorderlistqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_order_list_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripmerchantgalaxyorderlistqueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}
