package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingCouponQueryitemsAPIResponse 查询优惠券活动下的商品 API返回值
// alibaba.wdk.marketing.coupon.queryitems
//
// 查询优惠券活动下面的商品
type AlibabaWdkMarketingCouponQueryitemsAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingCouponQueryitemsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingCouponQueryitemsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingCouponQueryitemsAPIResponseModel).Reset()
}

// AlibabaWdkMarketingCouponQueryitemsAPIResponseModel is 查询优惠券活动下的商品 成功返回结果
type AlibabaWdkMarketingCouponQueryitemsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_coupon_queryitems_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回结果
	Result *MarketPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingCouponQueryitemsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingCouponQueryitemsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingCouponQueryitemsAPIResponse)
	},
}

// GetAlibabaWdkMarketingCouponQueryitemsAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingCouponQueryitemsAPIResponse
func GetAlibabaWdkMarketingCouponQueryitemsAPIResponse() *AlibabaWdkMarketingCouponQueryitemsAPIResponse {
	return poolAlibabaWdkMarketingCouponQueryitemsAPIResponse.Get().(*AlibabaWdkMarketingCouponQueryitemsAPIResponse)
}

// ReleaseAlibabaWdkMarketingCouponQueryitemsAPIResponse 将 AlibabaWdkMarketingCouponQueryitemsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingCouponQueryitemsAPIResponse(v *AlibabaWdkMarketingCouponQueryitemsAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingCouponQueryitemsAPIResponse.Put(v)
}
