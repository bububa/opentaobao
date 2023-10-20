package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscMjsActivityDeleteAPIResponse 删除满就送活动 API返回值
// taobao.promotionmisc.mjs.activity.delete
//
// 删除满就送活动
type TaobaoPromotionmiscMjsActivityDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionmiscMjsActivityDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscMjsActivityDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionmiscMjsActivityDeleteAPIResponseModel).Reset()
}

// TaobaoPromotionmiscMjsActivityDeleteAPIResponseModel is 删除满就送活动 成功返回结果
type TaobaoPromotionmiscMjsActivityDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"promotionmisc_mjs_activity_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功删除活动。
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscMjsActivityDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoPromotionmiscMjsActivityDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionmiscMjsActivityDeleteAPIResponse)
	},
}

// GetTaobaoPromotionmiscMjsActivityDeleteAPIResponse 从 sync.Pool 获取 TaobaoPromotionmiscMjsActivityDeleteAPIResponse
func GetTaobaoPromotionmiscMjsActivityDeleteAPIResponse() *TaobaoPromotionmiscMjsActivityDeleteAPIResponse {
	return poolTaobaoPromotionmiscMjsActivityDeleteAPIResponse.Get().(*TaobaoPromotionmiscMjsActivityDeleteAPIResponse)
}

// ReleaseTaobaoPromotionmiscMjsActivityDeleteAPIResponse 将 TaobaoPromotionmiscMjsActivityDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionmiscMjsActivityDeleteAPIResponse(v *TaobaoPromotionmiscMjsActivityDeleteAPIResponse) {
	v.Reset()
	poolTaobaoPromotionmiscMjsActivityDeleteAPIResponse.Put(v)
}
