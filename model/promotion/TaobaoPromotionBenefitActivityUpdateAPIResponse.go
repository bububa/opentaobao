package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionBenefitActivityUpdateAPIResponse 修改关联的活动权益 API返回值
// taobao.promotion.benefit.activity.update
//
// 修改卖家活动中关联的对应的权益。
type TaobaoPromotionBenefitActivityUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionBenefitActivityUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionBenefitActivityUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionBenefitActivityUpdateAPIResponseModel).Reset()
}

// TaobaoPromotionBenefitActivityUpdateAPIResponseModel is 修改关联的活动权益 成功返回结果
type TaobaoPromotionBenefitActivityUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"promotion_benefit_activity_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 更新是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionBenefitActivityUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoPromotionBenefitActivityUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionBenefitActivityUpdateAPIResponse)
	},
}

// GetTaobaoPromotionBenefitActivityUpdateAPIResponse 从 sync.Pool 获取 TaobaoPromotionBenefitActivityUpdateAPIResponse
func GetTaobaoPromotionBenefitActivityUpdateAPIResponse() *TaobaoPromotionBenefitActivityUpdateAPIResponse {
	return poolTaobaoPromotionBenefitActivityUpdateAPIResponse.Get().(*TaobaoPromotionBenefitActivityUpdateAPIResponse)
}

// ReleaseTaobaoPromotionBenefitActivityUpdateAPIResponse 将 TaobaoPromotionBenefitActivityUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionBenefitActivityUpdateAPIResponse(v *TaobaoPromotionBenefitActivityUpdateAPIResponse) {
	v.Reset()
	poolTaobaoPromotionBenefitActivityUpdateAPIResponse.Put(v)
}
