package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscActivityRangeListGetAPIResponse 查询活动参与的商品 API返回值
// taobao.promotionmisc.activity.range.list.get
//
// 查询活动参与的商品
type TaobaoPromotionmiscActivityRangeListGetAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionmiscActivityRangeListGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscActivityRangeListGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionmiscActivityRangeListGetAPIResponseModel).Reset()
}

// TaobaoPromotionmiscActivityRangeListGetAPIResponseModel is 查询活动参与的商品 成功返回结果
type TaobaoPromotionmiscActivityRangeListGetAPIResponseModel struct {
	XMLName xml.Name `xml:"promotionmisc_activity_range_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 活动参与的商品列表
	PromotionRangeList []PromotionRange `json:"promotion_range_list,omitempty" xml:"promotion_range_list>promotion_range,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscActivityRangeListGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.PromotionRangeList = m.PromotionRangeList[:0]
}

var poolTaobaoPromotionmiscActivityRangeListGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionmiscActivityRangeListGetAPIResponse)
	},
}

// GetTaobaoPromotionmiscActivityRangeListGetAPIResponse 从 sync.Pool 获取 TaobaoPromotionmiscActivityRangeListGetAPIResponse
func GetTaobaoPromotionmiscActivityRangeListGetAPIResponse() *TaobaoPromotionmiscActivityRangeListGetAPIResponse {
	return poolTaobaoPromotionmiscActivityRangeListGetAPIResponse.Get().(*TaobaoPromotionmiscActivityRangeListGetAPIResponse)
}

// ReleaseTaobaoPromotionmiscActivityRangeListGetAPIResponse 将 TaobaoPromotionmiscActivityRangeListGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionmiscActivityRangeListGetAPIResponse(v *TaobaoPromotionmiscActivityRangeListGetAPIResponse) {
	v.Reset()
	poolTaobaoPromotionmiscActivityRangeListGetAPIResponse.Put(v)
}
