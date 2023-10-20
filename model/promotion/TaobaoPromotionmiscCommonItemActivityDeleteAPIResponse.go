package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscCommonItemActivityDeleteAPIResponse 删除通用单品优惠活动 API返回值
// taobao.promotionmisc.common.item.activity.delete
//
// 删除通用单品优惠活动。
type TaobaoPromotionmiscCommonItemActivityDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionmiscCommonItemActivityDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscCommonItemActivityDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionmiscCommonItemActivityDeleteAPIResponseModel).Reset()
}

// TaobaoPromotionmiscCommonItemActivityDeleteAPIResponseModel is 删除通用单品优惠活动 成功返回结果
type TaobaoPromotionmiscCommonItemActivityDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"promotionmisc_common_item_activity_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否删除成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscCommonItemActivityDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoPromotionmiscCommonItemActivityDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionmiscCommonItemActivityDeleteAPIResponse)
	},
}

// GetTaobaoPromotionmiscCommonItemActivityDeleteAPIResponse 从 sync.Pool 获取 TaobaoPromotionmiscCommonItemActivityDeleteAPIResponse
func GetTaobaoPromotionmiscCommonItemActivityDeleteAPIResponse() *TaobaoPromotionmiscCommonItemActivityDeleteAPIResponse {
	return poolTaobaoPromotionmiscCommonItemActivityDeleteAPIResponse.Get().(*TaobaoPromotionmiscCommonItemActivityDeleteAPIResponse)
}

// ReleaseTaobaoPromotionmiscCommonItemActivityDeleteAPIResponse 将 TaobaoPromotionmiscCommonItemActivityDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionmiscCommonItemActivityDeleteAPIResponse(v *TaobaoPromotionmiscCommonItemActivityDeleteAPIResponse) {
	v.Reset()
	poolTaobaoPromotionmiscCommonItemActivityDeleteAPIResponse.Put(v)
}
