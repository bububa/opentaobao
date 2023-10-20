package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscItemActivityDeleteAPIResponse 删除无条件单品优惠活动 API返回值
// taobao.promotionmisc.item.activity.delete
//
// 删除无条件单品优惠活动
type TaobaoPromotionmiscItemActivityDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionmiscItemActivityDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscItemActivityDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionmiscItemActivityDeleteAPIResponseModel).Reset()
}

// TaobaoPromotionmiscItemActivityDeleteAPIResponseModel is 删除无条件单品优惠活动 成功返回结果
type TaobaoPromotionmiscItemActivityDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"promotionmisc_item_activity_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功删除活动。
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscItemActivityDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoPromotionmiscItemActivityDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionmiscItemActivityDeleteAPIResponse)
	},
}

// GetTaobaoPromotionmiscItemActivityDeleteAPIResponse 从 sync.Pool 获取 TaobaoPromotionmiscItemActivityDeleteAPIResponse
func GetTaobaoPromotionmiscItemActivityDeleteAPIResponse() *TaobaoPromotionmiscItemActivityDeleteAPIResponse {
	return poolTaobaoPromotionmiscItemActivityDeleteAPIResponse.Get().(*TaobaoPromotionmiscItemActivityDeleteAPIResponse)
}

// ReleaseTaobaoPromotionmiscItemActivityDeleteAPIResponse 将 TaobaoPromotionmiscItemActivityDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionmiscItemActivityDeleteAPIResponse(v *TaobaoPromotionmiscItemActivityDeleteAPIResponse) {
	v.Reset()
	poolTaobaoPromotionmiscItemActivityDeleteAPIResponse.Put(v)
}
