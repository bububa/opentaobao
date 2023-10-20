package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDentalStatementQueryAPIResponse ISV查询对账单 API返回值
// alibaba.alihealth.dental.statement.query
//
// ISV查询对账单
type AlibabaAlihealthDentalStatementQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDentalStatementQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDentalStatementQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDentalStatementQueryAPIResponseModel).Reset()
}

// AlibabaAlihealthDentalStatementQueryAPIResponseModel is ISV查询对账单 成功返回结果
type AlibabaAlihealthDentalStatementQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_dental_statement_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAlihealthDentalStatementQueryMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDentalStatementQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDentalStatementQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDentalStatementQueryAPIResponse)
	},
}

// GetAlibabaAlihealthDentalStatementQueryAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDentalStatementQueryAPIResponse
func GetAlibabaAlihealthDentalStatementQueryAPIResponse() *AlibabaAlihealthDentalStatementQueryAPIResponse {
	return poolAlibabaAlihealthDentalStatementQueryAPIResponse.Get().(*AlibabaAlihealthDentalStatementQueryAPIResponse)
}

// ReleaseAlibabaAlihealthDentalStatementQueryAPIResponse 将 AlibabaAlihealthDentalStatementQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDentalStatementQueryAPIResponse(v *AlibabaAlihealthDentalStatementQueryAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDentalStatementQueryAPIResponse.Put(v)
}
