package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionBenefitActivityDeleteAPIResponse 删除关联的活动权益 API返回值
// taobao.promotion.benefit.activity.delete
//
// 删除关联的活动权益
type TaobaoPromotionBenefitActivityDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionBenefitActivityDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionBenefitActivityDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionBenefitActivityDeleteAPIResponseModel).Reset()
}

// TaobaoPromotionBenefitActivityDeleteAPIResponseModel is 删除关联的活动权益 成功返回结果
type TaobaoPromotionBenefitActivityDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"promotion_benefit_activity_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionBenefitActivityDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoPromotionBenefitActivityDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionBenefitActivityDeleteAPIResponse)
	},
}

// GetTaobaoPromotionBenefitActivityDeleteAPIResponse 从 sync.Pool 获取 TaobaoPromotionBenefitActivityDeleteAPIResponse
func GetTaobaoPromotionBenefitActivityDeleteAPIResponse() *TaobaoPromotionBenefitActivityDeleteAPIResponse {
	return poolTaobaoPromotionBenefitActivityDeleteAPIResponse.Get().(*TaobaoPromotionBenefitActivityDeleteAPIResponse)
}

// ReleaseTaobaoPromotionBenefitActivityDeleteAPIResponse 将 TaobaoPromotionBenefitActivityDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionBenefitActivityDeleteAPIResponse(v *TaobaoPromotionBenefitActivityDeleteAPIResponse) {
	v.Reset()
	poolTaobaoPromotionBenefitActivityDeleteAPIResponse.Put(v)
}
