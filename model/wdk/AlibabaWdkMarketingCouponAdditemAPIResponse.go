package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingCouponAdditemAPIResponse 在优惠券活动下添加商品。【注意，此接口暂不支持并发！】 API返回值
// alibaba.wdk.marketing.coupon.additem
//
// 在优惠券活动下添加商品。【注意，此接口暂不支持并发！】
// 如果是商品券，则添加的商品为券适用的商品；
// 如果是品类券，则添加的商品为券排除的商品；
type AlibabaWdkMarketingCouponAdditemAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingCouponAdditemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingCouponAdditemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingCouponAdditemAPIResponseModel).Reset()
}

// AlibabaWdkMarketingCouponAdditemAPIResponseModel is 在优惠券活动下添加商品。【注意，此接口暂不支持并发！】 成功返回结果
type AlibabaWdkMarketingCouponAdditemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_coupon_additem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品报名活动的返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingCouponAdditemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingCouponAdditemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingCouponAdditemAPIResponse)
	},
}

// GetAlibabaWdkMarketingCouponAdditemAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingCouponAdditemAPIResponse
func GetAlibabaWdkMarketingCouponAdditemAPIResponse() *AlibabaWdkMarketingCouponAdditemAPIResponse {
	return poolAlibabaWdkMarketingCouponAdditemAPIResponse.Get().(*AlibabaWdkMarketingCouponAdditemAPIResponse)
}

// ReleaseAlibabaWdkMarketingCouponAdditemAPIResponse 将 AlibabaWdkMarketingCouponAdditemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingCouponAdditemAPIResponse(v *AlibabaWdkMarketingCouponAdditemAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingCouponAdditemAPIResponse.Put(v)
}
