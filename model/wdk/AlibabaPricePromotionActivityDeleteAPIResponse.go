package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPricePromotionActivityDeleteAPIResponse 删除档期活动 API返回值
// alibaba.price.promotion.activity.delete
//
// 删除盒马帮档期活动
type AlibabaPricePromotionActivityDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaPricePromotionActivityDeleteAPIResponseModel
}

// AlibabaPricePromotionActivityDeleteAPIResponseModel is 删除档期活动 成功返回结果
type AlibabaPricePromotionActivityDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_price_promotion_activity_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaPricePromotionActivityDeleteResult `json:"result,omitempty" xml:"result,omitempty"`
}
