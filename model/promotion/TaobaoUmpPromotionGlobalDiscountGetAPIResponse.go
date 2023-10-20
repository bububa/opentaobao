package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpPromotionGlobalDiscountGetAPIResponse 获取卖家最低折扣 API返回值
// taobao.ump.promotion.global.discount.get
//
// 提供卖家最低折扣查询功能
type TaobaoUmpPromotionGlobalDiscountGetAPIResponse struct {
	model.CommonResponse
	TaobaoUmpPromotionGlobalDiscountGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUmpPromotionGlobalDiscountGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUmpPromotionGlobalDiscountGetAPIResponseModel).Reset()
}

// TaobaoUmpPromotionGlobalDiscountGetAPIResponseModel is 获取卖家最低折扣 成功返回结果
type TaobaoUmpPromotionGlobalDiscountGetAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_promotion_global_discount_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *TaobaoUmpPromotionGlobalDiscountGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUmpPromotionGlobalDiscountGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUmpPromotionGlobalDiscountGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUmpPromotionGlobalDiscountGetAPIResponse)
	},
}

// GetTaobaoUmpPromotionGlobalDiscountGetAPIResponse 从 sync.Pool 获取 TaobaoUmpPromotionGlobalDiscountGetAPIResponse
func GetTaobaoUmpPromotionGlobalDiscountGetAPIResponse() *TaobaoUmpPromotionGlobalDiscountGetAPIResponse {
	return poolTaobaoUmpPromotionGlobalDiscountGetAPIResponse.Get().(*TaobaoUmpPromotionGlobalDiscountGetAPIResponse)
}

// ReleaseTaobaoUmpPromotionGlobalDiscountGetAPIResponse 将 TaobaoUmpPromotionGlobalDiscountGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUmpPromotionGlobalDiscountGetAPIResponse(v *TaobaoUmpPromotionGlobalDiscountGetAPIResponse) {
	v.Reset()
	poolTaobaoUmpPromotionGlobalDiscountGetAPIResponse.Put(v)
}
