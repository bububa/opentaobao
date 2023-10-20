package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionCouponAddAPIResponse 创建店铺优惠券接口 API返回值
// taobao.promotion.coupon.add
//
// 创建店铺优惠券。有效期内的店铺优惠券总数量不超过50张
type TaobaoPromotionCouponAddAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionCouponAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionCouponAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionCouponAddAPIResponseModel).Reset()
}

// TaobaoPromotionCouponAddAPIResponseModel is 创建店铺优惠券接口 成功返回结果
type TaobaoPromotionCouponAddAPIResponseModel struct {
	XMLName xml.Name `xml:"promotion_coupon_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 优惠券的id
	CouponId int64 `json:"coupon_id,omitempty" xml:"coupon_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionCouponAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CouponId = 0
}

var poolTaobaoPromotionCouponAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionCouponAddAPIResponse)
	},
}

// GetTaobaoPromotionCouponAddAPIResponse 从 sync.Pool 获取 TaobaoPromotionCouponAddAPIResponse
func GetTaobaoPromotionCouponAddAPIResponse() *TaobaoPromotionCouponAddAPIResponse {
	return poolTaobaoPromotionCouponAddAPIResponse.Get().(*TaobaoPromotionCouponAddAPIResponse)
}

// ReleaseTaobaoPromotionCouponAddAPIResponse 将 TaobaoPromotionCouponAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionCouponAddAPIResponse(v *TaobaoPromotionCouponAddAPIResponse) {
	v.Reset()
	poolTaobaoPromotionCouponAddAPIResponse.Put(v)
}
