package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyorderbookAPIResponse 星河-订单预订接口 API返回值
// alitrip.merchant.galaxy.order.book
//
// 为星河酒店解决方案的C端用户提供酒店预订能力
type AlitripmerchantgalaxyorderbookAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyorderbookAPIResponseModel
}

// AlitripmerchantgalaxyorderbookAPIResponseModel is 星河-订单预订接口 成功返回结果
type AlitripmerchantgalaxyorderbookAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_order_book_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripmerchantgalaxyorderbookResponse `json:"result,omitempty" xml:"result,omitempty"`
}
