package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionCouponSnsSendAPIResponse 微淘粉丝店铺优惠券发放接口 API返回值
// taobao.promotion.coupon.sns.send
//
// 通过接口批量发放店铺优惠券（每次只能发送100张，只能发给当前授权卖家店铺的微淘粉丝），发送成功则返回为空，发送失败则返回失败的买家列表和发送成功的买家和优惠券的number。注：如果所有买家都发放失败的话，is_success也为true，建议调用者根据返回的集合判断是否送入的买家都发放成功了
type TaobaoPromotionCouponSnsSendAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionCouponSnsSendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionCouponSnsSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionCouponSnsSendAPIResponseModel).Reset()
}

// TaobaoPromotionCouponSnsSendAPIResponseModel is 微淘粉丝店铺优惠券发放接口 成功返回结果
type TaobaoPromotionCouponSnsSendAPIResponseModel struct {
	XMLName xml.Name `xml:"promotion_coupon_sns_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 没有发送成功的买家
	FailureBuyers []ErrorMessage `json:"failure_buyers,omitempty" xml:"failure_buyers>error_message,omitempty"`
	// 发送成功的买家的昵称和优惠券的number
	CouponResults []CouponResult `json:"coupon_results,omitempty" xml:"coupon_results>coupon_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionCouponSnsSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FailureBuyers = m.FailureBuyers[:0]
	m.CouponResults = m.CouponResults[:0]
}

var poolTaobaoPromotionCouponSnsSendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionCouponSnsSendAPIResponse)
	},
}

// GetTaobaoPromotionCouponSnsSendAPIResponse 从 sync.Pool 获取 TaobaoPromotionCouponSnsSendAPIResponse
func GetTaobaoPromotionCouponSnsSendAPIResponse() *TaobaoPromotionCouponSnsSendAPIResponse {
	return poolTaobaoPromotionCouponSnsSendAPIResponse.Get().(*TaobaoPromotionCouponSnsSendAPIResponse)
}

// ReleaseTaobaoPromotionCouponSnsSendAPIResponse 将 TaobaoPromotionCouponSnsSendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionCouponSnsSendAPIResponse(v *TaobaoPromotionCouponSnsSendAPIResponse) {
	v.Reset()
	poolTaobaoPromotionCouponSnsSendAPIResponse.Put(v)
}
