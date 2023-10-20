package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscCommonItemActivityGetAPIResponse 查询通用单品优惠活动 API返回值
// taobao.promotionmisc.common.item.activity.get
//
// 查询通用单品优惠活动。
type TaobaoPromotionmiscCommonItemActivityGetAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionmiscCommonItemActivityGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscCommonItemActivityGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionmiscCommonItemActivityGetAPIResponseModel).Reset()
}

// TaobaoPromotionmiscCommonItemActivityGetAPIResponseModel is 查询通用单品优惠活动 成功返回结果
type TaobaoPromotionmiscCommonItemActivityGetAPIResponseModel struct {
	XMLName xml.Name `xml:"promotionmisc_common_item_activity_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 优惠活动
	Activity *CommonItemActivity `json:"activity,omitempty" xml:"activity,omitempty"`
	// 是否查询成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscCommonItemActivityGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Activity = nil
	m.IsSuccess = false
}

var poolTaobaoPromotionmiscCommonItemActivityGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionmiscCommonItemActivityGetAPIResponse)
	},
}

// GetTaobaoPromotionmiscCommonItemActivityGetAPIResponse 从 sync.Pool 获取 TaobaoPromotionmiscCommonItemActivityGetAPIResponse
func GetTaobaoPromotionmiscCommonItemActivityGetAPIResponse() *TaobaoPromotionmiscCommonItemActivityGetAPIResponse {
	return poolTaobaoPromotionmiscCommonItemActivityGetAPIResponse.Get().(*TaobaoPromotionmiscCommonItemActivityGetAPIResponse)
}

// ReleaseTaobaoPromotionmiscCommonItemActivityGetAPIResponse 将 TaobaoPromotionmiscCommonItemActivityGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionmiscCommonItemActivityGetAPIResponse(v *TaobaoPromotionmiscCommonItemActivityGetAPIResponse) {
	v.Reset()
	poolTaobaoPromotionmiscCommonItemActivityGetAPIResponse.Put(v)
}
