package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingExpirePromotionQueryAPIResponse 短保优惠查询 API返回值
// alibaba.hm.marketing.expire.promotion.query
//
// 短保优惠查询
type AlibabaHmMarketingExpirePromotionQueryAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingExpirePromotionQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingExpirePromotionQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingExpirePromotionQueryAPIResponseModel).Reset()
}

// AlibabaHmMarketingExpirePromotionQueryAPIResponseModel is 短保优惠查询 成功返回结果
type AlibabaHmMarketingExpirePromotionQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_expire_promotion_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingExpirePromotionQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingExpirePromotionQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingExpirePromotionQueryAPIResponse)
	},
}

// GetAlibabaHmMarketingExpirePromotionQueryAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingExpirePromotionQueryAPIResponse
func GetAlibabaHmMarketingExpirePromotionQueryAPIResponse() *AlibabaHmMarketingExpirePromotionQueryAPIResponse {
	return poolAlibabaHmMarketingExpirePromotionQueryAPIResponse.Get().(*AlibabaHmMarketingExpirePromotionQueryAPIResponse)
}

// ReleaseAlibabaHmMarketingExpirePromotionQueryAPIResponse 将 AlibabaHmMarketingExpirePromotionQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingExpirePromotionQueryAPIResponse(v *AlibabaHmMarketingExpirePromotionQueryAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingExpirePromotionQueryAPIResponse.Put(v)
}
