package promotion

import (
	"encoding/xml"

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

// TaobaoPromotionCouponAddAPIResponseModel is 创建店铺优惠券接口 成功返回结果
type TaobaoPromotionCouponAddAPIResponseModel struct {
	XMLName xml.Name `xml:"promotion_coupon_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 优惠券的id
	CouponId int64 `json:"coupon_id,omitempty" xml:"coupon_id,omitempty"`
}
