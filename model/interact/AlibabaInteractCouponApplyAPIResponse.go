package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractCouponApplyAPIResponse 优惠券领取鉴权接口 API返回值
// alibaba.interact.coupon.apply
//
// 鉴权接口，为coupon.apply接口鉴权
type AlibabaInteractCouponApplyAPIResponse struct {
	model.CommonResponse
	AlibabaInteractCouponApplyAPIResponseModel
}

// AlibabaInteractCouponApplyAPIResponseModel is 优惠券领取鉴权接口 成功返回结果
type AlibabaInteractCouponApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_coupon_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 无用参数，top限制一定要有出参，增加此参数
	Stub string `json:"stub,omitempty" xml:"stub,omitempty"`
}
