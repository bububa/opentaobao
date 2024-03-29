package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscMjsActivityListGetAPIResponse 查询满就送活动列表 API返回值
// taobao.promotionmisc.mjs.activity.list.get
//
// 查询满就送活动列表。注意，该接口的返回值中，只包含活动的主要信息，如activity_id，name，description，start_time，end_time，type，participate_range。优惠的详细信息，请通过taobao.promotionmisc.mjs.activity.get获取。
type TaobaoPromotionmiscMjsActivityListGetAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionmiscMjsActivityListGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscMjsActivityListGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionmiscMjsActivityListGetAPIResponseModel).Reset()
}

// TaobaoPromotionmiscMjsActivityListGetAPIResponseModel is 查询满就送活动列表 成功返回结果
type TaobaoPromotionmiscMjsActivityListGetAPIResponseModel struct {
	XMLName xml.Name `xml:"promotionmisc_mjs_activity_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 只包含活动的主要信息，如activity_id，aame，description，start_time，end_time，type，participate_range。优惠的其他详细信息，请通过taobao.promotionmisc.mjs.activity.get获取。
	MjsPromotionList []MjsPromotion `json:"mjs_promotion_list,omitempty" xml:"mjs_promotion_list>mjs_promotion,omitempty"`
	// 记录总条数。
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscMjsActivityListGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MjsPromotionList = m.MjsPromotionList[:0]
	m.TotalCount = 0
}

var poolTaobaoPromotionmiscMjsActivityListGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionmiscMjsActivityListGetAPIResponse)
	},
}

// GetTaobaoPromotionmiscMjsActivityListGetAPIResponse 从 sync.Pool 获取 TaobaoPromotionmiscMjsActivityListGetAPIResponse
func GetTaobaoPromotionmiscMjsActivityListGetAPIResponse() *TaobaoPromotionmiscMjsActivityListGetAPIResponse {
	return poolTaobaoPromotionmiscMjsActivityListGetAPIResponse.Get().(*TaobaoPromotionmiscMjsActivityListGetAPIResponse)
}

// ReleaseTaobaoPromotionmiscMjsActivityListGetAPIResponse 将 TaobaoPromotionmiscMjsActivityListGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionmiscMjsActivityListGetAPIResponse(v *TaobaoPromotionmiscMjsActivityListGetAPIResponse) {
	v.Reset()
	poolTaobaoPromotionmiscMjsActivityListGetAPIResponse.Put(v)
}
