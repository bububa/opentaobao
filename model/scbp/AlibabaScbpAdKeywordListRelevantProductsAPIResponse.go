package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordListRelevantProductsAPIResponse 查询和词匹配的推广产品 API返回值
// alibaba.scbp.ad.keyword.list.relevant.products
//
// 查询和词匹配的推广产品
type AlibabaScbpAdKeywordListRelevantProductsAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdKeywordListRelevantProductsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordListRelevantProductsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdKeywordListRelevantProductsAPIResponseModel).Reset()
}

// AlibabaScbpAdKeywordListRelevantProductsAPIResponseModel is 查询和词匹配的推广产品 成功返回结果
type AlibabaScbpAdKeywordListRelevantProductsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_list_relevant_products_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 优推品信息返回
	ResultList []RelevantProductDto `json:"result_list,omitempty" xml:"result_list>relevant_product_dto,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordListRelevantProductsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
}

var poolAlibabaScbpAdKeywordListRelevantProductsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdKeywordListRelevantProductsAPIResponse)
	},
}

// GetAlibabaScbpAdKeywordListRelevantProductsAPIResponse 从 sync.Pool 获取 AlibabaScbpAdKeywordListRelevantProductsAPIResponse
func GetAlibabaScbpAdKeywordListRelevantProductsAPIResponse() *AlibabaScbpAdKeywordListRelevantProductsAPIResponse {
	return poolAlibabaScbpAdKeywordListRelevantProductsAPIResponse.Get().(*AlibabaScbpAdKeywordListRelevantProductsAPIResponse)
}

// ReleaseAlibabaScbpAdKeywordListRelevantProductsAPIResponse 将 AlibabaScbpAdKeywordListRelevantProductsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdKeywordListRelevantProductsAPIResponse(v *AlibabaScbpAdKeywordListRelevantProductsAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdKeywordListRelevantProductsAPIResponse.Put(v)
}
