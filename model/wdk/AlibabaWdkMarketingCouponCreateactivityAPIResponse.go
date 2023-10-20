package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingCouponCreateactivityAPIResponse 优惠券活动创建 API返回值
// alibaba.wdk.marketing.coupon.createactivity
//
// 添加优惠券活动
type AlibabaWdkMarketingCouponCreateactivityAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingCouponCreateactivityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingCouponCreateactivityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingCouponCreateactivityAPIResponseModel).Reset()
}

// AlibabaWdkMarketingCouponCreateactivityAPIResponseModel is 优惠券活动创建 成功返回结果
type AlibabaWdkMarketingCouponCreateactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_coupon_createactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建优惠券活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingCouponCreateactivityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingCouponCreateactivityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingCouponCreateactivityAPIResponse)
	},
}

// GetAlibabaWdkMarketingCouponCreateactivityAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingCouponCreateactivityAPIResponse
func GetAlibabaWdkMarketingCouponCreateactivityAPIResponse() *AlibabaWdkMarketingCouponCreateactivityAPIResponse {
	return poolAlibabaWdkMarketingCouponCreateactivityAPIResponse.Get().(*AlibabaWdkMarketingCouponCreateactivityAPIResponse)
}

// ReleaseAlibabaWdkMarketingCouponCreateactivityAPIResponse 将 AlibabaWdkMarketingCouponCreateactivityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingCouponCreateactivityAPIResponse(v *AlibabaWdkMarketingCouponCreateactivityAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingCouponCreateactivityAPIResponse.Put(v)
}
