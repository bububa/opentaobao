package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyorderqueryinfoAPIResponse 订单详情改版 API返回值
// alitrip.merchant.galaxy.order.query.info
//
// 订单页详情查询
type AlitripmerchantgalaxyorderqueryinfoAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyorderqueryinfoAPIResponseModel
}

// AlitripmerchantgalaxyorderqueryinfoAPIResponseModel is 订单详情改版 成功返回结果
type AlitripmerchantgalaxyorderqueryinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_order_query_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlitripmerchantgalaxyorderqueryinfoResponse `json:"result,omitempty" xml:"result,omitempty"`
}
