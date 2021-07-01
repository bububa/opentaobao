package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPricePromotionItemAddAPIResponse
新增档期商品 API返回值
alibaba.price.promotion.item.add

批量新增档期活动商品 */
type AlibabaPricePromotionItemAddAPIResponse struct {
	model.CommonResponse
	AlibabaPricePromotionItemAddAPIResponseModel
}

// AlibabaPricePromotionItemAddAPIResponseModel is 新增档期商品 成功返回结果
type AlibabaPricePromotionItemAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_price_promotion_item_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaPricePromotionItemAddResult `json:"result,omitempty" xml:"result,omitempty"`
}
