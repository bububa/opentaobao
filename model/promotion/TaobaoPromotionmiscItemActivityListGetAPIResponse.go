package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscItemActivityListGetAPIResponse 查询无条件单品优惠活动列表 API返回值
// taobao.promotionmisc.item.activity.list.get
//
// 查询无条件单品优惠活动列表
type TaobaoPromotionmiscItemActivityListGetAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionmiscItemActivityListGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscItemActivityListGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionmiscItemActivityListGetAPIResponseModel).Reset()
}

// TaobaoPromotionmiscItemActivityListGetAPIResponseModel is 查询无条件单品优惠活动列表 成功返回结果
type TaobaoPromotionmiscItemActivityListGetAPIResponseModel struct {
	XMLName xml.Name `xml:"promotionmisc_item_activity_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 只包含活动的主要信息，如activity_id，name，description，start_time，end_time，participate_range。优惠的其他详细信息，请通过taobao.promotionmisc.item.activity.get获取。
	ItemPromotionList []ItemPromotion `json:"item_promotion_list,omitempty" xml:"item_promotion_list>item_promotion,omitempty"`
	// 记录总条数。
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscItemActivityListGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ItemPromotionList = m.ItemPromotionList[:0]
	m.TotalCount = 0
}

var poolTaobaoPromotionmiscItemActivityListGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionmiscItemActivityListGetAPIResponse)
	},
}

// GetTaobaoPromotionmiscItemActivityListGetAPIResponse 从 sync.Pool 获取 TaobaoPromotionmiscItemActivityListGetAPIResponse
func GetTaobaoPromotionmiscItemActivityListGetAPIResponse() *TaobaoPromotionmiscItemActivityListGetAPIResponse {
	return poolTaobaoPromotionmiscItemActivityListGetAPIResponse.Get().(*TaobaoPromotionmiscItemActivityListGetAPIResponse)
}

// ReleaseTaobaoPromotionmiscItemActivityListGetAPIResponse 将 TaobaoPromotionmiscItemActivityListGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionmiscItemActivityListGetAPIResponse(v *TaobaoPromotionmiscItemActivityListGetAPIResponse) {
	v.Reset()
	poolTaobaoPromotionmiscItemActivityListGetAPIResponse.Put(v)
}
