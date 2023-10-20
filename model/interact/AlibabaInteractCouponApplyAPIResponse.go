package interact

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaInteractCouponApplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractCouponApplyAPIResponseModel).Reset()
}

// AlibabaInteractCouponApplyAPIResponseModel is 优惠券领取鉴权接口 成功返回结果
type AlibabaInteractCouponApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_coupon_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 无用参数，top限制一定要有出参，增加此参数
	Stub string `json:"stub,omitempty" xml:"stub,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractCouponApplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Stub = ""
}

var poolAlibabaInteractCouponApplyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractCouponApplyAPIResponse)
	},
}

// GetAlibabaInteractCouponApplyAPIResponse 从 sync.Pool 获取 AlibabaInteractCouponApplyAPIResponse
func GetAlibabaInteractCouponApplyAPIResponse() *AlibabaInteractCouponApplyAPIResponse {
	return poolAlibabaInteractCouponApplyAPIResponse.Get().(*AlibabaInteractCouponApplyAPIResponse)
}

// ReleaseAlibabaInteractCouponApplyAPIResponse 将 AlibabaInteractCouponApplyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractCouponApplyAPIResponse(v *AlibabaInteractCouponApplyAPIResponse) {
	v.Reset()
	poolAlibabaInteractCouponApplyAPIResponse.Put(v)
}
