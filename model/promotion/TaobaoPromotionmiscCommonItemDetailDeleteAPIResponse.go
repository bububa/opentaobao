package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscCommonItemDetailDeleteAPIResponse 删除通用单品优惠详情 API返回值
// taobao.promotionmisc.common.item.detail.delete
//
// 删除通用单品优惠详情。
type TaobaoPromotionmiscCommonItemDetailDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionmiscCommonItemDetailDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscCommonItemDetailDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionmiscCommonItemDetailDeleteAPIResponseModel).Reset()
}

// TaobaoPromotionmiscCommonItemDetailDeleteAPIResponseModel is 删除通用单品优惠详情 成功返回结果
type TaobaoPromotionmiscCommonItemDetailDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"promotionmisc_common_item_detail_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否删除成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscCommonItemDetailDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoPromotionmiscCommonItemDetailDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionmiscCommonItemDetailDeleteAPIResponse)
	},
}

// GetTaobaoPromotionmiscCommonItemDetailDeleteAPIResponse 从 sync.Pool 获取 TaobaoPromotionmiscCommonItemDetailDeleteAPIResponse
func GetTaobaoPromotionmiscCommonItemDetailDeleteAPIResponse() *TaobaoPromotionmiscCommonItemDetailDeleteAPIResponse {
	return poolTaobaoPromotionmiscCommonItemDetailDeleteAPIResponse.Get().(*TaobaoPromotionmiscCommonItemDetailDeleteAPIResponse)
}

// ReleaseTaobaoPromotionmiscCommonItemDetailDeleteAPIResponse 将 TaobaoPromotionmiscCommonItemDetailDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionmiscCommonItemDetailDeleteAPIResponse(v *TaobaoPromotionmiscCommonItemDetailDeleteAPIResponse) {
	v.Reset()
	poolTaobaoPromotionmiscCommonItemDetailDeleteAPIResponse.Put(v)
}
