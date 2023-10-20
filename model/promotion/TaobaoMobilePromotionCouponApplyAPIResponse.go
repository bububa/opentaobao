package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMobilePromotionCouponApplyAPIResponse 优惠券领取(手淘专用) API返回值
// taobao.mobile.promotion.coupon.apply
//
// 优惠券领取
type TaobaoMobilePromotionCouponApplyAPIResponse struct {
	model.CommonResponse
	TaobaoMobilePromotionCouponApplyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMobilePromotionCouponApplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMobilePromotionCouponApplyAPIResponseModel).Reset()
}

// TaobaoMobilePromotionCouponApplyAPIResponseModel is 优惠券领取(手淘专用) 成功返回结果
type TaobaoMobilePromotionCouponApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"mobile_promotion_coupon_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 优惠券领取结果
	CouponApplyResult *CouponApplyResult `json:"coupon_apply_result,omitempty" xml:"coupon_apply_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMobilePromotionCouponApplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CouponApplyResult = nil
}

var poolTaobaoMobilePromotionCouponApplyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMobilePromotionCouponApplyAPIResponse)
	},
}

// GetTaobaoMobilePromotionCouponApplyAPIResponse 从 sync.Pool 获取 TaobaoMobilePromotionCouponApplyAPIResponse
func GetTaobaoMobilePromotionCouponApplyAPIResponse() *TaobaoMobilePromotionCouponApplyAPIResponse {
	return poolTaobaoMobilePromotionCouponApplyAPIResponse.Get().(*TaobaoMobilePromotionCouponApplyAPIResponse)
}

// ReleaseTaobaoMobilePromotionCouponApplyAPIResponse 将 TaobaoMobilePromotionCouponApplyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMobilePromotionCouponApplyAPIResponse(v *TaobaoMobilePromotionCouponApplyAPIResponse) {
	v.Reset()
	poolTaobaoMobilePromotionCouponApplyAPIResponse.Put(v)
}
