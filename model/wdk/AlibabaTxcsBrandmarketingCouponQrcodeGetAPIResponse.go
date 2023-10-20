package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponse 品牌营销导购员券页面二维码获取 API返回值
// alibaba.txcs.brandmarketing.coupon.qrcode.get
//
// 构建券页码二维码url
type AlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponse struct {
	model.CommonResponse
	AlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponseModel).Reset()
}

// AlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponseModel is 品牌营销导购员券页面二维码获取 成功返回结果
type AlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_txcs_brandmarketing_coupon_qrcode_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ApiResult *AlibabaTxcsBrandmarketingCouponQrcodeGetApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponse)
	},
}

// GetAlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponse 从 sync.Pool 获取 AlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponse
func GetAlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponse() *AlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponse {
	return poolAlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponse.Get().(*AlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponse)
}

// ReleaseAlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponse 将 AlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponse(v *AlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponse) {
	v.Reset()
	poolAlibabaTxcsBrandmarketingCouponQrcodeGetAPIResponse.Put(v)
}
