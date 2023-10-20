package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionBenefitActivityRelationAPIResponse 关联活动权益 API返回值
// taobao.promotion.benefit.activity.relation
//
// 卖家活动中需要通过该API来关联的对应的权益。
type TaobaoPromotionBenefitActivityRelationAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionBenefitActivityRelationAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionBenefitActivityRelationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionBenefitActivityRelationAPIResponseModel).Reset()
}

// TaobaoPromotionBenefitActivityRelationAPIResponseModel is 关联活动权益 成功返回结果
type TaobaoPromotionBenefitActivityRelationAPIResponseModel struct {
	XMLName xml.Name `xml:"promotion_benefit_activity_relation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 活动关联ID
	RelationId int64 `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
	// 请求是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionBenefitActivityRelationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RelationId = 0
	m.IsSuccess = false
}

var poolTaobaoPromotionBenefitActivityRelationAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionBenefitActivityRelationAPIResponse)
	},
}

// GetTaobaoPromotionBenefitActivityRelationAPIResponse 从 sync.Pool 获取 TaobaoPromotionBenefitActivityRelationAPIResponse
func GetTaobaoPromotionBenefitActivityRelationAPIResponse() *TaobaoPromotionBenefitActivityRelationAPIResponse {
	return poolTaobaoPromotionBenefitActivityRelationAPIResponse.Get().(*TaobaoPromotionBenefitActivityRelationAPIResponse)
}

// ReleaseTaobaoPromotionBenefitActivityRelationAPIResponse 将 TaobaoPromotionBenefitActivityRelationAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionBenefitActivityRelationAPIResponse(v *TaobaoPromotionBenefitActivityRelationAPIResponse) {
	v.Reset()
	poolTaobaoPromotionBenefitActivityRelationAPIResponse.Put(v)
}
