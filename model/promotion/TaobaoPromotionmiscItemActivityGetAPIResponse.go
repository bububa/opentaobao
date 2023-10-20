package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscItemActivityGetAPIResponse 查询无条件单品优惠活动 API返回值
// taobao.promotionmisc.item.activity.get
//
// 查询无条件单品优惠活动
type TaobaoPromotionmiscItemActivityGetAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionmiscItemActivityGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscItemActivityGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionmiscItemActivityGetAPIResponseModel).Reset()
}

// TaobaoPromotionmiscItemActivityGetAPIResponseModel is 查询无条件单品优惠活动 成功返回结果
type TaobaoPromotionmiscItemActivityGetAPIResponseModel struct {
	XMLName xml.Name `xml:"promotionmisc_item_activity_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 单品优惠活动信息。
	ItemPromotion *ItemPromotion `json:"item_promotion,omitempty" xml:"item_promotion,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscItemActivityGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ItemPromotion = nil
}

var poolTaobaoPromotionmiscItemActivityGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionmiscItemActivityGetAPIResponse)
	},
}

// GetTaobaoPromotionmiscItemActivityGetAPIResponse 从 sync.Pool 获取 TaobaoPromotionmiscItemActivityGetAPIResponse
func GetTaobaoPromotionmiscItemActivityGetAPIResponse() *TaobaoPromotionmiscItemActivityGetAPIResponse {
	return poolTaobaoPromotionmiscItemActivityGetAPIResponse.Get().(*TaobaoPromotionmiscItemActivityGetAPIResponse)
}

// ReleaseTaobaoPromotionmiscItemActivityGetAPIResponse 将 TaobaoPromotionmiscItemActivityGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionmiscItemActivityGetAPIResponse(v *TaobaoPromotionmiscItemActivityGetAPIResponse) {
	v.Reset()
	poolTaobaoPromotionmiscItemActivityGetAPIResponse.Put(v)
}
