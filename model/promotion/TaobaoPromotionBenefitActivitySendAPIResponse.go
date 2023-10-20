package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionBenefitActivitySendAPIResponse 活动权益发放接口 API返回值
// taobao.promotion.benefit.activity.send
//
// 活动权益发放接口，用于卖家针对活动进行权益发放
type TaobaoPromotionBenefitActivitySendAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionBenefitActivitySendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionBenefitActivitySendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionBenefitActivitySendAPIResponseModel).Reset()
}

// TaobaoPromotionBenefitActivitySendAPIResponseModel is 活动权益发放接口 成功返回结果
type TaobaoPromotionBenefitActivitySendAPIResponseModel struct {
	XMLName xml.Name `xml:"promotion_benefit_activity_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回活动详情级别的权益发放情况
	ResultList []BenefitSendResultExt `json:"result_list,omitempty" xml:"result_list>benefit_send_result_ext,omitempty"`
	// uniqueId
	UniqueId string `json:"unique_id,omitempty" xml:"unique_id,omitempty"`
	// 接口调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionBenefitActivitySendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
	m.UniqueId = ""
	m.IsSuccess = false
}

var poolTaobaoPromotionBenefitActivitySendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionBenefitActivitySendAPIResponse)
	},
}

// GetTaobaoPromotionBenefitActivitySendAPIResponse 从 sync.Pool 获取 TaobaoPromotionBenefitActivitySendAPIResponse
func GetTaobaoPromotionBenefitActivitySendAPIResponse() *TaobaoPromotionBenefitActivitySendAPIResponse {
	return poolTaobaoPromotionBenefitActivitySendAPIResponse.Get().(*TaobaoPromotionBenefitActivitySendAPIResponse)
}

// ReleaseTaobaoPromotionBenefitActivitySendAPIResponse 将 TaobaoPromotionBenefitActivitySendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionBenefitActivitySendAPIResponse(v *TaobaoPromotionBenefitActivitySendAPIResponse) {
	v.Reset()
	poolTaobaoPromotionBenefitActivitySendAPIResponse.Put(v)
}
