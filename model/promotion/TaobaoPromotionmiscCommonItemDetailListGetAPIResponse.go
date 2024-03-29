package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscCommonItemDetailListGetAPIResponse 查询通用单品优惠详情列表 API返回值
// taobao.promotionmisc.common.item.detail.list.get
//
// 查询通用单品优惠详情列表。
type TaobaoPromotionmiscCommonItemDetailListGetAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionmiscCommonItemDetailListGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscCommonItemDetailListGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionmiscCommonItemDetailListGetAPIResponseModel).Reset()
}

// TaobaoPromotionmiscCommonItemDetailListGetAPIResponseModel is 查询通用单品优惠详情列表 成功返回结果
type TaobaoPromotionmiscCommonItemDetailListGetAPIResponseModel struct {
	XMLName xml.Name `xml:"promotionmisc_common_item_detail_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 活动详情列表
	DetailList []CommonItemDetail `json:"detail_list,omitempty" xml:"detail_list>common_item_detail,omitempty"`
	// 数据总数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 是否查询成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscCommonItemDetailListGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DetailList = m.DetailList[:0]
	m.TotalCount = 0
	m.IsSuccess = false
}

var poolTaobaoPromotionmiscCommonItemDetailListGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionmiscCommonItemDetailListGetAPIResponse)
	},
}

// GetTaobaoPromotionmiscCommonItemDetailListGetAPIResponse 从 sync.Pool 获取 TaobaoPromotionmiscCommonItemDetailListGetAPIResponse
func GetTaobaoPromotionmiscCommonItemDetailListGetAPIResponse() *TaobaoPromotionmiscCommonItemDetailListGetAPIResponse {
	return poolTaobaoPromotionmiscCommonItemDetailListGetAPIResponse.Get().(*TaobaoPromotionmiscCommonItemDetailListGetAPIResponse)
}

// ReleaseTaobaoPromotionmiscCommonItemDetailListGetAPIResponse 将 TaobaoPromotionmiscCommonItemDetailListGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionmiscCommonItemDetailListGetAPIResponse(v *TaobaoPromotionmiscCommonItemDetailListGetAPIResponse) {
	v.Reset()
	poolTaobaoPromotionmiscCommonItemDetailListGetAPIResponse.Put(v)
}
