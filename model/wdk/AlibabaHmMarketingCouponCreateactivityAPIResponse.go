package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingCouponCreateactivityAPIResponse 优惠券活动创建 API返回值
// alibaba.hm.marketing.coupon.createactivity
//
// 添加优惠券活动
type AlibabaHmMarketingCouponCreateactivityAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingCouponCreateactivityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingCouponCreateactivityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingCouponCreateactivityAPIResponseModel).Reset()
}

// AlibabaHmMarketingCouponCreateactivityAPIResponseModel is 优惠券活动创建 成功返回结果
type AlibabaHmMarketingCouponCreateactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_coupon_createactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建优惠券活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingCouponCreateactivityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingCouponCreateactivityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingCouponCreateactivityAPIResponse)
	},
}

// GetAlibabaHmMarketingCouponCreateactivityAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingCouponCreateactivityAPIResponse
func GetAlibabaHmMarketingCouponCreateactivityAPIResponse() *AlibabaHmMarketingCouponCreateactivityAPIResponse {
	return poolAlibabaHmMarketingCouponCreateactivityAPIResponse.Get().(*AlibabaHmMarketingCouponCreateactivityAPIResponse)
}

// ReleaseAlibabaHmMarketingCouponCreateactivityAPIResponse 将 AlibabaHmMarketingCouponCreateactivityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingCouponCreateactivityAPIResponse(v *AlibabaHmMarketingCouponCreateactivityAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingCouponCreateactivityAPIResponse.Put(v)
}
