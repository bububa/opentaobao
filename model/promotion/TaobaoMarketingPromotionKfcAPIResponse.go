package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMarketingPromotionKfcAPIResponse 定向优惠活动名称与描述违禁词检查 API返回值
// taobao.marketing.promotion.kfc
//
// 活动名称与描述违禁词检查
type TaobaoMarketingPromotionKfcAPIResponse struct {
	model.CommonResponse
	TaobaoMarketingPromotionKfcAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMarketingPromotionKfcAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMarketingPromotionKfcAPIResponseModel).Reset()
}

// TaobaoMarketingPromotionKfcAPIResponseModel is 定向优惠活动名称与描述违禁词检查 成功返回结果
type TaobaoMarketingPromotionKfcAPIResponseModel struct {
	XMLName xml.Name `xml:"marketing_promotion_kfc_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMarketingPromotionKfcAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoMarketingPromotionKfcAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMarketingPromotionKfcAPIResponse)
	},
}

// GetTaobaoMarketingPromotionKfcAPIResponse 从 sync.Pool 获取 TaobaoMarketingPromotionKfcAPIResponse
func GetTaobaoMarketingPromotionKfcAPIResponse() *TaobaoMarketingPromotionKfcAPIResponse {
	return poolTaobaoMarketingPromotionKfcAPIResponse.Get().(*TaobaoMarketingPromotionKfcAPIResponse)
}

// ReleaseTaobaoMarketingPromotionKfcAPIResponse 将 TaobaoMarketingPromotionKfcAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMarketingPromotionKfcAPIResponse(v *TaobaoMarketingPromotionKfcAPIResponse) {
	v.Reset()
	poolTaobaoMarketingPromotionKfcAPIResponse.Put(v)
}
