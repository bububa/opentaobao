package icbuseller

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSellerCouponAuthVerifyAPIResponse 优惠券校验 API返回值
// alibaba.seller.coupon.auth.verify
//
// 优惠券校验
type AlibabaSellerCouponAuthVerifyAPIResponse struct {
	model.CommonResponse
	AlibabaSellerCouponAuthVerifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSellerCouponAuthVerifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSellerCouponAuthVerifyAPIResponseModel).Reset()
}

// AlibabaSellerCouponAuthVerifyAPIResponseModel is 优惠券校验 成功返回结果
type AlibabaSellerCouponAuthVerifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seller_coupon_auth_verify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 获取是否验证成功
	Result *AlibabaSellerCouponAuthVerifyResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSellerCouponAuthVerifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSellerCouponAuthVerifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSellerCouponAuthVerifyAPIResponse)
	},
}

// GetAlibabaSellerCouponAuthVerifyAPIResponse 从 sync.Pool 获取 AlibabaSellerCouponAuthVerifyAPIResponse
func GetAlibabaSellerCouponAuthVerifyAPIResponse() *AlibabaSellerCouponAuthVerifyAPIResponse {
	return poolAlibabaSellerCouponAuthVerifyAPIResponse.Get().(*AlibabaSellerCouponAuthVerifyAPIResponse)
}

// ReleaseAlibabaSellerCouponAuthVerifyAPIResponse 将 AlibabaSellerCouponAuthVerifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSellerCouponAuthVerifyAPIResponse(v *AlibabaSellerCouponAuthVerifyAPIResponse) {
	v.Reset()
	poolAlibabaSellerCouponAuthVerifyAPIResponse.Put(v)
}
