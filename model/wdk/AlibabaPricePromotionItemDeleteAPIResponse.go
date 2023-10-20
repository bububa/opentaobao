package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPricePromotionItemDeleteAPIResponse 批量删除档期 API返回值
// alibaba.price.promotion.item.delete
//
// 盒马帮批量删除档期商品
type AlibabaPricePromotionItemDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaPricePromotionItemDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaPricePromotionItemDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaPricePromotionItemDeleteAPIResponseModel).Reset()
}

// AlibabaPricePromotionItemDeleteAPIResponseModel is 批量删除档期 成功返回结果
type AlibabaPricePromotionItemDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_price_promotion_item_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaPricePromotionItemDeleteResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaPricePromotionItemDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaPricePromotionItemDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaPricePromotionItemDeleteAPIResponse)
	},
}

// GetAlibabaPricePromotionItemDeleteAPIResponse 从 sync.Pool 获取 AlibabaPricePromotionItemDeleteAPIResponse
func GetAlibabaPricePromotionItemDeleteAPIResponse() *AlibabaPricePromotionItemDeleteAPIResponse {
	return poolAlibabaPricePromotionItemDeleteAPIResponse.Get().(*AlibabaPricePromotionItemDeleteAPIResponse)
}

// ReleaseAlibabaPricePromotionItemDeleteAPIResponse 将 AlibabaPricePromotionItemDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaPricePromotionItemDeleteAPIResponse(v *AlibabaPricePromotionItemDeleteAPIResponse) {
	v.Reset()
	poolAlibabaPricePromotionItemDeleteAPIResponse.Put(v)
}
