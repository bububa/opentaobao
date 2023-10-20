package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingCouponQueryactivityAPIResponse 查询优惠券活动 API返回值
// alibaba.wdk.marketing.coupon.queryactivity
//
// 查询优惠券活动
type AlibabaWdkMarketingCouponQueryactivityAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingCouponQueryactivityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingCouponQueryactivityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingCouponQueryactivityAPIResponseModel).Reset()
}

// AlibabaWdkMarketingCouponQueryactivityAPIResponseModel is 查询优惠券活动 成功返回结果
type AlibabaWdkMarketingCouponQueryactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_coupon_queryactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingCouponQueryactivityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingCouponQueryactivityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingCouponQueryactivityAPIResponse)
	},
}

// GetAlibabaWdkMarketingCouponQueryactivityAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingCouponQueryactivityAPIResponse
func GetAlibabaWdkMarketingCouponQueryactivityAPIResponse() *AlibabaWdkMarketingCouponQueryactivityAPIResponse {
	return poolAlibabaWdkMarketingCouponQueryactivityAPIResponse.Get().(*AlibabaWdkMarketingCouponQueryactivityAPIResponse)
}

// ReleaseAlibabaWdkMarketingCouponQueryactivityAPIResponse 将 AlibabaWdkMarketingCouponQueryactivityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingCouponQueryactivityAPIResponse(v *AlibabaWdkMarketingCouponQueryactivityAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingCouponQueryactivityAPIResponse.Put(v)
}
