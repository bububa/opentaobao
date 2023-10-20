package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMobilePromotionBenefitActivitySendAPIResponse 手淘专用单用户发放接口 API返回值
// taobao.mobile.promotion.benefit.activity.send
//
// 卖家活动中需要通过该API来发放对应的权益。手淘专用单用户发放接口。
type TaobaoMobilePromotionBenefitActivitySendAPIResponse struct {
	model.CommonResponse
	TaobaoMobilePromotionBenefitActivitySendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMobilePromotionBenefitActivitySendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMobilePromotionBenefitActivitySendAPIResponseModel).Reset()
}

// TaobaoMobilePromotionBenefitActivitySendAPIResponseModel is 手淘专用单用户发放接口 成功返回结果
type TaobaoMobilePromotionBenefitActivitySendAPIResponseModel struct {
	XMLName xml.Name `xml:"mobile_promotion_benefit_activity_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 权益发放结果
	SendResult *SingleBenefitSendResult `json:"send_result,omitempty" xml:"send_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMobilePromotionBenefitActivitySendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SendResult = nil
}

var poolTaobaoMobilePromotionBenefitActivitySendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMobilePromotionBenefitActivitySendAPIResponse)
	},
}

// GetTaobaoMobilePromotionBenefitActivitySendAPIResponse 从 sync.Pool 获取 TaobaoMobilePromotionBenefitActivitySendAPIResponse
func GetTaobaoMobilePromotionBenefitActivitySendAPIResponse() *TaobaoMobilePromotionBenefitActivitySendAPIResponse {
	return poolTaobaoMobilePromotionBenefitActivitySendAPIResponse.Get().(*TaobaoMobilePromotionBenefitActivitySendAPIResponse)
}

// ReleaseTaobaoMobilePromotionBenefitActivitySendAPIResponse 将 TaobaoMobilePromotionBenefitActivitySendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMobilePromotionBenefitActivitySendAPIResponse(v *TaobaoMobilePromotionBenefitActivitySendAPIResponse) {
	v.Reset()
	poolTaobaoMobilePromotionBenefitActivitySendAPIResponse.Put(v)
}
