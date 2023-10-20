package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionCouponApplyAPIResponse 优惠券领取 API返回值
// taobao.promotion.coupon.apply
//
// 优惠券领取
type TaobaoPromotionCouponApplyAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionCouponApplyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionCouponApplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionCouponApplyAPIResponseModel).Reset()
}

// TaobaoPromotionCouponApplyAPIResponseModel is 优惠券领取 成功返回结果
type TaobaoPromotionCouponApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"promotion_coupon_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 失败详细描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 调用错误码，只有调用失败的时候才会有
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 领取结果，领取成功为true，否则为false
	ApplyResult bool `json:"apply_result,omitempty" xml:"apply_result,omitempty"`
	// 接口调用结果，调用成功为true，否则为false
	InvokeResult bool `json:"invoke_result,omitempty" xml:"invoke_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionCouponApplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.ResultCode = ""
	m.ApplyResult = false
	m.InvokeResult = false
}

var poolTaobaoPromotionCouponApplyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionCouponApplyAPIResponse)
	},
}

// GetTaobaoPromotionCouponApplyAPIResponse 从 sync.Pool 获取 TaobaoPromotionCouponApplyAPIResponse
func GetTaobaoPromotionCouponApplyAPIResponse() *TaobaoPromotionCouponApplyAPIResponse {
	return poolTaobaoPromotionCouponApplyAPIResponse.Get().(*TaobaoPromotionCouponApplyAPIResponse)
}

// ReleaseTaobaoPromotionCouponApplyAPIResponse 将 TaobaoPromotionCouponApplyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionCouponApplyAPIResponse(v *TaobaoPromotionCouponApplyAPIResponse) {
	v.Reset()
	poolTaobaoPromotionCouponApplyAPIResponse.Put(v)
}
