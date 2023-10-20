package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponse 修改关键词价格 API返回值
// alibaba.scbp.ad.keyword.update.keyword.price.batch
//
// 修改关键词价格
type AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponseModel).Reset()
}

// AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponseModel is 修改关键词价格 成功返回结果
type AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_update_keyword_price_batch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息集合
	ResultList []ErrorKeyword `json:"result_list,omitempty" xml:"result_list>error_keyword,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
}

var poolAlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponse)
	},
}

// GetAlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponse 从 sync.Pool 获取 AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponse
func GetAlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponse() *AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponse {
	return poolAlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponse.Get().(*AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponse)
}

// ReleaseAlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponse 将 AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponse(v *AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponse.Put(v)
}
