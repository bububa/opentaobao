package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingExpirePromotionQueryAPIResponse 短保优惠查询 API返回值
// alibaba.wdk.marketing.expire.promotion.query
//
// 短保优惠查询
type AlibabaWdkMarketingExpirePromotionQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingExpirePromotionQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingExpirePromotionQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingExpirePromotionQueryAPIResponseModel).Reset()
}

// AlibabaWdkMarketingExpirePromotionQueryAPIResponseModel is 短保优惠查询 成功返回结果
type AlibabaWdkMarketingExpirePromotionQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_expire_promotion_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingExpirePromotionQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingExpirePromotionQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingExpirePromotionQueryAPIResponse)
	},
}

// GetAlibabaWdkMarketingExpirePromotionQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingExpirePromotionQueryAPIResponse
func GetAlibabaWdkMarketingExpirePromotionQueryAPIResponse() *AlibabaWdkMarketingExpirePromotionQueryAPIResponse {
	return poolAlibabaWdkMarketingExpirePromotionQueryAPIResponse.Get().(*AlibabaWdkMarketingExpirePromotionQueryAPIResponse)
}

// ReleaseAlibabaWdkMarketingExpirePromotionQueryAPIResponse 将 AlibabaWdkMarketingExpirePromotionQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingExpirePromotionQueryAPIResponse(v *AlibabaWdkMarketingExpirePromotionQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingExpirePromotionQueryAPIResponse.Put(v)
}
