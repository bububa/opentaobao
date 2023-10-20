package aecreatives

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAffiliateProductQueryAPIResponse 联盟推广商品获取接口 API返回值
// aliexpress.affiliate.product.query
//
// 联盟推广商品搜索接口，用于搜索联盟推广商品数据
type AliexpressAffiliateProductQueryAPIResponse struct {
	model.CommonResponse
	AliexpressAffiliateProductQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressAffiliateProductQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressAffiliateProductQueryAPIResponseModel).Reset()
}

// AliexpressAffiliateProductQueryAPIResponseModel is 联盟推广商品获取接口 成功返回结果
type AliexpressAffiliateProductQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_affiliate_product_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	RespResult *ResponseDto `json:"resp_result,omitempty" xml:"resp_result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressAffiliateProductQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RespResult = nil
}

var poolAliexpressAffiliateProductQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressAffiliateProductQueryAPIResponse)
	},
}

// GetAliexpressAffiliateProductQueryAPIResponse 从 sync.Pool 获取 AliexpressAffiliateProductQueryAPIResponse
func GetAliexpressAffiliateProductQueryAPIResponse() *AliexpressAffiliateProductQueryAPIResponse {
	return poolAliexpressAffiliateProductQueryAPIResponse.Get().(*AliexpressAffiliateProductQueryAPIResponse)
}

// ReleaseAliexpressAffiliateProductQueryAPIResponse 将 AliexpressAffiliateProductQueryAPIResponse 保存到 sync.Pool
func ReleaseAliexpressAffiliateProductQueryAPIResponse(v *AliexpressAffiliateProductQueryAPIResponse) {
	v.Reset()
	poolAliexpressAffiliateProductQueryAPIResponse.Put(v)
}
