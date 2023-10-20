package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordPriceBatchupdateAPIResponse 关键词批量改价 API返回值
// alibaba.scbp.ad.keyword.price.batchupdate
//
// 关键词批量改价
type AlibabaScbpAdKeywordPriceBatchupdateAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdKeywordPriceBatchupdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordPriceBatchupdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdKeywordPriceBatchupdateAPIResponseModel).Reset()
}

// AlibabaScbpAdKeywordPriceBatchupdateAPIResponseModel is 关键词批量改价 成功返回结果
type AlibabaScbpAdKeywordPriceBatchupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_price_batchupdate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改失败关键词列表
	KeywordErrorResultList []KeywordErrorResultDto `json:"keyword_error_result_list,omitempty" xml:"keyword_error_result_list>keyword_error_result_dto,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordPriceBatchupdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.KeywordErrorResultList = m.KeywordErrorResultList[:0]
}

var poolAlibabaScbpAdKeywordPriceBatchupdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdKeywordPriceBatchupdateAPIResponse)
	},
}

// GetAlibabaScbpAdKeywordPriceBatchupdateAPIResponse 从 sync.Pool 获取 AlibabaScbpAdKeywordPriceBatchupdateAPIResponse
func GetAlibabaScbpAdKeywordPriceBatchupdateAPIResponse() *AlibabaScbpAdKeywordPriceBatchupdateAPIResponse {
	return poolAlibabaScbpAdKeywordPriceBatchupdateAPIResponse.Get().(*AlibabaScbpAdKeywordPriceBatchupdateAPIResponse)
}

// ReleaseAlibabaScbpAdKeywordPriceBatchupdateAPIResponse 将 AlibabaScbpAdKeywordPriceBatchupdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdKeywordPriceBatchupdateAPIResponse(v *AlibabaScbpAdKeywordPriceBatchupdateAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdKeywordPriceBatchupdateAPIResponse.Put(v)
}
