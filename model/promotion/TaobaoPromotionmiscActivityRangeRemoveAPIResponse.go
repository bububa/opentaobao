package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscActivityRangeRemoveAPIResponse 去除活动参与的商品 API返回值
// taobao.promotionmisc.activity.range.remove
//
// 去除活动参与的商品
type TaobaoPromotionmiscActivityRangeRemoveAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionmiscActivityRangeRemoveAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscActivityRangeRemoveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionmiscActivityRangeRemoveAPIResponseModel).Reset()
}

// TaobaoPromotionmiscActivityRangeRemoveAPIResponseModel is 去除活动参与的商品 成功返回结果
type TaobaoPromotionmiscActivityRangeRemoveAPIResponseModel struct {
	XMLName xml.Name `xml:"promotionmisc_activity_range_remove_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 去除活动参与的商品是否成功。
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscActivityRangeRemoveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoPromotionmiscActivityRangeRemoveAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionmiscActivityRangeRemoveAPIResponse)
	},
}

// GetTaobaoPromotionmiscActivityRangeRemoveAPIResponse 从 sync.Pool 获取 TaobaoPromotionmiscActivityRangeRemoveAPIResponse
func GetTaobaoPromotionmiscActivityRangeRemoveAPIResponse() *TaobaoPromotionmiscActivityRangeRemoveAPIResponse {
	return poolTaobaoPromotionmiscActivityRangeRemoveAPIResponse.Get().(*TaobaoPromotionmiscActivityRangeRemoveAPIResponse)
}

// ReleaseTaobaoPromotionmiscActivityRangeRemoveAPIResponse 将 TaobaoPromotionmiscActivityRangeRemoveAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionmiscActivityRangeRemoveAPIResponse(v *TaobaoPromotionmiscActivityRangeRemoveAPIResponse) {
	v.Reset()
	poolTaobaoPromotionmiscActivityRangeRemoveAPIResponse.Put(v)
}
