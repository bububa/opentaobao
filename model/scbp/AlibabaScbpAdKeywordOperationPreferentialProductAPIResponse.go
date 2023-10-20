package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordOperationPreferentialProductAPIResponse 操作优推品 API返回值
// alibaba.scbp.ad.keyword.operation.preferential.product
//
// 操作优推品
type AlibabaScbpAdKeywordOperationPreferentialProductAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdKeywordOperationPreferentialProductAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordOperationPreferentialProductAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdKeywordOperationPreferentialProductAPIResponseModel).Reset()
}

// AlibabaScbpAdKeywordOperationPreferentialProductAPIResponseModel is 操作优推品 成功返回结果
type AlibabaScbpAdKeywordOperationPreferentialProductAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_operation_preferential_product_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功数量
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordOperationPreferentialProductAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
}

var poolAlibabaScbpAdKeywordOperationPreferentialProductAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdKeywordOperationPreferentialProductAPIResponse)
	},
}

// GetAlibabaScbpAdKeywordOperationPreferentialProductAPIResponse 从 sync.Pool 获取 AlibabaScbpAdKeywordOperationPreferentialProductAPIResponse
func GetAlibabaScbpAdKeywordOperationPreferentialProductAPIResponse() *AlibabaScbpAdKeywordOperationPreferentialProductAPIResponse {
	return poolAlibabaScbpAdKeywordOperationPreferentialProductAPIResponse.Get().(*AlibabaScbpAdKeywordOperationPreferentialProductAPIResponse)
}

// ReleaseAlibabaScbpAdKeywordOperationPreferentialProductAPIResponse 将 AlibabaScbpAdKeywordOperationPreferentialProductAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdKeywordOperationPreferentialProductAPIResponse(v *AlibabaScbpAdKeywordOperationPreferentialProductAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdKeywordOperationPreferentialProductAPIResponse.Put(v)
}
