package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionCouponSendAPIResponse 店铺优惠券发放接口 API返回值
// taobao.promotion.coupon.send
//
// 通过接口批量发放店铺优惠券（每次只能发送100张，只能发给当前授权卖家店铺的会员），发送成功则返回为空，发送失败则返回失败的买家列表和发送成功的买家和优惠券的number。注：如果所有买家都发放失败的话，is_success也为true，建议调用者根据返回的集合判断是否送入的买家都发放成功了
type TaobaoPromotionCouponSendAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionCouponSendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionCouponSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionCouponSendAPIResponseModel).Reset()
}

// TaobaoPromotionCouponSendAPIResponseModel is 店铺优惠券发放接口 成功返回结果
type TaobaoPromotionCouponSendAPIResponseModel struct {
	XMLName xml.Name `xml:"promotion_coupon_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 没有发送成功的买家
	FailureBuyers []ErrorMessage `json:"failure_buyers,omitempty" xml:"failure_buyers>error_message,omitempty"`
	// 发送成功的买家的昵称和优惠券的number
	CouponResults []CouponResult `json:"coupon_results,omitempty" xml:"coupon_results>coupon_result,omitempty"`
	// true 成功，false失败
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionCouponSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FailureBuyers = m.FailureBuyers[:0]
	m.CouponResults = m.CouponResults[:0]
	m.IsSuccess = false
}

var poolTaobaoPromotionCouponSendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionCouponSendAPIResponse)
	},
}

// GetTaobaoPromotionCouponSendAPIResponse 从 sync.Pool 获取 TaobaoPromotionCouponSendAPIResponse
func GetTaobaoPromotionCouponSendAPIResponse() *TaobaoPromotionCouponSendAPIResponse {
	return poolTaobaoPromotionCouponSendAPIResponse.Get().(*TaobaoPromotionCouponSendAPIResponse)
}

// ReleaseTaobaoPromotionCouponSendAPIResponse 将 TaobaoPromotionCouponSendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionCouponSendAPIResponse(v *TaobaoPromotionCouponSendAPIResponse) {
	v.Reset()
	poolTaobaoPromotionCouponSendAPIResponse.Put(v)
}
