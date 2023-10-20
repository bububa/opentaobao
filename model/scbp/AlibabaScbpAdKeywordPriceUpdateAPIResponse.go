package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordPriceUpdateAPIResponse 关键词改价 API返回值
// alibaba.scbp.ad.keyword.price.update
//
// 关键词改价
type AlibabaScbpAdKeywordPriceUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdKeywordPriceUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordPriceUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdKeywordPriceUpdateAPIResponseModel).Reset()
}

// AlibabaScbpAdKeywordPriceUpdateAPIResponseModel is 关键词改价 成功返回结果
type AlibabaScbpAdKeywordPriceUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_price_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改关键词价格是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordPriceUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaScbpAdKeywordPriceUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdKeywordPriceUpdateAPIResponse)
	},
}

// GetAlibabaScbpAdKeywordPriceUpdateAPIResponse 从 sync.Pool 获取 AlibabaScbpAdKeywordPriceUpdateAPIResponse
func GetAlibabaScbpAdKeywordPriceUpdateAPIResponse() *AlibabaScbpAdKeywordPriceUpdateAPIResponse {
	return poolAlibabaScbpAdKeywordPriceUpdateAPIResponse.Get().(*AlibabaScbpAdKeywordPriceUpdateAPIResponse)
}

// ReleaseAlibabaScbpAdKeywordPriceUpdateAPIResponse 将 AlibabaScbpAdKeywordPriceUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdKeywordPriceUpdateAPIResponse(v *AlibabaScbpAdKeywordPriceUpdateAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdKeywordPriceUpdateAPIResponse.Put(v)
}
