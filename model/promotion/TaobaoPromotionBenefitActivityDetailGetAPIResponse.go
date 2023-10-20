package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionBenefitActivityDetailGetAPIResponse 活动关联的权益详情获取 API返回值
// taobao.promotion.benefit.activity.detail.get
//
// 活动关联的权益详情获取
type TaobaoPromotionBenefitActivityDetailGetAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionBenefitActivityDetailGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionBenefitActivityDetailGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionBenefitActivityDetailGetAPIResponseModel).Reset()
}

// TaobaoPromotionBenefitActivityDetailGetAPIResponseModel is 活动关联的权益详情获取 成功返回结果
type TaobaoPromotionBenefitActivityDetailGetAPIResponseModel struct {
	XMLName xml.Name `xml:"promotion_benefit_activity_detail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 活动关联的权益详情列表
	RelationBenefitDetails string `json:"relation_benefit_details,omitempty" xml:"relation_benefit_details,omitempty"`
	// 查询是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionBenefitActivityDetailGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RelationBenefitDetails = ""
	m.IsSuccess = false
}

var poolTaobaoPromotionBenefitActivityDetailGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionBenefitActivityDetailGetAPIResponse)
	},
}

// GetTaobaoPromotionBenefitActivityDetailGetAPIResponse 从 sync.Pool 获取 TaobaoPromotionBenefitActivityDetailGetAPIResponse
func GetTaobaoPromotionBenefitActivityDetailGetAPIResponse() *TaobaoPromotionBenefitActivityDetailGetAPIResponse {
	return poolTaobaoPromotionBenefitActivityDetailGetAPIResponse.Get().(*TaobaoPromotionBenefitActivityDetailGetAPIResponse)
}

// ReleaseTaobaoPromotionBenefitActivityDetailGetAPIResponse 将 TaobaoPromotionBenefitActivityDetailGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionBenefitActivityDetailGetAPIResponse(v *TaobaoPromotionBenefitActivityDetailGetAPIResponse) {
	v.Reset()
	poolTaobaoPromotionBenefitActivityDetailGetAPIResponse.Put(v)
}
