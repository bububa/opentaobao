package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscActivityRangeAddAPIResponse 增加活动参与的商品 API返回值
// taobao.promotionmisc.activity.range.add
//
// 增加活动参与的商品，部分商品参与的活动，最大支持指定150个商品。
type TaobaoPromotionmiscActivityRangeAddAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionmiscActivityRangeAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscActivityRangeAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionmiscActivityRangeAddAPIResponseModel).Reset()
}

// TaobaoPromotionmiscActivityRangeAddAPIResponseModel is 增加活动参与的商品 成功返回结果
type TaobaoPromotionmiscActivityRangeAddAPIResponseModel struct {
	XMLName xml.Name `xml:"promotionmisc_activity_range_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 增加商品范围是否成功。
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscActivityRangeAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoPromotionmiscActivityRangeAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionmiscActivityRangeAddAPIResponse)
	},
}

// GetTaobaoPromotionmiscActivityRangeAddAPIResponse 从 sync.Pool 获取 TaobaoPromotionmiscActivityRangeAddAPIResponse
func GetTaobaoPromotionmiscActivityRangeAddAPIResponse() *TaobaoPromotionmiscActivityRangeAddAPIResponse {
	return poolTaobaoPromotionmiscActivityRangeAddAPIResponse.Get().(*TaobaoPromotionmiscActivityRangeAddAPIResponse)
}

// ReleaseTaobaoPromotionmiscActivityRangeAddAPIResponse 将 TaobaoPromotionmiscActivityRangeAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionmiscActivityRangeAddAPIResponse(v *TaobaoPromotionmiscActivityRangeAddAPIResponse) {
	v.Reset()
	poolTaobaoPromotionmiscActivityRangeAddAPIResponse.Put(v)
}
