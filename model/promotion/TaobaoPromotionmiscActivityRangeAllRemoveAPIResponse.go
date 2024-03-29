package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscActivityRangeAllRemoveAPIResponse 清空活动参与的商品 API返回值
// taobao.promotionmisc.activity.range.all.remove
//
// 清空活动参与的商品
type TaobaoPromotionmiscActivityRangeAllRemoveAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionmiscActivityRangeAllRemoveAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscActivityRangeAllRemoveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionmiscActivityRangeAllRemoveAPIResponseModel).Reset()
}

// TaobaoPromotionmiscActivityRangeAllRemoveAPIResponseModel is 清空活动参与的商品 成功返回结果
type TaobaoPromotionmiscActivityRangeAllRemoveAPIResponseModel struct {
	XMLName xml.Name `xml:"promotionmisc_activity_range_all_remove_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 清空活动参与商品是否成功。
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscActivityRangeAllRemoveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoPromotionmiscActivityRangeAllRemoveAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionmiscActivityRangeAllRemoveAPIResponse)
	},
}

// GetTaobaoPromotionmiscActivityRangeAllRemoveAPIResponse 从 sync.Pool 获取 TaobaoPromotionmiscActivityRangeAllRemoveAPIResponse
func GetTaobaoPromotionmiscActivityRangeAllRemoveAPIResponse() *TaobaoPromotionmiscActivityRangeAllRemoveAPIResponse {
	return poolTaobaoPromotionmiscActivityRangeAllRemoveAPIResponse.Get().(*TaobaoPromotionmiscActivityRangeAllRemoveAPIResponse)
}

// ReleaseTaobaoPromotionmiscActivityRangeAllRemoveAPIResponse 将 TaobaoPromotionmiscActivityRangeAllRemoveAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionmiscActivityRangeAllRemoveAPIResponse(v *TaobaoPromotionmiscActivityRangeAllRemoveAPIResponse) {
	v.Reset()
	poolTaobaoPromotionmiscActivityRangeAllRemoveAPIResponse.Put(v)
}
