package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyofferqueryAPIResponse 星河-offer查询 API返回值
// alitrip.merchant.galaxy.offer.query
//
// 根据offer的ID查询offer信息
type AlitripmerchantgalaxyofferqueryAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyofferqueryAPIResponseModel
}

// AlitripmerchantgalaxyofferqueryAPIResponseModel is 星河-offer查询 成功返回结果
type AlitripmerchantgalaxyofferqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_offer_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripmerchantgalaxyofferqueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}
