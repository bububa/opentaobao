package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallPromotionCouponUseAPIResponse 券核销接口 API返回值
// tmall.promotion.coupon.use
//
// 核销用户的一张优惠券，返回核销结果
type TmallPromotionCouponUseAPIResponse struct {
	model.CommonResponse
	TmallPromotionCouponUseAPIResponseModel
}

// Reset 清空结构体
func (m *TmallPromotionCouponUseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallPromotionCouponUseAPIResponseModel).Reset()
}

// TmallPromotionCouponUseAPIResponseModel is 券核销接口 成功返回结果
type TmallPromotionCouponUseAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_promotion_coupon_use_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// resultCode
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// data
	Data *UseResultDo `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TmallPromotionCouponUseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ErrorMsg = ""
	m.Data = nil
}

var poolTmallPromotionCouponUseAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallPromotionCouponUseAPIResponse)
	},
}

// GetTmallPromotionCouponUseAPIResponse 从 sync.Pool 获取 TmallPromotionCouponUseAPIResponse
func GetTmallPromotionCouponUseAPIResponse() *TmallPromotionCouponUseAPIResponse {
	return poolTmallPromotionCouponUseAPIResponse.Get().(*TmallPromotionCouponUseAPIResponse)
}

// ReleaseTmallPromotionCouponUseAPIResponse 将 TmallPromotionCouponUseAPIResponse 保存到 sync.Pool
func ReleaseTmallPromotionCouponUseAPIResponse(v *TmallPromotionCouponUseAPIResponse) {
	v.Reset()
	poolTmallPromotionCouponUseAPIResponse.Put(v)
}
