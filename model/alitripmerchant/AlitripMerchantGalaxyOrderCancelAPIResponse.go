package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyordercancelAPIResponse 星河-取消预订 API返回值
// alitrip.merchant.galaxy.order.cancel
//
// 雅高酒店用户使用该接口，取消酒店预订
type AlitripmerchantgalaxyordercancelAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyordercancelAPIResponseModel
}

// AlitripmerchantgalaxyordercancelAPIResponseModel is 星河-取消预订 成功返回结果
type AlitripmerchantgalaxyordercancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripmerchantgalaxyordercancelResponse `json:"result,omitempty" xml:"result,omitempty"`
}
