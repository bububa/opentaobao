package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPricePromotionItemAddAPIResponse 新增档期商品 API返回值
// alibaba.price.promotion.item.add
//
// 批量新增档期活动商品
type AlibabaPricePromotionItemAddAPIResponse struct {
	model.CommonResponse
	AlibabaPricePromotionItemAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaPricePromotionItemAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaPricePromotionItemAddAPIResponseModel).Reset()
}

// AlibabaPricePromotionItemAddAPIResponseModel is 新增档期商品 成功返回结果
type AlibabaPricePromotionItemAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_price_promotion_item_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaPricePromotionItemAddResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaPricePromotionItemAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaPricePromotionItemAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaPricePromotionItemAddAPIResponse)
	},
}

// GetAlibabaPricePromotionItemAddAPIResponse 从 sync.Pool 获取 AlibabaPricePromotionItemAddAPIResponse
func GetAlibabaPricePromotionItemAddAPIResponse() *AlibabaPricePromotionItemAddAPIResponse {
	return poolAlibabaPricePromotionItemAddAPIResponse.Get().(*AlibabaPricePromotionItemAddAPIResponse)
}

// ReleaseAlibabaPricePromotionItemAddAPIResponse 将 AlibabaPricePromotionItemAddAPIResponse 保存到 sync.Pool
func ReleaseAlibabaPricePromotionItemAddAPIResponse(v *AlibabaPricePromotionItemAddAPIResponse) {
	v.Reset()
	poolAlibabaPricePromotionItemAddAPIResponse.Put(v)
}
