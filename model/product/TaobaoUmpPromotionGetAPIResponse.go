package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpPromotionGetAPIResponse 商品优惠详情查询 API返回值
// taobao.ump.promotion.get
//
// 商品优惠详情查询，可查询商品设置的详细优惠。包括限时折扣，满就送等官方优惠以及第三方优惠。
type TaobaoUmpPromotionGetAPIResponse struct {
	model.CommonResponse
	TaobaoUmpPromotionGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUmpPromotionGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUmpPromotionGetAPIResponseModel).Reset()
}

// TaobaoUmpPromotionGetAPIResponseModel is 商品优惠详情查询 成功返回结果
type TaobaoUmpPromotionGetAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_promotion_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 优惠详细信息
	Promotions *PromotionDisplayTop `json:"promotions,omitempty" xml:"promotions,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUmpPromotionGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Promotions = nil
}

var poolTaobaoUmpPromotionGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUmpPromotionGetAPIResponse)
	},
}

// GetTaobaoUmpPromotionGetAPIResponse 从 sync.Pool 获取 TaobaoUmpPromotionGetAPIResponse
func GetTaobaoUmpPromotionGetAPIResponse() *TaobaoUmpPromotionGetAPIResponse {
	return poolTaobaoUmpPromotionGetAPIResponse.Get().(*TaobaoUmpPromotionGetAPIResponse)
}

// ReleaseTaobaoUmpPromotionGetAPIResponse 将 TaobaoUmpPromotionGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUmpPromotionGetAPIResponse(v *TaobaoUmpPromotionGetAPIResponse) {
	v.Reset()
	poolTaobaoUmpPromotionGetAPIResponse.Put(v)
}
